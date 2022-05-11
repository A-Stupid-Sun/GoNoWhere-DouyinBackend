package dao

import (
	"douyin/model"
)

type userLoginDAO struct{}

func (l *userLoginDAO) Query(conditions map[string]interface{}, fields []string) (*model.UserLogin, error) {
	var u model.UserLogin
	err := db.Model(&model.UserLogin{}).
		Select(fields).
		Where(conditions).
		First(&u).Error

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (l *userLoginDAO) Create(values map[string]interface{}) error {
	err := db.Model(&model.UserLogin{}).Create(values).Error
	if err != nil {
		return err
	}

	return nil
}
