package model

import (
	"gorm.io/gorm"
	"time"
)

// User 实体
// gorm.model 包含自增主键id、以及创建、更新、删除时间四项，详情见文档 https://gorm.io/zh_CN/docs/models.html#gorm-Model
// 其他所有字段非空，一是业务要求，二是优化数据库性能（细节不表）
//
type User struct {
	gorm.Model
	UserID        int64   `gorm:"type:BIGINT;UNSIGNED;NOT NULL;UNIQUE" json:"user_id"`
	Name          string  `gorm:"type:VARCHAR(34);NOT NULL;UNIQUE;" json:"name" validate:"min=6,max=32"`
	PassWord      string  `gorm:"type:VARCHAR(65);NOT NULL;" json:"password" validate:"min=6,max=32"`
	FollowCount   int64   `gorm:"type:INT;UNSIGNED;NOT NULL;DEFAULT:0" json:"follow_count"`
	FollowerCount int64   `gorm:"type:INT;UNSIGNED;NOT NULL;DEFAULT:0" json:"follower_count"`
	Videos        []Video `gorm:"foreignKey:AuthorID;references:UserID"`
	//实际上 Videos 仅仅用来表示Video和User之间的关系，并且自动迁移的时候，设置关闭物理外键，就不会创建物理外键，使用逻辑外键
	//经过实测啦
}

// UserLogin （备选） 登录数据单独存放，因为登录、注册、改密操作相对来说比较少，算是冷数据吧
// 提供的接口只有用户名登录，那么只包含用户名和密码即可
type UserLogin struct {
	gorm.Model
	Name     string `gorm:"type:VARCHAR(34);NOT NULL;UNIQUE" json:"name" validate:"min=6,max=32"`
	PassWord string `gorm:"type:VARCHAR(65);NOT NULL;UNIQUE" json:"password" validate:"min=6,max=32"`
}

// Video 实体
// 发布时间要创建索引，加速按照时间访问
type Video struct {
	gorm.Model
	VideoID       int64  `gorm:"type:BIGINT;NOT NULL;UNIQUE" json:"video_id" validate:""`
	AuthorID      int64  `gorm:"type:BIGINT;NOT NULL" json:"author_id" validate:""`
	FavoriteCount int32  `gorm:"type:INT;NOT NULL;DEFAULT:0" json:"favorite_count" validate:""`
	CommentCount  int32  `gorm:"type:INT;NOT NULL;DEFAULT:0" json:"comment_count" validate:""`
	PlayURL       string `gorm:"type:VARCHAR(100);NOT NULL" json:"play_url" validate:""`
	CoverURL      string `gorm:"type:VARCHAR(100);NOT NULL" json:"cover_url" validate:""`
}

// Favorite 点赞实体
type Favorite struct {
	UserID   int64
	VideoID  int64
	CreateAt string
}

// Comments  实体
// 时间是必要的
type Comments struct {
	UserID   int64
	VideoID  int64
	Content  string
	CreateAt time.Time
}

// Follow 关注实体
// 需要添加时间属性，留下扩展空间，例如关注1周年等这些
type Follow struct {
	gorm.Model
	UserID   int64
	ToUserID int64
}
