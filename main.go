package fma;

import(
	"strconv"
	"encoding/json"
	"github.com/bramvdbogaerde/fma/httputil"
)

func GetFeaturedTracks(page ...int) Tracks{
  pageNumber := 1
  if len(page) >= 1{
		pageNumber = page[0]
	}

	options := map[string]string{};
	options["page"] = strconv.Itoa(pageNumber);

	return GetTracks(options)
}

func GetTracks(params map[string]string) Tracks{
	parameters := buildParameters(params)

	response := httputil.DoGet("http://freemusicarchive.org/api/get/tracks.json?"+parameters)
	tracksDataSet := Tracks{}
	json.Unmarshal([]byte(response), &tracksDataSet);

	return tracksDataSet;
}

func GetGenres(params map[string] string) Genres{
	parameters := buildParameters(params);
	response := httputil.DoGet("http://freemusicarchive.org/api/get/genres.json?"+parameters)

	genresDataSet := Genres{}
	json.Unmarshal([]byte(response), &genresDataSet);

	return genresDataSet;
}
