package lines

import (
	"fmt"
	"net/url"
	"strings"
)

// Mode of transport
type Mode struct {
	IsTFLService       bool   `json:"isTflService"`
	IsFarePaying       bool   `json:"isFarePaying"`
	IsScheduledService bool   `json:"isScheduledService"`
	Name               string `json:"modeName"`
}

// ListModes gets a list of valid modes
func (l *LineWrapper) ListModes() (modes []*Mode, err error) {
	err = l.http.GetJSON(fmt.Sprintf("%s/Line/Meta/Modes", RootURL), &modes)

	return
}

// ListLinesAndRoutesByModes gets all lines and their valid routes for given
// modes, including the name and id of the originating and terminating stops for
// each route
func (l *LineWrapper) ListLinesAndRoutesByModes(modes []string, serviceTypes []string) (res []*Line, err error) {
	if len(modes) > MaxIDs {
		return nil, ErrTooMuchInput
	}

	qs := url.Values{}

	for _, st := range serviceTypes {
		qs.Add("serviceTypes[]", st)
	}

	url := fmt.Sprintf("%s/Line/Mode/%s/Route?%s", RootURL, strings.Join(modes, ","), qs.Encode())

	err = l.http.GetJSON(url, &res)

	return
}
