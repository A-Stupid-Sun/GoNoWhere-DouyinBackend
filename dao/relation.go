package dao

import (
	"douyin/model"
	"errors"

	"gorm.io/gorm"
)

// Create 新增一对关注记录，即 userID->ToUserID
// 时间由数据库自动创建
func (*followDAO) Create(f *model.Follow) error {
	err := db.Model(&model.Follow{}).Create(f).Error
	return err
}

// Delete 删除一对关注记录，对应取关和拉黑操作
func (*followDAO) Delete(f *model.Follow) error {
	err := db.Model(&model.Follow{}).Where("user_id =? AND to_user_id =?", f.UserID, f.ToUserID).Error
	return err
}

// Query 查询一定条件下的用户列表
// 返回的是 User 的切片，包含所有满足条件的用户(只包含 field 字段)
func (*followDAO) Query(conditions map[string]interface{}, field ...string) ([]model.Follow, error) {
	var f []model.Follow
	err := db.Model(&model.Follow{}).Where(conditions).Select(field).Find(&f).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return f, nil
}
