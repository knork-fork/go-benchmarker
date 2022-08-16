package request

import (
	"net/http"
)

func Send(url string) bool {
	var result bool
	response, error := http.Get(url)

	if error != nil {
		result = false
	} else {
		result = response.StatusCode == 200
		response.Body.Close()
	}

	return result
}
