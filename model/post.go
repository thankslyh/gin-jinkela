package model

type SimplePost struct {
	Model
	Title string `json:"title"`
	ReadCount int `json:"read_count"`
	StarCount int `json:"star_count"`
	CoverImg string `json:"cover_img"`
}

type Post struct {
	SimplePost
	Content string `json:"content"`
}