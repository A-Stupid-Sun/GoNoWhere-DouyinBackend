package v1

import (
	"douyin/controller/api/v1/response"
	"douyin/errno"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type favoriteController struct{}

// Action 点赞，即赞操作,取消或者增加
func (f *favoriteController) Action(c *gin.Context) {
	// TODO 目前我想到操作数据库时，可能出现的问题就是，在并发情况下
	// 如，请求一读到点赞数是100，然后执行 update+1，点赞数为101，刷新到数据库；
	// 但是同时请求二也读到点赞数100，执行update+1，刷新到数据库，实际点赞数应该是102，但是结果确实101
	// 其实就是并发情况下的资源竞态问题，加锁即可（但是性能影响极大，300个并发请求就会导致大部分请求 4000多ms的延迟）
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	var actionType bool = true
	if c.Query("action_type") == "2" {
		actionType = false
	}
	res, err := service.FavoriteAction(userID, videoID, actionType)
	if err != nil {
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrFavoriteFail.Code,
			Message: errno.ErrFavoriteFail.Message,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// List 用户点赞列表
func (f *favoriteController) List(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.InvalidParma)
		return
	}
	res, err := service.FavoriteList(userID)
	if err != nil {
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrFavoriteVideoIDListFail.Code,
			Message: errno.ErrFavoriteVideoIDListFail.Message})
		return
	}
	c.JSON(http.StatusOK, res)
}
