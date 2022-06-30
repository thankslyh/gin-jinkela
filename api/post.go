package api

import (
	"fmt"
	"gorm.io/gorm"
	"jinkela/model"
	"net/http"
)

type Post struct {
	DB *gorm.DB
}

func (p *Post) GetList() ([]model.SimplePost, int, error) {
	var ret []model.SimplePost
	err := p.DB.Table("post").Where("is_disable = ?", 0).Find(&ret).Error
	if err != nil {
		return ret, http.StatusBadRequest, err
	}
	return ret, http.StatusOK, nil
}

func (p *Post) GetPostById(id int) (model.Post, int, error)  {
	var ret model.Post
	fmt.Println(p.DB)
	err := p.DB.Table("post").Find(&ret, "id = ?", id).Error
	if err != nil {
		return ret, http.StatusBadRequest, err
	}
	return ret, http.StatusOK, nil
}