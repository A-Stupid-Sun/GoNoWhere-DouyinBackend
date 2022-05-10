package v1

import (
	"douyin/controller/api/v1/response"
	"douyin/errno"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userController struct{}

// Login 用户登录入口
func (u *userController) Login(c *gin.Context) {
	n := c.Query("username")
	p := c.Query("password")
	if len(n) < 0 || len(n) > 32 || len(p) < 0 || len(p) > 32 {
		// TODO 尝试使用自定义状态码
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrQueryPramsInvalid.Code,
			Message: errno.ErrQueryPramsInvalid.Message,
		})
	}
	res, err := service.UserLogin(n, p)
	if err != nil {
		c.JSON(http.StatusOK, res.Status) //失败了就没有必要返回实际数据了
	}

	c.JSON(http.StatusOK, res)
}

// Register 用户注册
func (u *userController) Register(c *gin.Context) {
	n := c.Query("username")
	p := c.Query("password")
	if len(n) < 0 || len(n) > 32 || len(p) < 0 || len(p) > 32 {
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrQueryPramsInvalid.Code,
			Message: errno.ErrQueryPramsInvalid.Message,
		})
	}

	res, err := service.Register(n, p)

	if err != nil {
		c.JSON(http.StatusOK, res.Status)
	}
	c.JSON(http.StatusOK, res)
}

// Info 用户个人信息,id,name,关注数量，粉丝数量
func (u *userController) Info(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrQueryPramsInvalid.Code,
			Message: errno.ErrQueryPramsInvalid.Message,
		})
	}
	res, err := service.UserInfo(id)
	c.JSON(http.StatusOK, res)
}

//...
