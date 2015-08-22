package fma;

type Genre struct{
	Id int `json:"genre_id"` // Genre id (so you can search for tracks with this genre)
	ParentId int `json:"genre_parent_id"` // Parent id of genre
	Title string `json:"genre_title"` 	  // Title of genre
	Handle string `json:"genre_handle"`   //
	Color string `json:"genre_color"`     // Color as displayed on freemusicarchive.org
}
