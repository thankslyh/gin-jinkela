package api

import (
	"gorm.io/gorm"
	"jinkela/jerror"
	"jinkela/model"
	"jinkela/utils/auth"
	"log"
	"net/http"
	"time"
)

type User struct {
	DB *gorm.DB
}

// Register 注册api
func (user *User) Register(email, password string) (int, error)  {
	var ret model.User
	if len(email) < 11 {
		return http.StatusBadRequest, jerror.EmailFormatError
	}
	 err := user.DB.Table("users").Find(&ret, "email = ?", email).Error
	 if err != nil {
		 log.Fatalln(err)
		 return http.StatusBadRequest, err
	 }
	log.Println(ret)
	if ret.ID != 0 || ret.Email != "" {
		return http.StatusBadRequest, jerror.EmailAlreadyExsit
	}
	ret.Email, ret.Password = email, password
	ret.CreateTime, ret.UpdateTime = time.Now(), time.Now()
	if err := user.DB.Table("users").Create(&ret).Error; err != nil {
		return http.StatusOK, err
	}
	return http.StatusOK, nil
}

func (user *User) Login(email, password string) (int, string, error) {
	var ret model.User
	user.DB.Table("users").Find(&ret, "email = ?", email)
	if ret.Password != password {
		return http.StatusBadRequest, "", jerror.PasswordError
	}
	token, err := auth.GenToken(email, password, time.Hour * 2)
	if err != nil {
		return http.StatusBadRequest, "", err
	}
	return http.StatusOK, token, nil
}