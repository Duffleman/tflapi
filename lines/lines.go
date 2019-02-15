package lines

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	// MaxIDs is the maximum number of inputs to give a retrieve API
	MaxIDs = 20
)

// ErrTooMuchInput is returned when you've given too many IDs to a retrieve endpoint
var ErrTooMuchInput = errors.New("too many IDs given")

// Line is a TFL line of transport
type Line struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Mode          string          `json:"modeName"`
	CreatedAt     time.Time       `json:"created"`
	ModifiedAt    time.Time       `json:"modified"`
	Disruptions   []interface{}   `json:"disruptions"`
	LineStatuses  []interface{}   `json:"lineStatuses"`
	RouteSections []*RouteSection `json:"routeSections"`
	ServiceTypes  []*ServiceTypes `json:"serviceTypes"`
}

// ServiceTypes is a TFL service type
type ServiceTypes struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// ListLinesByIDs gets lines that match the specified line ids
func (l *LineWrapper) ListLinesByIDs(lineIDs []string) (res []*Line, err error) {
	if len(lineIDs) > MaxIDs {
		return nil, ErrTooMuchInput
	}

	url := fmt.Sprintf("%s/Line/%s", RootURL, strings.Join(lineIDs, ","))

	err = l.http.GetJSON(url, &res)

	return
}

// ListRoutes gets all valid routes for all lines, including the name and id of
// the originating and terminating stops for each route
func (l *LineWrapper) ListRoutes(serviceType string) (lines []*Line, err error) {
	qs := url.Values{
		"serviceTypes": []string{serviceType},
	}

	url := fmt.Sprintf("%s/Line/Route?%s", RootURL, qs.Encode())

	err = l.http.GetJSON(url, &lines)

	return
}

// ListRoutesByLine gets all valid routes for given line ids, including the name
// and id of the originating and terminating stops for each route.
func (l *LineWrapper) ListRoutesByLine(inLines []string, serviceType string) (lines []*Line, err error) {
	qs := url.Values{
		"serviceTypes": []string{serviceType},
	}

	if len(inLines) > MaxIDs {
		return nil, ErrTooMuchInput
	}

	url := fmt.Sprintf("%s/Line/%s/Route?%s", RootURL, strings.Join(inLines, ","), qs.Encode())

	err = l.http.GetJSON(url, &lines)

	return
}
