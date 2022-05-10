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

	//视频相关 103 开头
	ErrVideoUpload = &Errno{Code: 10301, Message: "视频上传失败"}

	// 用户相关 104 开头
	ErrPassWordWrong = &Errno{Code: 10401, Message: "密码错误"}

	// 评论相关 105 开头

	// 点赞相关 106 开头

	// 数据验证相关 107 开头

	ErrValidateFail = &Errno{Code: 10701, Message: "数据验证失败"}

	// 请求参数相关
	ErrQueryPramsInvalid = &Errno{Code: 10801, Message: "请求参数不合法"}
)
