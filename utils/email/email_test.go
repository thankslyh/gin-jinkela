package email_test

import (
	"jinkela/utils/email"
	"testing"
)

func TestEmail(t *testing.T) {
	email.SendEmail("注册验证码为：6666", "注册验证码", []string{"thankslyh@gmail.com", "anpeng1@foxmail.com"})
}
