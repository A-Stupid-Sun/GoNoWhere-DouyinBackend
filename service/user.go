package service

import (
	"douyin/config"
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"douyin/middleware"
	"douyin/model"
	"errors"
	"log"

	"github.com/yitter/idgenerator-go/idgen"
	"golang.org/x/crypto/bcrypt"
)

// 全局ID 生成器，隐式初始化
func init() {
	var options = idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)
}

// UserLogin 返回对应Login
func UserLogin(name, password string) (response.Login, error) {
	fields := []string{"pass_word", "user_id"}
	u, err := dao.UserLoginDAO.Query(map[string]interface{}{
		"name": name,
	}, fields)
	if err != nil {
		log.Print(err)
		return response.Login{}, err
	}

	//和数据库中加密的密码密文进行比较，如果密码正确 err 为空，否则给出错误
	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	if err != nil {
		return response.Login{
			Status: response.Status{Code: errno.ErrPassWordWrong.Code, Message: errno.ErrPassWordWrong.Message},
		}, err
	}

	// 为登录用户设置 token
	token, err := middleware.SetUpToken(u.UserID)
	if err != nil {
		return response.Login{
			Status: response.Status{
				Code:    errno.ErrTokenSetUpFail.Code,
				Message: errno.ErrTokenSetUpFail.Message,
			}}, err
	}
	// 所有处理都成功，返回对应的数据
	return response.Login{
		Status: response.Status{Code: 0, Message: "success"},
		UserID: u.UserID,
		Token:  token,
	}, nil
}

// Register 用户注册
func Register(name, password string) (response.Register, error) {
	UserID := idgen.NextId()

	//User 实体创建实例
	err := dao.UserDAO.Create(
		map[string]interface{}{
			"UserID": UserID,
		})

	//处理错误
	if err != nil {
		return response.Register{
			Status: response.Status{
				Code:    errno.ErrCreateUserFail.Code,
				Message: errno.ErrCreateUserFail.Message,
			}}, err
	}

	password, err = encryptPassWord(password)
	if err != nil {
		return response.Register{Status: response.Status{
			Code:    errno.ErrEncryptPassWordFail.Code,
			Message: errno.ErrEncryptPassWordFail.Message,
		}}, err
	}
	// UserLogin 实体创建实例
	err = dao.UserLoginDAO.Create(
		map[string]interface{}{
			"Name":     name,
			"PassWord": password,
			"UserID":   UserID,
		})

	if err != nil {
		return response.Register{Status: response.Status{
			Code:    errno.ErrCreateUserLoginFail.Code,
			Message: errno.ErrCreateUserLoginFail.Message,
		}}, err
	}
	token, err := middleware.SetUpToken(UserID)
	if err != nil {
		return response.Register{
			Status: response.Status{
				Code:    errno.ErrQueryUserLoginFail.Code,
				Message: errno.ErrQueryUserLoginFail.Message,
			}}, err
	}

	return response.Register{
		Status: response.Status{Code: 0, Message: "success"},
		UserID: UserID, Token: token}, nil
}

// UserInfo 用户信息
func UserInfo(id int64) (response.UserInfo, error) {
	// 获取到 user_id,follow_count,follower_count
	u, err := dao.UserDAO.Query(map[string]interface{}{"user_id": id})
	if err != nil {
		return response.UserInfo{
			Status: response.Status{
				Code:    errno.ErrQueryUserInfoFail.Code,
				Message: errno.ErrQueryUserInfoFail.Message,
			}}, err
	}

	name, err := getUserName(id)
	if err != nil {
		return response.UserInfo{
			Status: response.Status{
				Code:    errno.ErrQueryUserNameFail.Code,
				Message: errno.ErrQueryUserNameFail.Message,
			}}, err
	}
	return response.UserInfo{
		Status: response.Status{Code: 0, Message: "success"},
		User:   model.UserAPI{ID: u.UserID, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount, Name: name},
	}, nil
}

func encryptPassWord(password string) (string, error) {

	p, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptCost)
	if err != nil {
		return "", errors.New("创建用户失败")
	}
	return string(p), nil
}

func getUserName(userid int64) (string, error) {
	u, err := dao.UserLoginDAO.Query(map[string]interface{}{"user_id": userid}, []string{"name"})
	if err != nil {
		return "", err
	}

	return u.Name, nil
}
