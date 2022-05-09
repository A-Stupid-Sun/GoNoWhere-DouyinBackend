package router

import (
	_ "douyin/controller/api/v1"
	v1 "douyin/controller/api/v1"
	"github.com/gin-gonic/gin"
	_ "time"
)

// InitRouter 初始化路由
func InitRouter() (*gin.Engine, error) {
	r := gin.Default()

	// 测试
	r.GET("/test", v1.Ping)

	r.GET("/douyin/feed/", v1.FeedController.Feed)

}
