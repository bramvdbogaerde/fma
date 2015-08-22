package fma

type Track struct{
	Title string `json:"track_title"` // title of track
	Url string `json:"track_url"`     // url of track on freemusicarchive.org
	ImageFile string `json:"track_image_file"` // image of track
}
