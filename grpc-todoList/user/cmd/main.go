package main

import (
	"google.golang.org/grpc"
	config "grpc-todoList/conf"
	"grpc-todoList/pkg/discovery"
	"grpc-todoList/user/internal/handler"
	repository "grpc-todoList/user/internal/respository"
	"grpc-todoList/user/internal/service"
	"net"
)

func main() {
	config.Init()
	repository.Database()
	etcdAddress := []string{config.EtcdAddress}
	//连接etcd
	userServer := discovery.Server{Name: config.Domain2,
		Addr: config.GrpcAddress1}
	userServer.Register(etcdAddress)
	//创建grpc服务
	server := grpc.NewServer()
	defer server.Stop()
	//绑定服务
	service.RegisterUserServiceServer(server, handler.NewUserService())
	//开启端口
	lis, err := net.Listen("tcp", config.GrpcAddress1)
	if err != nil {
		panic(err)
	}
	//启动服务
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
