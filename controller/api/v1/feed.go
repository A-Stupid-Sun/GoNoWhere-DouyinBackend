package v1

import (
	"douyin/controller/api/v1/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type feedController struct{}

// Feed 推送视频流到客户端，按照视频的投稿时间倒序，即由近及远
func (*feedController) Feed(c *gin.Context) {

	data := response.Status{}

	// 处理和获取数据
	c.JSON(http.StatusOK, data)
}
