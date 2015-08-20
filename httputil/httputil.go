package httputil

import(
	"net/http"
	"io/ioutil"
)

func DoGet(url string) string{
	response,_ := http.Get(url);

	defer response.Body.Close();
	body,_ := ioutil.ReadAll(response.Body);

	return string(body);
}
