package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
)

// FavoriteAction 点赞操作
// actionType 为 false 取消点赞
// actionType 为 true  点赞
// 新增或删除点赞记录
// 更改视频表中的 favorite_count 字段值
func FavoriteAction(userID, videoID int64, actionType bool) (response.Status, error) {
	handleErr := func(errType *errno.Errno) response.Status {
		return response.Status{
			Code:    errType.Code,
			Message: errType.Message,
		}
	}
	if actionType {
		if err := favoriteAdd(userID, videoID); err != nil {
			return handleErr(errno.ErrFavoriteAddFail), err
		}
		return response.StatusOK, nil
	}
	if err := favoriteSub(userID, videoID); err != nil {
		return handleErr(errno.ErrFavoriteAddFail), err
	}
	return response.StatusOK, nil
}

// 增加点赞操作
// 后期可以考虑加上单机缓存& 并发执行两个操作
func favoriteAdd(userID, videoID int64) error {
	err := dao.FavoriteDAO.Add(userID, videoID) //增加一条点赞记录
	if err != nil {
		return err
	}
	err = dao.VideoDAO.AddFavorite(videoID, 1)
	if err != nil {
		return err
	}
	return nil
}

// 取消点赞操作
// 后期可以考虑加上单机缓存& 并发执行两个操作
func favoriteSub(userID, videoID int64) error {
	err := dao.FavoriteDAO.Sub(userID, videoID) //增加一条点赞记录
	if err != nil {
		return err
	}
	err = dao.VideoDAO.SubFavorite(videoID, 1)
	if err != nil {
		return err
	}
	return nil
}
