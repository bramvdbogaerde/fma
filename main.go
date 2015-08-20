package fma;

import(
	"strconv"
	"encoding/json"
	"github.com/bramvdbogaerde/fma/httputil"
)

func GetFeaturedTracks(page ...int) []Track{
  pageNumber := 1
  if len(page) >= 1{
		pageNumber = page[0]
	}

	options := map[string]string{};
	options["page"] = strconv.Itoa(pageNumber);

	return getTracks(options)
}

func getTracks(params map[string]string) []Track{
	parameters := ""
	for key,_ := range(params){
		parameters += key+"="+params[key];
	}

	response := httputil.DoGet("http://freemusicarchive.org/api/get/tracks.json?"+parameters)
	dataset := &DataSet{}
	json.Unmarshal([]byte(response), &dataset);

	return dataset.DataSet;
}
