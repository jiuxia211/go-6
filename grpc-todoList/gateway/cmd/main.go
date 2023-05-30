package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	config "grpc-todoList/conf"
	"grpc-todoList/gateway/internal/service"
	"grpc-todoList/gateway/routes"
	"grpc-todoList/pkg/discovery"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Register   *discovery.Resolver
	CancelFunc context.CancelFunc
	ctx        context.Context
)

func main() {
	config.Init()
	Register = discovery.NewResolver([]string{config.EtcdAddress}, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)
	go startListen(ctx)
	{
		//创建一个缓冲大小为1的信号通道，用于接收操作系统发送的信号。
		osSignals := make(chan os.Signal, 1)
		//函数将指定的信号注册到信号通道osSignals上。
		//在这个示例中，我们注册了os.Interrupt、os.Kill、syscall.SIGTERM、syscall.SIGINT和syscall.SIGKILL这些信号。
		//中断信号（Interrupt）：Ctrl+C
		//终止进程信号（Kill）：Ctrl+\
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
}
func startListen(ctx context.Context) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")),
	}
	addrUser := fmt.Sprintf("%s:///%s", Register.Scheme(), config.Domain2)
	addrTask := fmt.Sprintf("%s:///%s", Register.Scheme(), config.Domain3)
	fmt.Println(addrUser)
	fmt.Println(addrTask)
	userConn, _ := grpc.DialContext(ctx, addrUser, opts...)
	taskConn, _ := grpc.DialContext(ctx, addrTask, opts...)
	//这个userService里存着可调用的接口
	userService := service.NewUserServiceClient(userConn)
	taskService := service.NewTaskServiceClient(taskConn)
	fmt.Println(userService, "zzz", taskService)
	ginRouter := routes.NewRouter(userService, taskService)
	server := &http.Server{
		Addr:           config.HttpPort,
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}

}
