package jerror

import "errors"

var (
	EmailFormatError = errors.New("手机号格式错误")
	EmailAlreadyExsit = errors.New("手机号已存在")
	VerifyCodeError = errors.New("验证码错误")
)