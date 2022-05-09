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
	//实际上 Videos 仅仅用来表示Video和User之间的关系，并且自动迁移的时候，设置关闭物理外键，就不会创建物理外键，使用逻辑外键
	//经过实测啦
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
	PassWord string `gorm:"type:varchar(65);not null;unique" json:"password" validate:"min=6,max=32"`
}

// Video 实体
// 发布时间要创建索引，加速按照时间访问
type Video struct {
	id            int64  `gorm:"type:INT;primaryKey;autoIncrement"`
	VideoID       int64  `gorm:"type:BIGINT;not null;UNIQUE" json:"video_id" validate:""`
	AuthorID      int64  `gorm:"type:BIGINT;not null" json:"author_id" validate:""`
	FavoriteCount int32  `gorm:"type:INT;not null;default:0" json:"favorite_count" validate:""`
	CommentCount  int32  `gorm:"type:INT;not null;default:0" json:"comment_count" validate:""`
	PlayURL       string `gorm:"type:varchar(100);not null" json:"play_url" validate:""`
	CoverURL      string `gorm:"type:varchar(100);not null" json:"cover_url" validate:""`
}
type VideoAPI struct {
	Author        UserAPI
	VideoID       int64  `json:"id"`
	FavoriteCount int32  `json:"favorite_count"`
	CommentCount  int32  `json:"comment_count"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	IsFavorite    bool   `json:"is_favorite"`
}

// Favorite 点赞实体
type Favorite struct {
	id       int64  `gorm:"type:INT;primaryKey;autoIncrement"`
	UserID   int64  `gorm:"type:BIGINT;not null" json:"user_id" validate:""`
	VideoID  int64  `gorm:"type:BIGINT;not null;" json:"video_id" validate:""`
	CreateAt string `gorm:"type:timestamp;not null;default:current_timestamp()"`
}

// Comments  实体
type Comments struct {
	id       int64  `gorm:"type:INT;primaryKey;autoIncrement"`
	UserID   int64  `gorm:"type:BIGINT;not null;" json:"user_id" validate:""`
	VideoID  int64  `gorm:"type:BIGINT;not null;" json:"video_id" validate:""`
	Content  string `gorm:"type:varchar(100);not null;" json:"content"`
	CreateAt string `gorm:"type:timestamp;not null;default:current_timestamp()"`
}

// Follow 关注实体
type Follow struct {
	id       int64  `gorm:"type:INT;primaryKey;autoIncrement"`
	UserID   int64  `gorm:"type:BIGINT;not null;" json:"user_id" validate:""`
	ToUserID int64  `gorm:"type:BIGINT;not null;" json:"to_user_id" validate:""`
	CreateAt string `gorm:"type:timestamp;not null;default:current_timestamp()"`
}
