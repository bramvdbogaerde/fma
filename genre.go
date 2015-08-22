package fma;

type Genre struct{
	Id int `json:"genre_id"`
	ParentId int `json:"genre_parent_id"`
	Title string `json:"genre_title"`
	Handle string `json:"genre_handle"`
	Color string `json:"genre_color"`
}
