package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	mvccpb2 "go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`    // 地址
	Version string `json:"version"` // 版本
	Weight  int64  `json:"weight"`  // 权重
}

func (server *Server) Register(etcdAddr []string) {
	fmt.Println("zzzzz")
	//先创建一个客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second})
	if err != nil {
		fmt.Println("创建etcd客户端失败")
	}
	//创建lease对象
	lease := clientv3.NewLease(cli)
	//创建租约
	leaseResp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println("创建租约失败")
	}

	//获取租约id
	leaseId := leaseResp.ID
	keepAliveChan, err := lease.KeepAlive(context.TODO(), leaseId)
	if err != nil {
		fmt.Println("自动永久续租失败")
	}

	go func() {
		for {
			select {
			case keepResp := <-keepAliveChan:
				if keepResp != nil {
					fmt.Println("自动续租中")
				} else {
					fmt.Println("租约失效了")
					server.Unregister(cli)
				}
			}

		}
	}()
	data, err := json.Marshal(server)
	kv := clientv3.NewKV(cli)
	fmt.Println(BuildRegisterPath(*server), string(data))
	_, err = kv.Put(context.TODO(), BuildRegisterPath(*server), string(data), clientv3.WithLease(leaseId))
	if err != nil {
		fmt.Println("put Server 失败")
	}

}
func (server *Server) Unregister(cli *clientv3.Client) {
	kv := clientv3.NewKV(cli)
	_, err := kv.Delete(context.TODO(), server.Name)
	if err != nil {
		fmt.Println("删除节点失败")
	}
}

// NewMaster 这个方法暂时没用到
func (server *Server) NewMaster(etcdAddr []string) {
	//先创建一个客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second})
	if err != nil {
		fmt.Println("创建etcd客户端失败")
	}
	watchRespChan := cli.Watch(context.TODO(), "jiuxia-key")
	kv := clientv3.NewKV(cli)
	getResp, err := kv.Get(context.TODO(), "jiuxia-key")
	if err != nil {
		fmt.Println(err)
	}
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值", string(getResp.Kvs[0].Value), "正常运行")
	}
	go watch(watchRespChan)

}
func watch(watchRespChan clientv3.WatchChan) {
	for {
		select {
		case res, ok := <-watchRespChan:
			if ok {
				for _, event := range res.Events {
					switch event.Type {
					case mvccpb2.PUT:
						fmt.Println(event.Kv.Key, "已上线:", string(event.Kv.Value), "Revision:",
							event.Kv.CreateRevision, event.Kv.ModRevision)
					case mvccpb2.DELETE:
						fmt.Println("删除了", "Revision:", event.Kv.ModRevision, event.Kv.Value)
					}
				}

			}
		}
	}
}
