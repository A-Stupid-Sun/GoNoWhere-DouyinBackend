package dao

import (
	"douyin/model"
	"errors"

	"gorm.io/gorm"
)

type commentDAO struct{}

// Add 增加评论的操作， c 里面只包含userID、videoID、content三项内容
func (*commentDAO) Add(c *model.Comment) error {
	err := db.Model(&model.Comment{}).
		Create(c).
		Error
	return err
}

// Delete 删除评论，只需提供 commentID 即可
func (*commentDAO) Delete(id int64) error {
	err := db.Model(&model.Comment{}).Delete(&model.Comment{}, 10).Error
	return err
}

func (*commentDAO) List(videoID int64) ([]model.Comment, error) {
	var c []model.Comment
	err := db.Model(&model.Comment{}).
		Select([]string{"id", "content", "user_id", "create_at"}).
		Where("video_id = ?", videoID).
		Find(&c).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return c, nil
}
