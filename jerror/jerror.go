package jerror

import "errors"

var (
	EmailFormatError = errors.New("手机号格式错误")
	EmailAlreadyExsit = errors.New("手机号已存在")
	VerifyCodeError = errors.New("验证码错误")
	PasswordError = errors.New("密码错误！")
	TagExist = errors.New("该标签已存在！")
)