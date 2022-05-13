package response

import "douyin/model"

// 主要定义返回的数据结构，因为返回给前端的数据并不是定义数据模型里面的那样

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

var StatusOK = Status{Code: 0, Message: "success"}

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
