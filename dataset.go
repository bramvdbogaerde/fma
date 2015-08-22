package fma;

// datasets

type DataSet struct{
	Title string `json:"title"`
	Total int `json:"total"`
	TotalPages int `json:"total_pages"`
	Limit int `json:"limit"`
}

// http://freemusicarchive.com/api/tracks.{{format}}
type Tracks struct{
	DataSet
	Tracks []Track `json:"dataset"`
}

// http://freemusicarchive.com/api/genres.{{format}}
type Genres struct{
	DataSet
	Genres []Genre `json:"dataset"`
}
