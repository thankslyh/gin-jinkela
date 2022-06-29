package model

type SimplePost struct {
	Model
	Title string `json:"title"`
	ReadNum int `json:"read_num"`
	StarNum int `json:"star_num"`
	CoverImg string `json:"cover_img"`
}

type Post struct {
	SimplePost
	Content string `json:"content"`
}