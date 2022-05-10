package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"douyin/middleware"
	"github.com/yitter/idgenerator-go/idgen"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func init() {
	var options = idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)
}

// UserLogin 返回对应Login
func UserLogin(name, password string) (response.LoginResp, error) {
	u, err := dao.UserLoginDAO.Query(map[string]interface{}{
		"UserName": name,
	})
	if err != nil {
		log.Print(err)
		return response.LoginResp{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	if err != nil {
		return response.LoginResp{
			Status: response.Status{Code: errno.ErrPassWordWrong.Code, Message: errno.ErrPassWordWrong.Message},
		}, err
	}

	token, err := middleware.SetUpToken(u.UserID)
	if err != nil {
		return response.LoginResp{
			Status: response.Status{
				Code:    errno.ErrTokenSetUpFail.Code,
				Message: errno.ErrTokenSetUpFail.Message,
			}}, err
	}
	return response.LoginResp{
		Status: response.Status{Code: 0, Message: "success"},
		UserID: u.UserID,
		Token:  token,
	}, nil
}

// Register 用户注册
func Register(name, password string) (response.RegisterResp, error) {
	UserID := idgen.NextId()

	err := dao.UserDAO.Create(
		map[string]interface{}{
			"Name":   name,
			"UserID": UserID,
		})

	log.Print("_______________")
	if err != nil {
		return response.RegisterResp{
			Status: response.Status{
				Code:    errno.ErrCreateUserFail.Code,
				Message: errno.ErrCreateUserFail.Message,
			}}, err
	}
	err = dao.UserLoginDAO.Create(
		map[string]interface{}{
			"Name":     name,
			"PassWord": password,
			"UserID":   UserID,
		})

	if err != nil {
		return response.RegisterResp{Status: response.Status{
			Code:    errno.ErrCreateUserLoginFail.Code,
			Message: errno.ErrCreateUserLoginFail.Message,
		}}, err
	}
	token, err := middleware.SetUpToken(UserID)
	if err != nil {
		return response.RegisterResp{
			Status: response.Status{
				Code:    errno.ErrQueryUserLoginFail.Code,
				Message: errno.ErrQueryUserLoginFail.Message,
			}}, err
	}

	return response.RegisterResp{
		Status: response.Status{Code: 0, Message: "success"},
		UserID: UserID, Token: token}, nil
}

// UserInfo 用户信息
func UserInfo(id int64) (response.UserInfoResp, error) {
	u, err := dao.UserDAO.Query(map[string]interface{}{"UserID": id})
	if err != nil {
		return response.UserInfoResp{
			Status: response.Status{
				Code:    errno.ErrQueryUserInfoFail.Code,
				Message: errno.ErrQueryUserInfoFail.Message,
			}}, err
	}

	return response.UserInfoResp{
		Status: response.Status{Code: 0, Message: "success"},
		User:   *u,
	}, err
}
