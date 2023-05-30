package handler

import (
	"context"
	"grpc-todoList/pkg/e"
	repository "grpc-todoList/user/internal/respository"
	"grpc-todoList/user/internal/service"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success
	err = user.UserCreate(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}

	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}
