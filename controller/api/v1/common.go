package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	FeedController     = &feedController{}
	UserController     = &userController{}
	CommentController  = &commentController{}
	RelationController = relationController{}
	FavoriteController = &favoriteController{}
	PublishController  = &publishController{}
)

// Ping 测试服务器联通
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Msg": "OK",
	})
}
