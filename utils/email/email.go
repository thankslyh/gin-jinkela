package email

import (
	"errors"
	em "github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func SendEmail(text, subject string, to []string) error {
	e := em.NewEmail()
	e.From = "测试 <mesterli@126.com>"
	e.To = to
	e.Subject = subject
	e.Text = []byte(text)
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "mesterli@126.com", "LPYGPJXNBAUUJUFQ", "smtp.126.com"))
	if err != nil {
		log.Fatal("发送失败\n", err.Error())
		return errors.New("发送失败" + err.Error())
	}
	log.Println("发送成功........")
	return nil
}

func CheckEmail(email string) {

}
