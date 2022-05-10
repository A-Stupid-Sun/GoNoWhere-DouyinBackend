package response

import (
	"douyin/model"
	"time"
)

type Status struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_msg"`
}

type FeedResp struct {
	Status
	NextTime time.Time
	Videos   []model.VideoAPI `json:"video_list"`
}

type UserInfoResp struct {
	Status
	model.User
}

type LoginResp struct {
	Status
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
type RegisterResp struct {
	Status
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type PublishResp struct {
	Status
	Videos     []model.VideoAPI
	IsFavorite bool
}

type FavoriteResp struct {
	Status
	Videos []model.VideoAPI `json:"omitempty"`
}
