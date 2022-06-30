package router

import (
	"github.com/gin-gonic/gin"
	"jinkela/api"
	"net/http"
	"strconv"
)

type Post struct {
	GetList, GetPostById gin.HandlerFunc
	Base string
}

var postApi api.Post

func getPostApi() *api.Post {
	postApi.DB = getDB()
	return &postApi
}

var PostRoute = Post{
	GetList: getList,
	GetPostById: getPostById,
	Base: "/post",
}

func getList(ctx *gin.Context)  {
	data, code, err := getPostApi().GetList()
	if code != http.StatusOK {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg": err.Error(),
			"data": data,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"code": code,
		"msg": "success",
		"data": data,
	})
}

func getPostById(ctx *gin.Context) {
	id := ctx.Query("id")
	iniId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": 500,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	data, code, err := getPostApi().GetPostById(iniId)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg": err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"code": code,
		"msg": "success",
		"data": data,
	})
}