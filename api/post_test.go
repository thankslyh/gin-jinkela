package api_test

import (
	"jinkela/api"
	"jinkela/db"
	"testing"
)

func TestPost_GetList(t *testing.T) {
	apiPost := api.Post{
		DB: db.GetMysqlDB(),
	}
	apiPost.GetPostList("")
}
