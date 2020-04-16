package detector

import (
	"errors"
	"net/http"
)

type IHttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// IsDown will make a GET request to the indicated url.
func IsDown(url string, client IHttpClient) (bool, error) {
	response, err := client.Get(url)

	if err != nil {
		return false, errors.New("cannot detect the status of the endpoint")
	}

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return true, nil
	}

	return false, nil
}
