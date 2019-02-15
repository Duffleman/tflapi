package stoppoint

import (
	"github.com/duffleman/tflapi/http"
)

const (
	// RootURL is the base URL for the TFL API
	RootURL = "https://api.tfl.gov.uk"
)

type StopPointWrapper struct {
	http *http.Client
}

func NewWrapper(c *http.Client) *StopPointWrapper {
	return &StopPointWrapper{http: c}
}
