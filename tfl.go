package tflapi

import (
	"net/http"

	dflHttp "github.com/duffleman/tflapi/http"
	"github.com/duffleman/tflapi/lines"
	"github.com/duffleman/tflapi/stoppoint"
)

const (
	// RootURL is the base URL for the TFL API
	RootURL = "https://api.tfl.gov.uk"
)

// Client is the main exported struct which contains all the methods this API provides
type Client struct {
	applicationID  string
	applicationKey string
	http           *dflHttp.Client

	Line      *lines.LineWrapper
	StopPoint *stoppoint.StopPointWrapper
}

// NewClient returns a TFL API Client
func NewClient(id, key string, client *http.Client) *Client {
	var c *http.Client

	if client != nil {
		c = client
	} else {
		c = http.DefaultClient
	}

	http := dflHttp.NewClient(c)

	line := lines.NewWrapper(http)
	stoppoint := stoppoint.NewWrapper(http)

	return &Client{
		applicationID:  id,
		applicationKey: key,
		http:           http,
		Line:           line,
		StopPoint:      stoppoint,
	}
}
