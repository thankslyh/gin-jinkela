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
	"/api/tags/get",
	"/api/post/list",
	"/api/post/get",
}

func Run(r *gin.Engine) {
	root := r.Group("/api")
	root.Use(auth.AuthRequired(whitelist))
	rootUser := root.Group(UserRoute.Base)
	rootUser.POST("/register", UserRoute.Register)
	rootUser.POST("/send-verify-code", UserRoute.SendVerifyCode)
	rootUser.POST("/login", UserRoute.Login)
	rootUser.GET("/info", UserRoute.Info)

	rootTag := root.Group(TagRoute.Base)
	rootTag.POST("/add", TagRoute.Add)
	rootTag.GET("/get", TagRoute.GetTags)

	rootPost := root.Group(PostRoute.Base)
	rootPost.GET("/list", PostRoute.GetList)
	rootPost.GET("/get", PostRoute.GetPostById)
}