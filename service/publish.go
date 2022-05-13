package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"douyin/utils/upload"
	"fmt"
	"mime/multipart"
)

// PublishVideo 上传数据到七牛云对象存储,并在数据库里面更新数据
// 如果以上两个步骤任何一个出错，都执行回滚操作,确保上述操作是原子操作
// 要么全部成功，要么全部失败
var coverURL = "https://cdn.cnbj1.fds.api.mi-img.com/product-images/redmik40sb4bd68/31.jpg"

func PublishVideo(file multipart.File, header *multipart.FileHeader, userID int64) (response.Status, error) {
	handleErr := func(errorType *errno.Errno) response.Status {
		return response.Status{Code: errorType.Code, Message: errorType.Message}
	}
	url, err := upload.ToQiNiu(file, header.Size)
	if err != nil {
		fmt.Println(err)
		return handleErr(errno.ErrUpLoadToQiNiuFail), err
	}
	// TODO 暂未实现封面照片相关
	err = dao.VideoDAO.Create(map[string]interface{}{"author_id": userID, "play_url": url, "cover_url": coverURL, "video_id": 4123})
	if err != nil {
		return handleErr(errno.ErrCreateVideoRecordFail), err
	}
	return response.StatusOK, nil
}
