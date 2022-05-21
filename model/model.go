package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID       uint64    `gorm:"comment:自增主键"`
	CreateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UpdateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	gorm.DeletedAt
}

// User 实体
// 其他所有字段非空，一是业务要求，二是优化数据库性能（细节不表）
type User struct {
	ID       uint64    `gorm:"comment:自增主键"`
	CreateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	gorm.DeletedAt
	UserID        int64   `gorm:"type:bigint;unsigned;not null;unique;uniqueIndex:idx_user_id" json:"user_id"`
	FollowCount   int     `gorm:"type:int;unsigned;not null;default:0" json:"follow_count"`
	FollowerCount int     `gorm:"type:int;unsigned;not null;default:0" json:"follower_count"`
	Videos        []Video `gorm:"foreignKey:AuthorID;references:UserID" json:"omitempty"`
}

// UserAPI 主要提供给接口使用
type UserAPI struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// UserLogin （备选） 登录数据单独存放，因为登录、注册、改密操作相对来说比较少，算是冷数据吧
// 我个人觉得，把密码和 name 做一个拆分，是比较好的选择，一是加密之后的密码较长，二是大部分情况下用不到密码和用户名，很多操作只需要id即可
//
type UserLogin struct {
	ID       uint64    `gorm:"comment:自增主键"`
	CreateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UpdateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	gorm.DeletedAt

	UserID   int64  `gorm:"type:bigint;unsigned;not null;unique;uniqueIndex:idx_user_id" json:"user_id"`
	Name     string `gorm:"type:varchar(34);not null;unique;uniqueIndex:idx_user_name" json:"name" validate:"min=6,max=32"`
	PassWord string `gorm:"type:varchar(65);not null" json:"password" validate:"min=6,max=32"`
}

// Video 实体
// 发布时间要创建索引，加速按照时间访问
type Video struct {
	ID       uint64 `gorm:"comment:自增主键"`
	CreateAt string `gorm:"type:timeStamp;not null;default:current_timestamp();index:idx_create_time,sort:desc"`
	gorm.DeletedAt
	VideoID       int64  `gorm:"type:BIGINT;not null;UNIQUE" json:"video_id" validate:""`
	AuthorID      int64  `gorm:"type:BIGINT;not null;index:idx_author_id" json:"author_id" validate:""`
	FavoriteCount int32  `gorm:"type:INT;not null;default:0" json:"favorite_count" validate:""`
	CommentCount  int32  `gorm:"type:INT;not null;default:0" json:"comment_count" validate:""`
	PlayURL       string `gorm:"type:varchar(100);not null" json:"play_url" validate:""`
	CoverURL      string `gorm:"type:varchar(100);not null" json:"cover_url" validate:""`
}

// VideoAPI 主要提供给查询操作使用
type VideoAPI struct {
	VideoID       int64   `json:"id,omitempty"`
	Author        UserAPI `json:"author,omitempty"`
	FavoriteCount int32   `json:"favorite_count"`
	CommentCount  int32   `json:"comment_count"`
	PlayURL       string  `json:"play_url"`
	CoverURL      string  `json:"cover_url"`
	IsFavorite    bool    `json:"is_favorite"`
}

// Favorite 点赞实体
type Favorite struct {
	ID       uint64    `gorm:"comment:自增主键"`
	CreateAt time.Time `gorm:"type:timeStamp;not null;default:current_timestamp();comment:创建时间"`
	UserID   int64     `gorm:"type:BIGINT;not nul;index:idx_user_id;comment:点赞用户ID" json:"user_id"`
	VideoID  int64     `gorm:"type:BIGINT;not null;index:idx_video_id;comment:被点赞视频ID" json:"video_id" `
}

// Comment  实体
type Comment struct {
	ID       int64     `gorm:"comment:自增主键"`
	CreateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UpdateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	gorm.DeletedAt
	UserID  int64  `gorm:"type:BIGINT;not null;评论用户ID" json:"user_id"`
	VideoID int64  `gorm:"type:BIGINT;not null;index:idx_video_id;comment:被评论视频ID" json:"video_id" validate:""`
	Content string `gorm:"type:varchar(100);not null;comment:评论内容" json:"content"`
}

type CommentAPI struct {
	ID       int64
	User     UserAPI
	Content  string
	CreateAt string
}

// Follow 关注实体
type Follow struct {
	gorm.Model
	UserID   int64 `gorm:"type:BIGINT;not null;index:idx_user_id;comment:粉丝用户ID" json:"user_id" validate:""`
	ToUserID int64 `gorm:"type:BIGINT;not null;index:idx_to_user_id;comment:被关注用户ID" json:"to_user_id" validate:""`
}
