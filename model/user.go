package model

import "time"

type Model struct {
	ID                     int64   `gorm:"primaryKey" json:"id" :"id"`
	CreateTime, UpdateTime time.Time `json:"create_time"`
}

type User struct {
	Model
	Email string `json:"email"`
	Password string `json:"password"`
	UserId                 int64 `json:"user_id"`
	Phone                  string `json:"phone"`
	NickName               string `json:"nick_name"`

}
