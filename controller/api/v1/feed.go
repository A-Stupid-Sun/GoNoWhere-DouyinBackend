package v1

import "github.com/gin-gonic/gin"

type FeedController struct{}

// Feed 推送视频流到客户端，按照视频的投稿时间倒序，即由近及远
func (f *FeedController) Feed(c *gin.Context) {

}
