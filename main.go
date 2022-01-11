package main

import (
	"github.com/gin-gonic/gin"
	"jinkela/router"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	router.Run(r)
	r.Run(":4396")
}
