// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.5
// source: taskService.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaskService_TaskCreate_FullMethodName         = "/pb.TaskService/TaskCreate"
	TaskService_SetTask_FullMethodName            = "/pb.TaskService/SetTask"
	TaskService_GetAllTask_FullMethodName         = "/pb.TaskService/GetAllTask"
	TaskService_GetFinishTask_FullMethodName      = "/pb.TaskService/GetFinishTask"
	TaskService_GetUnFinishTask_FullMethodName    = "/pb.TaskService/GetUnFinishTask"
	TaskService_Delete_FullMethodName             = "/pb.TaskService/Delete"
	TaskService_DeleteFinishTask_FullMethodName   = "/pb.TaskService/DeleteFinishTask"
	TaskService_DeleteUnFinishTask_FullMethodName = "/pb.TaskService/DeleteUnFinishTask"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	TaskCreate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	SetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	GetAllTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	GetFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	GetUnFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	Delete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	DeleteFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	DeleteUnFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) TaskCreate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_TaskCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_SetTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetAllTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_GetAllTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_GetFinishTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetUnFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_GetUnFinishTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Delete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_DeleteFinishTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteUnFinishTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_DeleteUnFinishTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	TaskCreate(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	SetTask(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	GetAllTask(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	GetFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	GetUnFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	Delete(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	DeleteFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	DeleteUnFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) TaskCreate(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreate not implemented")
}
func (UnimplementedTaskServiceServer) SetTask(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTask not implemented")
}
func (UnimplementedTaskServiceServer) GetAllTask(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTask not implemented")
}
func (UnimplementedTaskServiceServer) GetFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinishTask not implemented")
}
func (UnimplementedTaskServiceServer) GetUnFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnFinishTask not implemented")
}
func (UnimplementedTaskServiceServer) Delete(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTaskServiceServer) DeleteFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFinishTask not implemented")
}
func (UnimplementedTaskServiceServer) DeleteUnFinishTask(context.Context, *TaskRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUnFinishTask not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_TaskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_TaskCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskCreate(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_SetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SetTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetAllTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetAllTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetAllTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetAllTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetFinishTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetFinishTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetFinishTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetFinishTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetUnFinishTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetUnFinishTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetUnFinishTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetUnFinishTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Delete(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteFinishTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteFinishTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_DeleteFinishTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteFinishTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteUnFinishTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteUnFinishTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_DeleteUnFinishTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteUnFinishTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskCreate",
			Handler:    _TaskService_TaskCreate_Handler,
		},
		{
			MethodName: "SetTask",
			Handler:    _TaskService_SetTask_Handler,
		},
		{
			MethodName: "GetAllTask",
			Handler:    _TaskService_GetAllTask_Handler,
		},
		{
			MethodName: "GetFinishTask",
			Handler:    _TaskService_GetFinishTask_Handler,
		},
		{
			MethodName: "GetUnFinishTask",
			Handler:    _TaskService_GetUnFinishTask_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TaskService_Delete_Handler,
		},
		{
			MethodName: "DeleteFinishTask",
			Handler:    _TaskService_DeleteFinishTask_Handler,
		},
		{
			MethodName: "DeleteUnFinishTask",
			Handler:    _TaskService_DeleteUnFinishTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taskService.proto",
}
