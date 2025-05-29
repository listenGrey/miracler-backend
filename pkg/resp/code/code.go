package code

const (
	Success       = 0
	BadRequest    = 400
	Unauthorized  = 401
	Forbidden     = 403
	NotFound      = 404
	InternalError = 500

	// 业务错误码建议用 10000+，防止和 HTTP 状态码混淆
	UserNotFound  = 10001
	PasswordWrong = 10002
)

var msg = map[int]string{
	Success:       "success",
	BadRequest:    "请求参数错误",
	Unauthorized:  "未授权",
	Forbidden:     "无权限访问",
	NotFound:      "资源不存在",
	InternalError: "服务器内部错误",

	UserNotFound:  "用户不存在",
	PasswordWrong: "密码错误",
}

func GetMsg(code int) string {
	return msg[code]
}
