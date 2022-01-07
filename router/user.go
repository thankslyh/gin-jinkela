package router

import (
	"github.com/gin-gonic/gin"
	"jinkela/api"
	"jinkela/db"
	et "jinkela/utils/email"
	"log"
	"net/http"
	"regexp"
	"time"
)

// LPYGPJXNBAUUJUFQ

const EmailReg = "^[A-Za-z0-9\\u4e00-\\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"


var userApi = api.User{
	DB: DB,
}
func Register(ctx *gin.Context)  {
	email := ctx.PostForm("email")
	verifyCode := ctx.PostForm("verifyCode")
	code := db.GetRedisDB().Get(email).Val()
	if code != verifyCode || code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "验证码错误！",
		})
	}
	password := ctx.PostForm("password")
	if emailBool, _ := regexp.MatchString(EmailReg, email); !emailBool {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "邮箱格式不对!",
		})
		return
	}
	if len(password) < 8 || len(password) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "密码格式不对!",
		})
		return
	}
	code, err := userApi.Register(email, password)
	ctx.JSON(code, err)
}

func SendVerifyCode(ctx *gin.Context)  {
	email := ctx.PostForm("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "邮箱不能为空!",
		})
	}
	if emailBool, _ := regexp.MatchString(EmailReg, email); !emailBool {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "邮箱格式不对!",
		})
		return
	}
	if err := et.SendEmail("您正在注册xxxxx服务，验证码为：6666", "注册", []string{email}); err != nil {
		log.Fatal("邮件发送失败，失败原因：" + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": err.Error(),
		})
		return
	}
	db.GetRedisDB().Set(email, "6666", time.Millisecond * 60)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "验证发已发送您的邮箱，如若没看到，请看垃圾邮箱",
	})
}
