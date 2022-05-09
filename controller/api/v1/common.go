package v1

import (
	"douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	FeedController     = &feedController{}
	UserController     = &userController{}
	CommentController  = &commentController{}
	RelationController = relationController{}
	FavoriteController = &favoriteController{}
	PublishController  = &publishController{}
)

type Status struct {
	Code    int           `json:"status_code"`
	Message string        `json:"status_msg"`
	Videos  []model.Video `json:"video_list"`
}

type FeedData struct {
	Status
	NextTime time.Time
}

type UserData struct {
	Status
	model.UserAPI
	UserID int64  `json:"user_id omitempty"`
	Token  string `json:"token omitempty"`
}

type PublishData struct {
	Status
	Videos     []model.VideoAPI
	IsFavorite bool
}

type FavoriteData struct {
	Status
	Videos []model.VideoAPI `json:"omitempty"`
}

// Ping 测试服务器联通
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Msg": "OK",
	})
}
