package lines

import (
	"github.com/duffleman/tflapi/http"
)

const (
	// RootURL is the base URL for the TFL API
	RootURL = "https://api.tfl.gov.uk"
)

type LineWrapper struct {
	http *http.Client
}

func NewWrapper(c *http.Client) *LineWrapper {
	return &LineWrapper{http: c}
}
