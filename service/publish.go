package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"douyin/model"
	"douyin/utils/upload"
	"fmt"
	"mime/multipart"

	"github.com/yitter/idgenerator-go/idgen"
)

// PublishVideo 上传数据到七牛云对象存储,并在数据库里面更新数据
// 如果以上两个步骤任何一个出错，都执行回滚操作,确保上述操作是原子操作
// 要么全部成功，要么全部失败
var coverURL = "https://cdn.cnbj1.fds.api.mi-img.com/product-images/redmik40sb4bd68/31.jpg"

func PublishVideo(file multipart.File, header *multipart.FileHeader, userID int64) (response.Status, error) {
	handleErr := func(errorType *errno.Errno) response.Status {
		return response.Status{Code: errorType.Code, Message: errorType.Message}
	}

	// 生成视频id
	id := idgen.NextId() // 暂时和用户ID使用同一个id生成器
	url, err := upload.ToQiNiu(file, header.Size, id)
	if err != nil {
		fmt.Println(err)
		return handleErr(errno.ErrUpLoadToQiNiuFail), err
	}
	// TODO 暂未实现封面照片相关

	err = dao.VideoDAO.Create(map[string]interface{}{
		"author_id": userID,
		"play_url":  url,
		"cover_url": coverURL,
		"video_id":  id})
	if err != nil {
		return handleErr(errno.ErrCreateVideoRecordFail), err
	}
	return response.StatusOK, nil
}

// PublishList 返回用户发布的所有的视频，包括视频的点赞数和评论数等视频相关信息
func PublishList(userID int64) (response.PublishList, error) {
	handleErr := func(errorType *errno.Errno) response.PublishList {
		return response.PublishList{
			Status: response.Status{
				Code:    errorType.Code,
				Message: errorType.Message,
			}}
	}
	// 首先查询视频
	videos, err := dao.VideoDAO.Query(
		map[string]interface{}{"author_id": userID},
		"play_url", "cover_url", "favorite_count", "comment_count", "video_id")
	if err != nil {
		return handleErr(errno.ErrQueryVideosFail), err
	}

	v := newVideoAPIList(videos) //构造数据

	for i, video := range v {

		// 作者自己是否点赞
		v[i].IsFavorite = dao.FavoriteDAO.IsFavorite(userID, video.VideoID)
	}
	return response.PublishList{
		Status:     response.StatusOK,
		VideoLists: v,
	}, err
}

// 构造 VideoAPI 切片
func newVideoAPIList(videos []model.Video) []model.VideoAPI {
	var v []model.VideoAPI
	for _, i := range videos {
		v = append(v, model.VideoAPI{
			VideoID:       i.VideoID,
			PlayURL:       i.PlayURL,
			CoverURL:      i.CoverURL,
			CommentCount:  i.CommentCount,
			FavoriteCount: i.FavoriteCount,
		})
	}

	return v
}
