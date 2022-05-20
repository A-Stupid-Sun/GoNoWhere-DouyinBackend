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

	// 需要鉴权token
	auth := r.Group("", middleware.JWTToken())
	{
		auth.POST("/douyin/publish/action/", v1.PublishController.Publish)  //用户投稿
		auth.GET("/douyin/publish/list/", v1.PublishController.PublishList) //发布列表
		auth.GET("/douyin/user/", v1.UserController.Info)                   //用户信息
		auth.POST("/douyin/favorite/action/", v1.FavoriteController.Action) //赞操作
		auth.GET("/douyin/favorite/list/", v1.FavoriteController.List)      //点赞列表
	}

	return r
}
