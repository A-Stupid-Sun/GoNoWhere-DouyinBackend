package model

import (
	"douyin/config"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// BeforeSave 用户注册时，在保存数据库之前对密码进行加密,使用接口导致钩子调用失败

func (u *UserLogin) BeforeSave(tx *gorm.DB) (err error) {
	p, err := bcrypt.GenerateFromPassword([]byte(u.PassWord), config.BcryptCost)
	if err != nil {
		return errors.New("创建用户失败")
	}
	u.PassWord = string(p)
	fmt.Println(u.PassWord)
	return
}
