package erromsg

const (
	SUCCSE                     = 200
	ERRORS                     = 500
	ERROR_USERNAME_EXIST       = 5001
	ERROR_PASSWORD             = 5002
	ERROR_USERNAME_NOT_EXIST   = 5003
	ERROR_TOKEN_EXIST          = 5004
	ERROR_TOKEN_OVERDUE        = 5005
	ERROR_TOKEN                = 5006
	ERROR_TOKEN_FORMAT         = 5007
	ERROR_USER_NO_AUTH         = 5008
	ERROR_EMAIL_EXIST          = 5009
	ERROR_EMAIL_NOEXIST        = 5010
	ERROR_EMAIL_RegisteredSUCC = 5011
	ERROR_EMAIL_RegisteredFAIL = 5012
	ERROR_EMAIL_CODEFailure    = 5013
)

var msgList = map[int]string{
	SUCCSE:                     "OK",
	ERRORS:                     "FAIL",
	ERROR_USERNAME_EXIST:       "用户名已存在！",
	ERROR_PASSWORD:             "密码错误",
	ERROR_USERNAME_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:          "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_OVERDUE:        "TOKEN已过期,请重新登陆",
	ERROR_TOKEN:                "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_FORMAT:         "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_AUTH:         "该用户无权限",
	ERROR_EMAIL_EXIST:          "邮件已发送至用户绑定的邮箱，请勿重复发送",
	ERROR_EMAIL_NOEXIST:        "邮件未发送",
	ERROR_EMAIL_RegisteredSUCC: "邮箱发送成功",
	ERROR_EMAIL_RegisteredFAIL: "邮箱发送失败",
	ERROR_EMAIL_CODEFailure:    "验证码失效",
}

//获取错误提醒
func GetErrosMsg(code int) string {
	return msgList[code]
}
