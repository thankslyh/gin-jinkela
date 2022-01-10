package auth_test

import (
	"jinkela/utils/auth"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	token, _ := auth.GenToken("thankslyh@gmail.com", "Li!21577", time.Millisecond * 60)
	t.Log(token)
}
