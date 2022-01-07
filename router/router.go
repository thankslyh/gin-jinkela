package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB)  {
	DB = db
}

type Route struct {
	IsGroup bool
	Path string
	Method string
	HandlerFunc gin.HandlerFunc
	Children []Route
}

func Run(r *gin.Engine) {
	root := r.Group("/api")

	rootUser := root.Group("/user")
	rootUser.POST("/register", Register)
	rootUser.POST("/send-verify-code", SendVerifyCode)
}