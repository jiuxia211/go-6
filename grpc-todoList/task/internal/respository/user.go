package repository

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	NickName       string
}
