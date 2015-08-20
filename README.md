FreeMusicArchive client in Go
=============================

This is a FreeMusicArchive client for Go, it interacts with the FreeMusicArchive API.

## Documentation

## Package: fma
### Methods
#### GetFeaturedTracks(page int)
returns Track[]
Example
```
// Prints the featured tracks titles
tracks := fma.GetFeaturedTracks();

for _,track := range(tracks){
	fmt.Println(track.Title)
}
```