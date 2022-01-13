package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jinkela/middlewares/auth"
)

var DB *gorm.DB

type Route struct {
	IsGroup bool
	Path string
	Method string
	HandlerFunc gin.HandlerFunc
	Children []Route
}

// 白名单
var whitelist = []string{
	"/api/user/login",
	"/api/user/register",
	"/api/user/send-verify-code",
}

func Run(r *gin.Engine) {
	root := r.Group("/api")
	root.Use(auth.AuthRequired(whitelist))
	rootUser := root.Group("/user")
	rootUser.POST("/register", Register)
	rootUser.POST("/send-verify-code", SendVerifyCode)
	rootUser.POST("/login", Login)
	rootUser.GET("/info", Info)
}