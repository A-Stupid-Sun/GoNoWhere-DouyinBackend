package dao

import (
	"douyin/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type videoDAO struct{}

const kMAXVideoCount = 30

// Create 新增数据
func (*videoDAO) Create(values map[string]interface{}) error {
	err := db.Model(&model.Video{}).Create(values).Error
	if err != nil {
		log.Print("新增视频实录失败")
	}
	return err
}

// QueryVideos 根据参数中的条件查询视频
func (*videoDAO) QueryVideos(conditions map[string]interface{}, fields ...string) ([]model.Video, error) {
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
	var v []model.Video
	err := db.Model(&model.Video{}).
		Order("create_at desc").
		Select([]string{"author_id", "play_url", "cover_url", "favorite_count", "comment_count", "create_at"}).
		Where("create_at < ?", latestTime).
		Find(&v).
		Limit(kMAXVideoCount).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return v, nil
}

// AddFavorite 增加 count 数量的点赞数
// count 参数主要为了更好的扩展性
// 使用 count 可以批量更新减少更新时带来的一些问题
func (*videoDAO) AddFavorite(videoID int64, count int) error {
	err := db.Model(model.Video{}).
		Where("video_id = ?", videoID).
		Update("favorite_count", gorm.Expr("favorite_count + ?", count)).
		Error

	if err != nil {
		return err
	}
	return nil
}

// SubFavorite  减少 count 数量的点赞数
// count 参数主要为了更好的扩展性
// 使用 count 可以批量更新减少更新时带来的一些问题
func (*videoDAO) SubFavorite(videoID int64, count int) error {
	err := db.Model(model.Video{}).
		Update("favorite_count", gorm.Expr("favorite_count + ?", count)).
		Error

	if err != nil {
		return err
	}
	return nil
}

// QueryVideo 根据单个videoID查询单个视频信息
// 返回一个视频的信息
func (*videoDAO) QueryVideo(videoID int64, fields ...string) (model.Video, error) {
	var v model.Video
	err := db.Model(&model.Video{}).
		Select(fields).
		Where("video_id = ?", videoID).
		Error

	if err != nil {
		return model.Video{}, err
	}
	return v, nil
}

// QueryVideosByID 通过 ID 列表批量查询视频
func (*videoDAO) QueryVideosByID(IDList []int64, fields ...string) ([]model.Video, error) {
	var v []model.Video
	err := db.Model(&model.Video{}).
		Select(fields).
		Where("video_id IN ? ", IDList).
		Find(&v).
		Error
	// 处理错误和未查询到记录的err
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return v, nil
}
