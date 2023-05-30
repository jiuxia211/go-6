package main

import (
	"google.golang.org/grpc"
	config "grpc-todoList/conf"
	"grpc-todoList/pkg/discovery"
	"grpc-todoList/task/internal/handler"
	repository "grpc-todoList/task/internal/respository"
	"grpc-todoList/task/internal/service"
	"net"
)

func main() {
	config.Init()
	repository.Database()
	etcdAddress := []string{config.EtcdAddress}
	//连接etcd
	taskServer := discovery.Server{Name: config.Domain3,
		Addr: config.GrpcAddress2}
	taskServer.Register(etcdAddress)
	//创建grpc服务
	server := grpc.NewServer()
	defer server.Stop()
	//绑定服务
	service.RegisterTaskServiceServer(server, handler.NewTaskService())
	//开启端口
	lis, err := net.Listen("tcp", config.GrpcAddress2)
	if err != nil {
		panic(err)
	}
	//启动服务
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
