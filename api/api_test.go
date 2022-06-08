package api_test

import (
	"jinkela/api"
	"jinkela/db"
	"testing"
)

func TestUser_Register(t *testing.T) {
	apiUser := api.User{
		DB: db.GetMysqlDB(),
	}
	apiUser.Register("826567584@qq.com", "Li!21577")
}

func TestTag_Add(t *testing.T) {
	apiTag := api.Tag{
		DB: db.GetMysqlDB(),
	}
	apiTag.Add("nodejs", "NodeJs")
}
