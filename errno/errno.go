package errno

type Errno struct {
	Code    int
	Message string
}

var OK = &Errno{Code: 0, Message: "OK"}
var Fail = &Errno{Code: -1, Message: "Error"}

var (
	// 数据库相关 101 开头
	ErrDataBase            = &Errno{Code: 10101, Message: "数据库错误"}
	ErrQueryUserInfoFail   = &Errno{Code: 10102, Message: "查询用户信息错误"}
	ErrQueryUserLoginFail  = &Errno{Code: 10103, Message: "查询用户登录信息错误"}
	ErrCreateUserFail      = &Errno{Code: 10104, Message: "创建用户信息失败"}
	ErrCreateUserLoginFail = &Errno{Code: 10105, Message: "创建用户登录信息失败"}

	// Token 相关 102 开头
	ErrTokenExpired   = &Errno{Code: 10201, Message: "Token 已过期"}
	ErrTokenSetUpFail = &Errno{Code: 10202, Message: "Token 生成失败"}
	ErrNoToken        = &Errno{Code: 10203, Message: "No Token"}

	//视频相关 103 开头
	ErrVideoUpload           = &Errno{Code: 10301, Message: "视频上传失败"}
	ErrCreateVideoRecordFail = &Errno{Code: 10302, Message: "数据库新增视频记录失败"}
	ErrQueryVideosFail       = &Errno{Code: 10303, Message: "查询视频信息失败"}

	// 用户相关 104 开头
	ErrPassWordWrong       = &Errno{Code: 10401, Message: "密码错误"}
	ErrEncryptPassWordFail = &Errno{Code: 10402, Message: "密码加密失败"}
	ErrQueryUserNameFail   = &Errno{Code: 10403, Message: "获取用户名失败"}

	// 评论相关 105 开头
	ErrCommentAddFail       = &Errno{Code: 10501, Message: "新增评论失败"}
	ErrCommentDelFail       = &Errno{Code: 10502, Message: "删除评论失败"}
	ErrQueryCommentListFail = &Errno{Code: 10503, Message: "查询评论列表失败"}

	// 点赞相关 106 开头
	ErrFavoriteAddFail         = &Errno{Code: 10601, Message: "点赞失败，请稍后再试"}
	ErrFavoriteSubFail         = &Errno{Code: 10602, Message: "取消点赞失败，请稍后再试"}
	ErrFavoriteFail            = &Errno{Code: 10606, Message: "点赞操作失败，请稍后再试"}
	ErrFavoriteVideoIDListFail = &Errno{Code: 10603, Message: "获取用户点赞视频ID列表失败"}
	// 数据验证相关 107 开头

	ErrValidateFail = &Errno{Code: 10701, Message: "数据验证失败"}

	// 请求参数相关 108 开头
	ErrQueryPramsInvalid = &Errno{Code: 10801, Message: "请求参数不合法"}

	// 上传文件相关 109 开头
	ErrUpLoadToQiNiuFail = &Errno{Code: 10901, Message: "上传七牛云失败"}
	ErrFileTooMuchBig    = &Errno{Code: 10902, Message: "文件太大"}
)
