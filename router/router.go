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

	r.GET("/douyin/feed", v1.FeedController.Feed)
	r.POST("/douyin/user/register", v1.UserController.Register)
	r.POST("/douyin/user/login", v1.UserController.Login)
	r.GET("/douyin/user", v1.UserController.Info)

	return r

}
