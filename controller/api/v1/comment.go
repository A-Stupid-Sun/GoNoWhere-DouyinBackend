package v1

import (
	"douyin/controller/api/v1/response"
	"douyin/model"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct{}

// Action 执行增加、删除评论操作
func (*commentController) Action(c *gin.Context) {
	// 把新增评论和取消评论的操作分开
	if c.Query("action_type") == "1" {
		actionAdd(c)
	} else {
		actionDel(c)
	}
}

// List 某视频的所有评论，(按时间倒序，也就是由近到远)
func (*commentController) List(c *gin.Context) {
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.InvalidParma)
		return
	}
	resp := service.CommentList(videoID)
	c.JSON(http.StatusOK, resp)
}

// 只处理新增评论的操作
func actionAdd(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.InvalidParma)
		return
	}
	content := c.Query("comment_text")

	comment := model.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}
	resp := service.CommentAdd(comment)
	c.JSON(http.StatusOK, resp)
}

// 只处理删除评论的操作
func actionDel(c *gin.Context) {
	commentID, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.InvalidParma)
		return
	}
	resp := service.CommentDel(commentID)
	c.JSON(http.StatusOK, resp)
}
