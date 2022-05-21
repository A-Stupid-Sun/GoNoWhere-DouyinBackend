package v1

import (
	"douyin/model"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type relationController struct{}

// Action 关注操作，关注、取关
func (r *relationController) Action(c *gin.Context) {
	if c.Query("action_type") == "1" {
		add(c)
	} else {
		cancel(c)
	}
}

// FollowList 用户的关注列表
func (r *relationController) FollowList(c *gin.Context) {

}

// FollowerList 用户的粉丝列表
func (r *relationController) FollowerList(c *gin.Context) {

}

func add(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	resp, _ := service.FollowAdd(&model.Follow{UserID: userID, ToUserID: toUserId})
	c.JSON(http.StatusOK, resp)
}

func cancel(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)

	resp, _ := service.FollowCancel(&model.Follow{UserID: userID, ToUserID: toUserId})
	c.JSON(http.StatusOK, resp)

}
