package response

import (
	"douyin/errno"
	"douyin/model"
)

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

var OK = Status{Code: 0, Message: "success"}
var NoToken = Status{Code: errno.ErrNoToken.Code, Message: errno.ErrNoToken.Message}
var TokenExpired = Status{Code: errno.ErrTokenExpired.Code, Message: errno.ErrTokenExpired.Message}
var InvalidParma = Status{Code: errno.ErrValidateFail.Code, Message: errno.ErrValidateFail.Message}

type Feed struct {
	Status
	NextTime   int64            `json:"next_time"`
	VideoLists []model.VideoAPI `json:"video_list"`
}

type VideoList struct {
	Status     `json:"status"`
	VideoLists []model.VideoAPI `json:"video_list"`
}

//type PublishList struct {
//	Status     `json:"status"`
//	VideoLists []model.VideoAPI `json:"video_list"`
//}
//
//type FavoriteList struct {
//	Status     `json:"status"`
//	VideoLists []model.VideoAPI `json:"video_list"`
//}

type FavoriteList VideoList
type PublishList VideoList

type CommentList struct {
	Status
	CommentLists []model.CommentAPI `json:"comment_lists"`
}

type UserList struct {
	Status
	UserList []model.UserAPI `json:"user_list"`
}
