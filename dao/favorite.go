package dao

import (
	"douyin/model"
)

type favoriteDAO struct{}

// QueryCountOfVideo 查询某一条视频的点赞数
func (*favoriteDAO) QueryCountOfVideo(conditions map[string]interface{}) (int64, error) {
	var count int64
	err := db.Model(&model.Favorite{}).
		Select("id").
		Where(conditions).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

// IsFavorite 查询 userID 的用户是否对videoID的视频点赞
// 如果以及点赞返回 true，否则返回false
// 如果中间发送错误，返回false
func (*favoriteDAO) IsFavorite(userID int64, videoID int64) bool {
	var count int64
	err := db.Model(&model.Favorite{}).
		Select("id").
		Where("user_id = ? AND video_id = ?", userID, videoID).
		Count(&count).Error
	if err != nil {
		return false
	}
	return count != 0
}

// Add 增加 一条点赞记录
func (*favoriteDAO) Add(userID, videoID int64) error {
	f := model.Favorite{
		UserID:  userID,
		VideoID: videoID,
	}
	err := db.Model(&model.Favorite{}).
		Create(&f).Error

	if err != nil {
		return err
	}
	return nil
}

// Sub 删除一条点赞记录
func (*favoriteDAO) Sub(userID, videoID int64) error {
	f := model.Favorite{
		UserID:  userID,
		VideoID: videoID,
	}
	err := db.Model(&model.Favorite{}).
		Delete(&f).Error

	if err != nil {
		return err
	}
	return nil
}
