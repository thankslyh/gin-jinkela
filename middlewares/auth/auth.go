package auth

import (
	"github.com/gin-gonic/gin"
	"jinkela/utils/auth"
	"log"
	"net/http"
	"strconv"
)

func Include(strs []string, target string) bool {
	for _, str := range strs {
		if str == target {
			return true
		}
	}
	return false
}

func AuthRequired(whiteList []string) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		log.Println(ctx.Request.URL)
		token := ctx.GetHeader("Authorization")
		if !Include(whiteList, ctx.Request.URL.Path) {
			t, err := auth.VerifyToken(token)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusUnauthorized,
					"msg": "token 失效",
				})
				return
			}
			if claims, ok := t.Claims.(*auth.CustomClaims); ok && t.Valid {
				strUserId := strconv.FormatInt(int64(claims.UserId), 10)
				ctx.Set("userId", strUserId)
			}
			return
		}
	}
}