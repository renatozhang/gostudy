package util

const (
	ErrCodeSuccess           = 0
	ErrCodeParmeter          = 1001
	ErrCodeUserExist         = 1002
	ErrCodeServerBusy        = 1003
	ErrCodeUserNotExists     = 1004
	ErrCodeUserPasswordWrong = 1005
	ErrCodeCaptionHit        = 1006
	ErrCodeContentHit        = 1007
	ErrCodeNotLogin          = 1008
)

func GetMessage(code int) (message string) {
	switch code {
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeParmeter:
		message = "参数错误"
	case ErrCodeUserExist:
		message = "用户名已经存在"
	case ErrCodeServerBusy:
		message = "服务器繁忙"
	case ErrCodeUserNotExists:
		message = "用户不存在"
	case ErrCodeUserPasswordWrong:
		message = "用户名或密码不正确"
	case ErrCodeCaptionHit:
		message = "标题中含有非法内容，请修改后发表"
	case ErrCodeContentHit:
		message = "内容中含有非法内容，请修改后发表"
	case ErrCodeNotLogin:
		message = "用户未登录"
	default:
		message = "未知错误"
	}
	return
}
