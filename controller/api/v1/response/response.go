package response

import "douyin/model"

type Status struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_msg"`
}

type UserInfo struct {
	Status
	User model.UserAPI
}

type Login struct {
	Status
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
type Register struct {
	Status
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

//type Feed struct {
//	Status
//	NextTime time.Time
//	Videos   []model.VideoAPI `json:"video_list"`
//}

//type Publish struct {
//	Status
//	Videos     []model.VideoAPI
//	IsFavorite bool
//}
//
//type Favorite struct {
//	Status
//	Videos []model.VideoAPI `json:"omitempty"`
//}
