package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 目前是手动创建一个Controller实体，接下来可能会改成第二节中赵征老师那样的Go风格的单例模式
// 摸钱手动创建的，扩展性不是很好
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
