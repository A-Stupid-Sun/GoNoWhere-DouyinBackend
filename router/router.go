package router

import (
	_ "douyin/controller/api/v1"
	v1 "douyin/controller/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "time"
)

// InitRouter 初始化路由
func InitRouter() (*gin.Engine, error) {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Msg": "ok",
		})
	})

	r.GET("/douyin/feed/", v1.FeedController.Feed)

}
