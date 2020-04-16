package detector

import (
	"log"
	"net/http"
	"testing"
)

type MyFakeHttpClient struct {

}

func (m MyFakeHttpClient) Get(url string) (resp *http.Response, err error) {
	log.Println("Fake request post to: " + url)
	return &http.Response{StatusCode: 0}, nil
}

func TestGetStatusCode(t *testing.T) {
	httpClient := MyFakeHttpClient{}
	detector := Client{
		httpClient: httpClient,
	}

	isDown, err := detector.IsDown("https://eventbrite.com")

	if err != nil {
		t.Error("No error expected for this test")
	}

	if !isDown {
		t.Errorf("IsDown returned: %t, which is invalid.", isDown)
	}
}