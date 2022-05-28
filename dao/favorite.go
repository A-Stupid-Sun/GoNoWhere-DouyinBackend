package dao

import (
	"douyin/model"
	"errors"

	"gorm.io/gorm"
)

type favoriteDAO struct{}

// QueryCountOfFavorite 从点赞记录里面查询某一条视频的点赞数
// 从点赞记录里面获取的记录一定是最真实的
func (*favoriteDAO) QueryCountOfFavorite(conditions map[string]interface{}) (int64, error) {
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
	return db.Model(&model.Favorite{}).Create(&f).Error

}

// Sub 删除一条点赞记录
func (*favoriteDAO) Sub(userID, videoID int64) error {
	f := model.Favorite{
		UserID:  userID,
		VideoID: videoID,
	}
	return db.Model(&model.Favorite{}).Delete(&f).Error
}

// FavoriteListByUserID 获取某用户点赞的所有视频的 ID 列表
// 对于 favorite 实体来说，只有对应的 userID 或者 videoID 才是比较重要的，所以这里就只返回一个int64
// 类型的切片，而不是 favorite 类型的切片
func (*favoriteDAO) FavoriteListByUserID(userID int64) ([]int64, error) {
	var f []model.Favorite
	err := db.Model(&model.Favorite{}).
		Select("video_id").
		Where("user_id").
		Find(&f).Error

	// 除了判断err不为空，也要处理 RecordNotFound 问题
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	var res []int64
	for _, i := range f {
		res = append(res, i.VideoID)
	}
	return res, nil
}
