package v1

import (
	"douyin/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type feedController struct{}

// NewFeedController 是 feedController 的构造器
// 返回一个 feedController 类型的指针
func NewFeedController() *feedController {
	return &feedController{}
}

// Feed 推送视频流到客户端，按照视频的投稿时间倒序，即由近及远
// 首先根据参数获取视频列表最新时间阈值，如果此值为空则默认使用当前时间
func (*feedController) Feed(c *gin.Context) {
	latestTime := toTimeString(c.Query("latest_time"))
	resp, err := service.Feed(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, resp.Status)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// toTimeString 时间戳转换为字符串形式的格式化时间
// 参数为精确到秒的时间戳
// 如果字符串为空或者发送其他错误，那么就返回当前时间的字符串格式
func toTimeString(sec string) string {
	if len(sec) == 0 {
		return time.Now().Format("2006-01-02 15:04:05")
	}
	t, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		return time.Now().Format("2006-01-02 15:04:05")
	}

	return time.Unix(t/1000, 0).Format("2006-01-02 15:04:05")
}
