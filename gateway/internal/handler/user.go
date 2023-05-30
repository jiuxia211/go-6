package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"grpc-todoList/gateway/internal/service"
	"grpc-todoList/pkg/e"
	"grpc-todoList/pkg/res"
	"grpc-todoList/pkg/util"
)

func UserRegister(c *gin.Context) {
	var userReq service.UserRequest
	err := c.Bind(&userReq)
	if err != nil {
		panic(errors.New("User Service--" + err.Error()))
	}
	// gin.Key中获取用户实例
	userService := c.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	if err != nil {
		panic(errors.New("User Service--" + err.Error()))
	}
	r := res.Response{Data: userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code))}
	c.JSON(200, r)
}
func UserLogin(c *gin.Context) {
	var userReq service.UserRequest
	err := c.Bind(&userReq)
	if err != nil {
		panic(errors.New("User Service--" + err.Error()))
	}
	userService := c.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	if err != nil {
		panic(errors.New("User Service--" + err.Error()))
	}
	token, err := util.GenerateToken(userReq.UserName, uint(userResp.UserDetail.UserID))
	r := res.Response{Data: res.TokenData{
		User:  userResp.UserDetail,
		Token: token,
	},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code))}
	c.JSON(200, r)
}
