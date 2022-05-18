package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"time"
)

// Feed 获取视频列表,latestTime 限制视频列表中的最新的视频时间
// 也就是获取视频应该是由近及远的
// 首先从DAO 层回去对应 的Video 的切片，之后进行 VideoAPI 的组装
// 遍历VideoAPI 切片，查询作者信息
func Feed(latestTime string) (response.Feed, error) {
	handleErr := func(errType *errno.Errno) response.Feed {
		return response.Feed{
			Status: response.Status{
				Code:    errType.Code,
				Message: errType.Message,
			},
		}
	}

	videos, err := dao.VideoDAO.QueryLatest(latestTime)
	nextTime, err := time.Parse(time.RFC3339, videos[len(videos)-1].CreateAt)
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
	}
	return response.Feed{VideoLists: v, Status: response.StatusOK, NextTime: nextTime.Unix()}, nil
}
