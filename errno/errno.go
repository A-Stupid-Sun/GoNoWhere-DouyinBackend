package errno

type Errno struct {
	Code    int
	Message string
}

var OK = &Errno{Code: 0, Message: "OK"}
var Fail = &Errno{Code: -1, Message: "Error"}

var (
	// 数据库相关 101 开头

	ErrDataBase = &Errno{Code: 10101, Message: "数据库错误"}

	// Token 相关 102 开头

	ErrTokenExpired = &Errno{Code: 10201, Message: "Token 已过期"}

	//视频相关 103 开头
	ErrVideoUpload = &Errno{Code: 10301, Message: "视频上传失败"}

	// 用户相关 104 开头
	ErrPassWordWrong = &Errno{Code: 10401, Message: "密码错误"}

	// 评论相关 105 开头

	// 点赞相关 106 开头

	// 数据验证相关 107 开头

	ErrValidateFail = &Errno{Code: 10701, Message: "数据验证失败"}
)
