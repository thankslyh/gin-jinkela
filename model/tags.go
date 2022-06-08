package model

type Tag struct {
	Model
	Name string `json:"name"`
	Code string `json:"code"`
}