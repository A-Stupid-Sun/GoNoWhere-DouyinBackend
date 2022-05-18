package router

import (
	v1 "douyin/controller/api/v1"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	// 测试
	r.GET("/test/", v1.Ping)
	r.GET("/douyin/feed/", v1.FeedController.Feed)               //视频流接口
	r.POST("/douyin/user/register/", v1.UserController.Register) //用户注册
	r.POST("/douyin/user/login/", v1.UserController.Login)       //用户登录

	authorization := r.Group("", middleware.JWTToken())
	{
		// 需要鉴权token
		authorization.POST("/douyin/publish/action/", v1.PublishController.Publish)  //用户投稿
		authorization.GET("/douyin/publish/list/", v1.PublishController.PublishList) //发布列表
		authorization.GET("/douyin/user/", v1.UserController.Info)                   //用户信息

	}

	return r
}
