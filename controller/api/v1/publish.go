package v1

import (
	"douyin/controller/api/v1/response"
	"douyin/errno"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type publishController struct{}

// Publish 上传视频， 上传到OSS->返回视频URL->更新数据库

func handleErr(errorType *errno.Errno) response.Status {
	return response.Status{
		Code:    errorType.Code,
		Message: errorType.Message,
	}
}
func (p *publishController) Publish(c *gin.Context) {
	// 提取处理错误的出公共部分

	//TODO 上传到OSS->返回视频URL->更新数据库
	// 以上操作为原子操作，必须确保全部成功，或者全部失败
	// 错误案例1：视频已经上传到OSS，但是在更新数据库时，失败，也没有删除OSS中以及上传的文件
	// 错误案例2：视频未成功上传到OSS，返回一个不正确的url(一般是错误信息)，
	// 数据库里面更新了不完整的视频信息，也就是url不正确，返回给客户端的时候根本没法用
	file, fileHeader, _ := c.Request.FormFile("data")
	userID, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, handleErr(errno.ErrQueryPramsInvalid))
		return
	}

	// 如果视频过大100Mb
	if fileHeader.Size > (100 << 20) {
		c.JSON(http.StatusOK, handleErr(errno.ErrFileTooMuchBig))
		return
	}

	_, err = service.PublishVideo(file, fileHeader, userID)
	if err != nil {
		c.JSON(http.StatusOK, handleErr(errno.ErrVideoUpload))
		return
	}
	c.JSON(http.StatusOK, response.StatusOK) // 处理成功
}

// PublishList 用户投稿的视频列表
func (p *publishController) PublishList(c *gin.Context) {

	userID, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, handleErr(errno.ErrQueryPramsInvalid))
		return
	}
	res, err := service.PublishList(userID)
	if err != nil {
		c.JSON(http.StatusOK, handleErr(errno.ErrDataBase))
		return
	}
	c.JSON(http.StatusOK, res) //处理成功
}
