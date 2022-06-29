package api

import (
	"gorm.io/gorm"
	"jinkela/jerror"
	"jinkela/model"
	"net/http"
	"time"
)

type Tag struct {
	DB *gorm.DB
}

// Add 添加标签
func (tag *Tag) Add(code, name string) (int, error)  {
	var ret model.Tag
	err := tag.DB.Table("tags").Find(&ret, "code = ?", code).Error
	if err != nil {
		return http.StatusBadRequest, err
	}
	if ret.ID != 0 {
		return http.StatusBadRequest, jerror.TagExist
	}
	ret.Code, ret.Name = code, name
	ret.CreateTime, ret.UpdateTime = time.Now(), time.Now()
	if err := tag.DB.Table("tags").Create(&ret).Error; err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

// GetAll 查询所有标签
func (tag *Tag) GetAll() ([]model.Tag, int, error)  {
	var ret []model.Tag
	err := tag.DB.Table("tags").Where("is_disable = ?", 0).Find(&ret).Error
	if err != nil {
		return ret, http.StatusBadRequest, err
	}
	return ret, http.StatusOK, nil
}