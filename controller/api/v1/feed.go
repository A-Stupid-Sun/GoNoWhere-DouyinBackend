package v1

import (
	"douyin/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type feedController struct{}

// Feed 推送视频流到客户端，按照视频的投稿时间倒序，即由近及远
func (*feedController) Feed(c *gin.Context) {
	latestTime, ok := c.GetQuery("latest_time")
	if !ok {
		latestTime = time.Now().Format("2006-01-02 15:04:05")
	}
	fmt.Println(latestTime)
	resp, err := service.Feed(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, resp.Status)
		return
	}

	c.JSON(http.StatusOK, resp)
}
