package router

import (
	"github.com/gin-gonic/gin"
	"jinkela/api"
	"jinkela/db"
	"jinkela/model"
	"jinkela/utils/auth"
	et "jinkela/utils/email"
	"log"
	"net/http"
	"regexp"
	"time"
)

// LPYGPJXNBAUUJUFQ

const EmailReg = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`

type User struct {
	Register, SendVerifyCode, Login, Info gin.HandlerFunc
	Base string
}

var userApi = api.User{
	DB: db.GetMysqlDB(),
}

var UserRoute = User {
	Register: register,
	SendVerifyCode: sendVerifyCode,
	Login: login,
	Info: info,
	Base: "/user",
}

func register(ctx *gin.Context)  {
	email := ctx.PostForm("email")
	verifyCode := ctx.PostForm("verifyCode")
	inpCode := db.GetRedisDB().Get(email).Val()
	log.Println("验证码=========:string=", inpCode, email)
	if inpCode != verifyCode || inpCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "验证码错误！",
		})
		return
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
			"code": http.StatusBadRequest,
			"msg": "密码格式不对!",
		})
		return
	}
	code, err := userApi.Register(email, password)
	if code != http.StatusOK {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg": err,
		})
		return
	}
	log.Println(code, err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "注册成功！",
	})
}

func sendVerifyCode(ctx *gin.Context)  {
	email := ctx.PostForm("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "邮箱不能为空!",
		})
	}
	if matched, _ := regexp.MatchString(EmailReg, email); !matched {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "邮箱格式不对!",
		})
		return
	}
	if err := et.SendEmail("您正在注册xxxxx服务，验证码为：6666", "注册", []string{email}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": err.Error(),
		})
		log.Fatal("邮件发送失败，失败原因：" + err.Error())
		return
	}
	db.GetRedisDB().Set(email, "6666", time.Minute * 2)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "验证发已发送您的邮箱，如若没看到，请看垃圾邮箱",
	})
}

func login(ctx *gin.Context) {
	var user model.User
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if matched, _ := regexp.MatchString(EmailReg, email); !matched {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg": "邮箱格式不对!",
		})
		return
	}
	db.GetMysqlDB().Table("users").Find(&user, "email = ?", email)
	if user.Password != password {
		ctx.JSON(400, gin.H{
			"code": 400,
			"msg": "密码错误！",
		})
		return
	}
	token, err := auth.GenToken(int(user.UserId), time.Hour * 2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": "生成token错误！",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"token": token,
		},
	})
}

func info(ctx *gin.Context) {
	userIdInter := ctx.MustGet("userId")
	userIdInt, _ := userIdInter.(int)
	userInfo, err := userApi.Info(userIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": "查询出错！",
		})
		return
	}
	//userInfoByte, err := json.Marshal(userInfo);
	//if  err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"code": http.StatusBadRequest,
	//		"msg": "解析出错！",
	//	})
	//	return
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "请求成功",
		"data": userInfo,
	})
}
