package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"douyin/model"
	"time"
)

func getNextTime(v []model.Video) int64 {
	if len(v) == 0 {
		return time.Now().UnixMilli()
	}
	return v[len(v)-1].CreateAt.UnixMilli()
}

// Feed 获取视频列表,latestTime 限制视频列表中的最新的视频时间
// 也就是获取视频应该是由近及远的
// 首先从DAO 层回去对应 的Video 的切片，之后进行 VideoAPI 的组装
// 之后遍历 VideoAPI 切片，查询作者信息(这里并没有使用数据表的联结)
func Feed(latestTime string, userID int64) (response.Feed, error) {
	handleErr := func(errType *errno.Errno) response.Feed {
		return response.Feed{
			Status: response.Status{
				Code:    errType.Code,
				Message: errType.Message,
			},
		}
	}

	videos, err := dao.VideoDAO.QueryLatest(latestTime)
	//nextTime, err := time.Parse(time.RFC3339, videos[len(videos)-1].CreateAt)
	nextTime := getNextTime(videos)
	if err != nil {
		return handleErr(errno.ErrQueryVideosFail), err
	}
	v := newVideoAPIList(videos)
	for i := 0; i < len(v); i++ {
		//查询视频作者信息
		resp, err := UserInfo(videos[i].AuthorID)
		if err != nil {
			return handleErr(errno.ErrQueryUserInfoFail), err
		}
		v[i].Author = resp.User //作者信息
		v[i].IsFavorite = dao.FavoriteDAO.IsFavorite(userID, v[i].VideoID)
		v[i].Author.IsFollow = dao.FollowDAO.IsFollow(userID, videos[i].AuthorID)
	}
	return response.Feed{VideoLists: v, Status: response.OK, NextTime: nextTime}, nil
}
