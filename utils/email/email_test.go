package email_test

import (
	"fmt"
	"jinkela/utils/email"
	"regexp"
	"testing"
)

func TestEmail(t *testing.T) {
	email.SendEmail("注册验证码为：6666", "注册验证码", []string{"thankslyh@gmail.com", "anpeng1@foxmail.com"})
}

func TestVerifyEmail(t *testing.T)  {
	const EmailReg = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	fmt.Println(regexp.Match(EmailReg, []byte("thankslyh@gmail.com")))
}
