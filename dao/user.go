package dao

import (
	"douyin/model"
)

type userDAO struct{}

// Create 根据 values 里面的参数 （k-v） 创建 User 实体
func (*userDAO) Create(values map[string]interface{}) error {
	//TODO 增
	err := db.Model(&model.User{}).Create(values).Error
	if err != nil {
		return err
	}

	return nil
}

func (*userDAO) Delete() {
	//TODO 删
}

func (*userDAO) Update() {
	//TODO 改
}

func (*userDAO) Query(conditions map[string]interface{}) (*model.User, error) {
	var u model.User
	err := db.Model(&model.User{}).
		Select([]string{"user_id", "follow_count", "follower_count"}).
		Where(conditions).
		First(&u).Error

	if err != nil {
		return nil, err
	}

	return &u, nil
}
