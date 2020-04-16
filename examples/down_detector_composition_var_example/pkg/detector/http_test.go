package detector

import (
	"net/http"
	"testing"
)


func TestGetStatusCode(t *testing.T) {
	oldHttpClientGet := httpClientGet

	defer func() {
		httpClientGet = oldHttpClientGet
	}()

	httpClientGet = func(url string) (resp *http.Response, err error) {
		return &http.Response{}, nil
	}

	isDown, err := IsDown("http://eventbrite.com")

	if err != nil {
		t.Error("No error expected for this test")
	}

	if !isDown {
		t.Errorf("IsDown returned: %t, which is invalid.", isDown)
	}
}