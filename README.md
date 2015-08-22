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
	
}
``
