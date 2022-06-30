package model

import "time"

type Model struct {
	ID                     int64   `gorm:"primaryKey" json:"id" :"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type User struct {
	Model
	Email string `json:"email"`
	Password string `json:"password"`
	UserId                 int64 `json:"user_id"`
	Phone                  string `json:"phone"`
	NickName               string `json:"nick_name"`

}
