package auth_test

import (
	"jinkela/utils/auth"
	"testing"
	"time"
)

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRoYW5rc2x5aEBnbWFpbC5jb20iLCJwYXNzd29yZCI6IkxpITIxNTc3IiwiZXhwIjoxNjQxODk3ODg2LCJpc3MiOiJ0aGFua3NseWhAZ21haWwuY29tIn0.0OCfylYuB84PsDqUlgoWPGyUpJXHYYIa7paRyqWxSt0
func TestGenToken(t *testing.T) {
	token, _ := auth.GenToken(1, time.Millisecond * 60)
	t.Log(token)
}

func TestVerifyToken(t *testing.T) {
	auth.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRoYW5rc2x5aEBnbWFpbC5jb20iLCJwYXNzd29yZCI6IkxpITIxNTc3IiwiZXhwIjoxNjQxODk3ODg2LCJpc3MiOiJ0aGFua3NseWhAZ21haWwuY29tIn0.0OCfylYuB84PsDqUlgoWPGyUpJXHYYIa7paRyqWxSt0")
}