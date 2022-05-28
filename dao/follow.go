package dao

import "douyin/model"

type followDAO struct{}

// IsFollow userID 用户是否以及关注 toUserID 用户
//如果关注 返回ture ，发生异常或者不存在粉丝关系否则返回 false
func (*followDAO) IsFollow(userID, toUserID int64) bool {
	var count int64
	err := db.Model(&model.Follow{}).
		Select("id").
		Where("user_id =? AND to_user_id = ?", userID, toUserID).
		Count(&count).Error
	if err != nil {
		return false
	}

	return count != 0
}
