FreeMusicArchive client in Go
=============================

This is a FreeMusicArchive client for Go, it interacts with the FreeMusicArchive API.

__Still work in progress, this isn't finished yet__

## Documentation

## Package: fma
### Methods
#### GetFeaturedTracks(page int)
returns Tracks


Example
```
// Prints the featured tracks titles
tracks := fma.GetFeaturedTracks();

for _,track := range(tracks.Tracks){
	fmt.Println(track.Title)
}
```

#### GetTracks(params map[string] string)
return Genres

Example
```
// Prints all the available genres
genres := fma.GetGenres()

for _,genre := range(genres.Genres){
	fmt.Println(genre.Title)
}
```
### Structs
#### Track
```
type Track struct{
	Title string `json:"track_title"` // track title
	Url string `json:"track_url"`     // url to track on freemusicarchive.org
	ImageFile string `json:"track_image_file"` // image for track
}
```

#### Genre
```
type Genre struct{
	Id int `json:"genre_id"` // Genre id (so you can search for tracks with this genre)
	ParentId int `json:"genre_parent_id"` // Parent id of genre
	Title string `json:"genre_title"` 	  // Title of genre
	Handle string `json:"genre_handle"`   //
	Color string `json:"genre_color"`     // Color as displayed on freemusicarchive.org
}
```
