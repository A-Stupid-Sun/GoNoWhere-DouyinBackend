package router

import (
	v1 "douyin/controller/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	// 测试
	r.GET("/test", v1.Ping)

	r.GET("/douyin/feed", v1.FeedController.Feed)                   //视频流接口
	r.POST("/douyin/user/register", v1.UserController.Register)     //用户注册
	r.POST("/douyin/user/login", v1.UserController.Login)           //用户登录
	r.GET("/douyin/user", v1.UserController.Info)                   //用户信息
	r.POST("/douyin/publish/action/", v1.PublishController.Publish) //投稿接口

	return r

}
