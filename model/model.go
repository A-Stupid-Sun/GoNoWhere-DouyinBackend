package model

import (
	"gorm.io/gorm"
)

// User 实体
// gorm.model 包含自增主键id、以及创建、更新、删除时间四项，详情见文档 https://gorm.io/zh_CN/docs/models.html#gorm-Model
// 其他所有字段非空，一是业务要求，二是优化数据库性能（细节不表）
//
type User struct {
	gorm.Model
	UserID        int64   `gorm:"type:BIGINT;unsigned;not null;unique" json:"user_id"`
	Name          string  `gorm:"type:varchar(34);not null;unique;" json:"name" validate:"min=6,max=32"`
	PassWord      string  `gorm:"type:varchar(65);not null;" json:"password" validate:"min=6,max=32"`
	FollowCount   int64   `gorm:"type:INT;unsigned;not null;default:0" json:"follow_count"`
	FollowerCount int64   `gorm:"type:INT;unsigned;not null;default:0" json:"follower_count"`
	Videos        []Video `gorm:"foreignKey:AuthorID;references:UserID"`
}

// UserAPI 主要提供给接口使用
type UserAPI struct {
	UserID        int64  `json:"user_id omitempty"`
	Name          string `json:"name omitempty"`
	FollowCount   int    `json:"follow_count omitempty"`
	FollowerCount int    `json:"follower_count omitempty"`
	IsFollow      bool   `json:"is_follow omitempty"`
}

// UserLogin （备选） 登录数据单独存放，因为登录、注册、改密操作相对来说比较少，算是冷数据吧
// 提供的接口只有用户名登录，那么只包含用户名和密码即可
type UserLogin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(34);not null;unique" json:"name" validate:"min=6,max=32"`
	PassWord string `gorm:"type:varchar(65);not null" json:"password" validate:"min=6,max=32"`
}

// Video 实体
// 发布时间要创建索引，加速按照时间访问
type Video struct {
	gorm.Model
	VideoID       int64  `gorm:"type:BIGINT;not null;UNIQUE" json:"video_id" validate:""`
	AuthorID      int64  `gorm:"type:BIGINT;not null" json:"author_id" validate:""`
	FavoriteCount int32  `gorm:"type:INT;not null;default:0" json:"favorite_count" validate:""`
	CommentCount  int32  `gorm:"type:INT;not null;default:0" json:"comment_count" validate:""`
	PlayURL       string `gorm:"type:varchar(100);not null" json:"play_url" validate:""`
	CoverURL      string `gorm:"type:varchar(100);not null" json:"cover_url" validate:""`
}

// VideoAPI 主要提供给查询操作使用
type VideoAPI struct {
	Author        UserAPI `json:"author"`
	VideoID       int64   `json:"id"`
	FavoriteCount int32   `json:"favorite_count"`
	CommentCount  int32   `json:"comment_count"`
	PlayURL       string  `json:"play_url"`
	CoverURL      string  `json:"cover_url"`
	IsFavorite    bool    `json:"is_favorite"`
}

// Favorite 点赞实体
type Favorite struct {
	gorm.Model
	UserID  int64 `gorm:"type:BIGINT;not null" json:"user_id" validate:""`
	VideoID int64 `gorm:"type:BIGINT;not null" json:"video_id" validate:""`
}

// Comments  实体
type Comments struct {
	gorm.Model
	UserID  int64  `gorm:"type:BIGINT;not null;" json:"user_id" validate:""`
	VideoID int64  `gorm:"type:BIGINT;not null;" json:"video_id" validate:""`
	Content string `gorm:"type:varchar(100);not null;" json:"content"`
}

// Follow 关注实体
type Follow struct {
	gorm.Model
	UserID   int64 `gorm:"type:BIGINT;not null;" json:"user_id" validate:""`
	ToUserID int64 `gorm:"type:BIGINT;not null;" json:"to_user_id" validate:""`
}
