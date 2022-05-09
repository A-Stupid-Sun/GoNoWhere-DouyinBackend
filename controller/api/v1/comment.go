package v1

import "github.com/gin-gonic/gin"

type commentController struct{}

// Action 执行增加、删除评论操作 (注意形参命名!!!)
func (cc *commentController) Action(ctx *gin.Context) {

}

// List 某视频的所有评论，(按时间倒序！！！)
func (cc *commentController) List(ctx *gin.Context) {

}
