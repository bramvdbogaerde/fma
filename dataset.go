package fma;

type DataSet struct{
	Title string `json:"title"`
	Total int `json:"total"`
	DataSet []Track `json:"dataset"`
}
