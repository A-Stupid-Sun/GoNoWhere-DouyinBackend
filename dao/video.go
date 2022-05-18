package dao

import (
	"douyin/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
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

// QueryLatest 查询最新的若干条视频，但保证不超过 latestTime
func (*videoDAO) QueryLatest(latestTime string) ([]model.Video, error) {
	fmt.Println("videoDAO:", latestTime)
	var v []model.Video
	err := db.Model(&model.Video{}).
		Order("create_at desc").
		Select([]string{"author_id", "play_url", "cover_url", "favorite_count", "comment_count", "create_at"}).
		Where("create_at < ?", latestTime).
		Find(&v).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return v, nil
}
