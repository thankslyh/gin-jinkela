package router

import (
	"github.com/gin-gonic/gin"
	"jinkela/api"
	"net/http"
)

type Tags struct {
	Add, GetTags, Update, Delete, Info gin.HandlerFunc
	Base string
}

var tagApi api.Tag

func getTagApi() *api.Tag {
	tagApi.DB = getDB()
	return &tagApi
}
var TagRoute = Tags{
	Add: add,
	GetTags: getTags,
	Base: "/tags",
}

func add(c *gin.Context)  {
	tagCode := c.PostForm("code")
	tagName := c.PostForm("name")
	code, err := getTagApi().Add(tagCode, tagName)
	if code != http.StatusOK {
		c.JSON(code, gin.H{
			"code": code,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg": "成功",
		"data": "",
	})
}

func getTags(c *gin.Context)  {
	data, code, err := getTagApi().GetAll()
	if code != http.StatusOK {
		c.JSON(code, gin.H{
			"code": code,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg": "success",
		"data": data,
	})

}