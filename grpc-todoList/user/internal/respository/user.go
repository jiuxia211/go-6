package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"grpc-todoList/user/internal/service"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	NickName       string
}

const (
	PasswordCost = 12 //密码加密难度
)

func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("用户名不存在")
}
func (user *User) UserCreate(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return errors.New("用户名已注册")
	}
	_ = user.SetPassword(req.Password)
	user.UserName = req.UserName
	user.NickName = req.NickName
	fmt.Println("zz", user)
	if err := DB.Create(&user).Error; err != nil {
		fmt.Println("创建用户失败")
	}
	fmt.Println("zzz", user)
	return nil
}
func (user *User) SetPassword(password string) error {
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	} else {
		user.PasswordDigest = string(passwordDigest)
		return nil
	}
}
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.ID),
		UserName: item.UserName,
		NickName: item.NickName,
	}
	return &userModel
}
