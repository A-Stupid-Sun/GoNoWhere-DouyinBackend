package dao

import (
	"douyin/model"
	"log"
)

type videoDAO struct{}

// Create 新增数据
func (*videoDAO) Create(values map[string]interface{}) error {
	err := db.Model(&model.Video{}).Create(values).Error
	if err != nil {
		log.Print("新增视频实录失败")
	}
	return err
}

// Query 根据参数中的条件查询视频
func (*videoDAO) Query(conditions map[string]interface{}, fields ...string) ([]model.Video, error) {
	var v []model.Video
	err := db.Model(&model.Video{}).
		Select(fields).
		Where(conditions).
		Find(&v).Error
	if err != nil {
		return nil, err
	}

	return v, nil
}
