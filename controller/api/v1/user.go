package v1

import (
	"douyin/controller/api/v1/response"
	"douyin/errno"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct{}

// Login 用户登录入口,参数验证，调用 service 层服务,
func (u *userController) Login(c *gin.Context) {
	n := c.Query("username")
	p := c.Query("password")

	if len(n) == 0 || len(n) > 32 || len(p) == 0 || len(p) > 32 {
		// 尝试使用自定义状态码
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrQueryPramsInvalid.Code,
			Message: errno.ErrQueryPramsInvalid.Message,
		})
	}

	res, err := service.UserLogin(n, p)
	if err != nil {
		c.JSON(http.StatusOK, res.Status) //如果就没有必要返回实际数据了，只返回状态码和信息即可
		return
	}

	c.JSON(http.StatusOK, res)
}

// Register 用户注册，如果用户名和密码长度不符合规范，直接返回错误，不执行后续操作
// 合法情况下，调用 service 层对应服务
func (u *userController) Register(c *gin.Context) {
	n := c.Query("username")
	p := c.Query("password")

	if len(n) == 0 || len(n) > 32 || len(p) == 0 || len(p) > 32 {
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrQueryPramsInvalid.Code,
			Message: errno.ErrQueryPramsInvalid.Message,
		})
	}

	res, err := service.Register(n, p)
	if err != nil {
		c.JSON(http.StatusOK, res.Status)
		return
	}
	c.JSON(http.StatusOK, res)
}

// Info 获取用户个人信息,如：id,name,关注数量，粉丝数量,如果请求出错（参数不合法、service 层出错），只返回状态码和信息
func (u *userController) Info(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.Status{
			Code:    errno.ErrQueryPramsInvalid.Code,
			Message: errno.ErrQueryPramsInvalid.Message,
		})
	}

	// 调用 service 层服务
	res, err := service.UserInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, res.Status)
		return
	}

	c.JSON(http.StatusOK, res)
}

//...
