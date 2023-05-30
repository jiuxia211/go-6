package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"grpc-todoList/task/internal/service"
)

type Task struct {
	gorm.Model
	Uid       uint32
	Title     string
	Content   string
	Status    uint32
	StartTime uint32
	EndTime   uint32
}

func (task *Task) TaskCreate(req *service.TaskRequest) (err error) {
	task.Uid = req.Uid
	task.Title = req.Title
	task.Content = req.Content
	task.Status = req.Status
	task.StartTime = req.StartTime
	task.EndTime = req.EndTime
	if err := DB.Create(&task).Error; err != nil {
		fmt.Println("创建事务失败")
		return err
	}
	return nil
}
func (task *Task) SetTask(req *service.TaskRequest) (err error) {
	if err := DB.Where("id=?", req.TaskID).First(&task).Error; err != nil {
		fmt.Println("未找到对应事务")
		return err
	}
	if task.Uid != req.Uid {
		return errors.New("没有修改该事务的权限")
	}
	task.Uid = req.Uid
	task.Title = req.Title
	task.Content = req.Content
	task.Status = req.Status
	task.StartTime = req.StartTime
	task.EndTime = req.EndTime
	if err := DB.Save(&task).Error; err != nil {
		fmt.Println("更新事务信息失败")
		return err
	}
	return nil
}
func (task *Task) GetAllTask(req *service.TaskRequest) (err error, tasks []Task) {
	fmt.Println(req.Uid, "zz", req.Page)
	err = DB.Where("uid=?", req.Uid).Limit(5).Offset((req.Page - 1) * 5).Find(&tasks).Error
	if err != nil {
		fmt.Println("查询失败")
		return err, nil
	}
	return nil, tasks
}
func (task *Task) GetFinishTask(req *service.TaskRequest) (err error, tasks []Task) {
	err = DB.Where("uid=? AND status=?", req.Uid, 1).Limit(5).Offset((req.Page - 1) * 5).Find(&tasks).Error
	if err != nil {
		fmt.Println("查询失败")
		return err, nil
	}
	return nil, tasks
}
func (task *Task) GetUnFinishTask(req *service.TaskRequest) (err error, tasks []Task) {
	err = DB.Where("uid=? AND status=?", req.Uid, 0).Limit(5).Offset((req.Page - 1) * 5).Find(&tasks).Error
	if err != nil {
		fmt.Println("查询失败")
		return err, nil
	}
	return nil, tasks
}
func (task *Task) DeleteTask(req *service.TaskRequest) (err error) {
	if err = DB.First(&task, req.TaskID).Error; err == gorm.ErrRecordNotFound {
		fmt.Println("未找到事务")
		return err
	}
	if task.Uid != req.Uid {
		return errors.New("没有修改该事务的权限")
	}
	if err = DB.Delete(&Task{}, req.TaskID).Error; err != nil {
		fmt.Println("删除失败")
		return err
	}
	return nil
}
func (task *Task) DeleteFinishTask(req *service.TaskRequest) (err error) {
	if err = DB.Where("uid=? AND status=?", req.Uid, 1).Delete(&Task{}).Error; err != nil {
		fmt.Println("删除失败")
		return err
	}
	return nil
}
func (task *Task) DeleteUnFinishTask(req *service.TaskRequest) (err error) {
	if err = DB.Where("uid=? AND status=?", req.Uid, 0).Delete(&Task{}).Error; err != nil {
		fmt.Println("删除失败")
		return err
	}
	return nil
}
func BuildTask(item Task) *service.TaskModel {
	taskModel := service.TaskModel{
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
		Tid:       uint32(item.ID),
	}
	return &taskModel
}
func BuildTasks(items []Task) (tasks []*service.TaskModel) {
	for _, item := range items {
		tasks = append(tasks, BuildTask(item))
	}
	return tasks
}
