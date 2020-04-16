package detector

import (
	"errors"
	"net/http"
)

var httpClientGet = http.Get

// IsDown will make a GET request to the indicated url.
func IsDown(url string) (bool, error) {
	response, err := httpClientGet(url)

	if err != nil {
		return false, errors.New("cannot detect the status of the endpoint")
	}

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return true, nil
	}

	return false, nil
}
