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
// 初期使用第三方库，后期优化，可以考虑自己本地实现一个支持并发的自增ID生成器，张雷老师代码里面提供了一种方法
// atomic.AddInt64(p *int64,delta int64) 这类方法
func init() {
	var options = idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)

}

// UserLogin 调用 dao层 服务获取加密后的密文，通过 bcrypt.CompareHashAndPassword() 进行比较
// 如果匹配，则生成对应的 token
func UserLogin(name, password string) (response.Login, error) {
	fields := []string{"pass_word", "user_id"}
	u, err := dao.UserLoginDAO.Query(map[string]interface{}{
		"name": name,
	}, fields)

	// 创建处理错误的匿名函数
	handleError := func(errorType *errno.Errno) response.Login {
		return response.Login{
			Status: response.Status{Code: errorType.Code, Message: errorType.Message},
		}
	}

	if err != nil {
		log.Print(err)
		return handleError(errno.ErrDataBase), err
	}

	//和数据库中加密的密码密文进行比较，如果密码正确 err 为空，否则给出错误
	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	if err != nil {
		return handleError(errno.ErrPassWordWrong), err
	}

	// 为登录用户设置 token
	token, err := middleware.SetUpToken(u.UserID)
	if err != nil {
		return handleError(errno.ErrTokenSetUpFail), err
	}
	// 所有处理都成功，返回对应的数据
	return response.Login{
		Status: response.Status{Code: 0, Message: "success"},
		UserID: u.UserID,
		Token:  token,
	}, nil
}

// Register 用户注册
// 首先为用户生成全局唯一ID，调用 dao 层服务创建User和UserLogin实体在数据库里面的记录
// 在dao层创建 UserLogin 实体的数据库记录之前， 还会对密码进行加密存储
// dao 层服务完成并且没有错误时，生成 token
func Register(name, password string) (response.Register, error) {
	UserID := idgen.NextId()

	//User 实体创建实例
	err := dao.UserDAO.Create(
		map[string]interface{}{
			"UserID": UserID,
		})

	// 处理错误
	handleError := func(errorType *errno.Errno) response.Register {
		return response.Register{
			Status: response.Status{Code: errorType.Code, Message: errorType.Message},
		}
	}

	//处理错误
	if err != nil {
		return handleError(errno.ErrCreateUserFail), err
	}

	password, err = encryptPassWord(password) //密码加密
	if err != nil {
		return handleError(errno.ErrEncryptPassWordFail), err
	}
	// UserLogin 实体创建实例
	err = dao.UserLoginDAO.Create(
		map[string]interface{}{
			"Name":     name,
			"PassWord": password,
			"UserID":   UserID,
		})

	if err != nil {
		return handleError(errno.ErrCreateUserLoginFail), err
	}
	token, err := middleware.SetUpToken(UserID)
	if err != nil {
		return handleError(errno.ErrTokenSetUpFail), err
	}

	return response.Register{
		Status: response.Status{Code: 0, Message: "success"},
		UserID: UserID, Token: token}, nil
}

// UserInfo 获取用户信息
// 包括 user_id,name,follow_count,follower_count,is_favorite
// 最后一个字段应该和具体业务有关（我暂时还不太理解）
func UserInfo(id int64) (response.UserInfo, error) {
	// 处理错误
	handleError := func(errorType *errno.Errno) response.UserInfo {
		return response.UserInfo{
			Status: response.Status{Code: errorType.Code, Message: errorType.Message},
		}
	}

	// 获取到 user_id,follow_count,follower_count
	u, err := dao.UserDAO.Query(map[string]interface{}{"user_id": id})
	if err != nil {
		return handleError(errno.ErrQueryUserInfoFail), err
	}

	// 获取 user_name
	name, err := getUserName(id)
	if err != nil {
		return handleError(errno.ErrQueryUserNameFail), err
	}
	return response.UserInfo{
		Status: response.Status{Code: 0, Message: "success"},
		User:   model.UserAPI{ID: u.UserID, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount, Name: name},
	}, nil
}

// 密码加密
func encryptPassWord(password string) (string, error) {

	p, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptCost)
	if err != nil {
		return "", errors.New("创建用户失败")
	}
	return string(p), nil
}

// 操作 dao层获取 user_name
func getUserName(userid int64) (string, error) {
	u, err := dao.UserLoginDAO.Query(map[string]interface{}{"user_id": userid}, []string{"name"})
	if err != nil {
		return "", err
	}

	return u.Name, nil
}
