package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"grpc-todoList/gateway/internal/service"
	"grpc-todoList/pkg/e"
	"grpc-todoList/pkg/res"
	"grpc-todoList/pkg/util"
	"strconv"
)

func TaskCreate(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskCreate(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
func SetTask(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	tid, _ := strconv.ParseInt(c.Param("tid"), 10, 64)
	taskReq.TaskID = uint32(tid)
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.SetTask(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
func GetAllTask(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	println(taskReq.Page, "zz", taskReq.Title)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.GetAllTask(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
func GetFinishTask(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.GetFinishTask(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
func GetUnFinishTask(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.GetUnFinishTask(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
func Delete(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	tid, _ := strconv.ParseInt(c.Param("tid"), 10, 64)
	taskReq.TaskID = uint32(tid)
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.Delete(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
func DeleteFinishTask(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.DeleteFinishTask(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
func DeleteUnFinishTask(c *gin.Context) {
	var taskReq service.TaskRequest
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.Bind(&taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	taskReq.Uid = uint32(uid)
	// gin.Key中获取用户实例
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.DeleteUnFinishTask(context.Background(), &taskReq)
	if err != nil {
		panic(errors.New("Task Service--" + err.Error()))
	}
	r := res.Response{Data: taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code))}
	c.JSON(200, r)
}
