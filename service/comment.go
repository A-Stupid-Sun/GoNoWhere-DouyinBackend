package service

import (
	"douyin/controller/api/v1/response"
	"douyin/dao"
	"douyin/errno"
	"douyin/model"
)

// CommentAdd 新增评论操作 c 里面包含了userID、videoID、content三项
// 其他的信息，如创建时间由数据库自动生成
// 成功返回 response.OK
// 失败返回对应的错误信息
func CommentAdd(c model.Comment) response.Status {
	if err := dao.CommentDAO.Add(&c); err != nil {
		return response.Status{
			Code:    errno.ErrCommentAddFail.Code,
			Message: errno.ErrCommentAddFail.Message,
		}
	}

	return response.OK
}

// CommentDel 删除评论操作，只需要提供评论id即可
// 成功返回 response.OK
// 失败返回对应的错误信息
func CommentDel(id int64) response.Status {
	if err := dao.CommentDAO.Delete(id); err != nil {
		return response.Status{
			Code:    errno.ErrCommentDelFail.Code,
			Message: errno.ErrCommentDelFail.Message,
		}
	}
	return response.OK
}

// CommentList 返回某视频的所有的评论
func CommentList(videoID int64) response.CommentList {
	handleErr := func(errType *errno.Errno) response.CommentList {
		return response.CommentList{Status: response.Status{
			Code:    errType.Code,
			Message: errType.Message,
		}}
	}
	comments, err := dao.CommentDAO.List(videoID)
	if err != nil {
		return handleErr(errno.ErrQueryCommentListFail)
	}
	c := newCommentAPIList(comments)

	return response.CommentList{Status: response.OK, CommentLists: c}
}

// 构造 commentList
func newCommentAPIList(cc []model.Comment) []model.CommentAPI {
	var c []model.CommentAPI

	for _, i := range cc {
		resp, _ := UserInfo(i.UserID)
		c = append(c, model.CommentAPI{
			ID:       i.ID,
			User:     resp.User,
			Content:  i.Content,
			CreateAt: i.CreateAt.String(),
		})
	}
	return c
}
