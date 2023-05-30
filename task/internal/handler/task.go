package handler

import (
	"context"
	"grpc-todoList/pkg/e"
	repository "grpc-todoList/task/internal/respository"
	"grpc-todoList/task/internal/service"
)

type TaskService struct {
	service.UnimplementedTaskServiceServer
}

func NewTaskService() *TaskService {
	return &TaskService{}
}
func (*TaskService) TaskCreate(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err = task.TaskCreate(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	var tasks []*service.TaskModel
	tasks = append(tasks, repository.BuildTask(task))
	resp.TaskDetail = tasks
	return resp, nil
}
func (*TaskService) SetTask(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err = task.SetTask(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	var tasks []*service.TaskModel
	tasks = append(tasks, repository.BuildTask(task))
	resp.TaskDetail = tasks
	return resp, nil
}
func (*TaskService) GetAllTask(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err, tasks := task.GetAllTask(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.TaskDetail = repository.BuildTasks(tasks)
	return resp, nil
}
func (*TaskService) GetFinishTask(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err, tasks := task.GetFinishTask(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.TaskDetail = repository.BuildTasks(tasks)
	return resp, nil
}
func (*TaskService) GetUnFinishTask(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err, tasks := task.GetUnFinishTask(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.TaskDetail = repository.BuildTasks(tasks)
	return resp, nil
}
func (*TaskService) Delete(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err = task.DeleteTask(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	return resp, nil
}
func (*TaskService) DeleteFinishTask(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err = task.DeleteFinishTask(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	return resp, nil
}
func (*TaskService) DeleteUnFinishTask(ctx context.Context, req *service.TaskRequest) (resp *service.TaskDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TaskDetailResponse)
	resp.Code = e.Success
	err = task.DeleteUnFinishTask(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	return resp, nil
}
