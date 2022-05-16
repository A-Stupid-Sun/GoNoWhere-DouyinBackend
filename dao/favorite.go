package dao

import "douyin/model"

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
