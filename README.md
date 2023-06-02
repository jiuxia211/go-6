# go-6

gin+grpc+etcd

西二在线golang第六轮考核~~~~

接口文档 https://www.showdoc.com.cn/2275177058881415?page_id=10206616509552948

运行

1.在config.ini 的mysql部分修改对应的数据库信息，并先手动创建etcd

2.启动etcd

启动用户模块 make build-user

启动事务模块 make build-task

启动网关 make build-gateway