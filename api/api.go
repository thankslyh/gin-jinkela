package api

import (
	"fmt"
	"gorm.io/gorm"
	"jinkela/jerror"
	"jinkela/model"
	"net/http"
	"time"
)

type User struct {
	DB *gorm.DB
}

// Register 注册api
func (user *User) Register(email, password string) (int, error)  {
	ret := &model.User{}
	if len(email) < 11 {
		return http.StatusBadRequest, jerror.EmailFormatError
	}
	user.DB.Table("users").Where("email = ?", email).Find(ret)
	fmt.Println(ret)
	if ret.ID != 0 {
		return http.StatusBadRequest, jerror.EmailAlreadyExsit
	}
	ret.Email, ret.Password = email, password
	ret.CreateTime, ret.UpdateTime = time.Now(), time.Now()
	if err := user.DB.Table("users").Create(&ret).Error; err != nil {
		return http.StatusOK, err
	}
	return http.StatusOK, nil
}