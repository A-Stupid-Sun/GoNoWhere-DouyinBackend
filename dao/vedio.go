package dao

import (
	"douyin/model"
	"log"
)

type videoDAO struct{}

// Create 新增数据
func (*videoDAO) Create(condition map[string]interface{}) error {
	err := db.Model(&model.Video{}).Create(condition).Error
	if err != nil {
		log.Print("新增视频实录失败")
	}
	return err
}
