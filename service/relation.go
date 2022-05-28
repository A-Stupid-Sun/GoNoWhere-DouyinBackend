package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"douyin/model"
)

func handleErr(errType *errno.Errno) response.Status {
	return response.Status{
		Code:    errType.Code,
		Message: errType.Message,
	}
}

// FollowAdd 新增关注
// User 数据库也要进行对应的操作
func FollowAdd(f *model.Follow) (response.Status, error) {
	err := dao.FollowDAO.Create(f)
	if err != nil {
		return handleErr(errno.ErrAddFollowFail), err
	}
	return response.OK, nil
}

// FollowCancel 取关
func FollowCancel(f *model.Follow) (response.Status, error) {
	err := dao.FollowDAO.Delete(f)
	if err != nil {
		return handleErr(errno.ErrCancelFollowFail), err
	}
	return response.OK, nil
}

// FollowList 某用户的关注列表
func FollowList(userID int64) (response.UserList, error) {
	follows, err := dao.FollowDAO.Query(map[string]interface{}{"user_id": userID}, "to_user_id")
	if err != nil {
		return response.UserList{}, err
	}

	var u []model.UserAPI
	for _, f := range follows {
		resp, _ := UserInfo(f.ToUserID)
		u = append(u, resp.User)
	}
	return response.UserList{
		Status:   response.OK,
		UserList: u,
	}, nil
}

// FollowerList 某用户的粉丝列表
func FollowerList(userID int64) (response.UserList, error) {
	follows, err := dao.FollowDAO.Query(map[string]interface{}{"to_user_id": userID}, "user_id")
	if err != nil {
		return response.UserList{}, err
	}

	var u []model.UserAPI
	for _, f := range follows {
		resp, _ := UserInfo(f.ToUserID)
		u = append(u, resp.User)
	}
	return response.UserList{
		Status:   response.OK,
		UserList: u,
	}, nil
}
