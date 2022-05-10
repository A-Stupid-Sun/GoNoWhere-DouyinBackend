package dao

import (
	"douyin/model"
	"fmt"
)

type userDAO struct{}

//Demo

func (u *userDAO) Create(values map[string]interface{}) error {
	//TODO 增
	fmt.Println(db)
	err := db.Model(&model.User{}).Create(values).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userDAO) Delete() {
	//TODO 删
}

func (u *userDAO) Update() {
	//TODO 改
}

func (u *userDAO) Query(conditions map[string]interface{}) (*model.User, error) {
	//TODO 查
	return nil, nil
}
