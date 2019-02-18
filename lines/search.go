package lines

import (
	"fmt"
	"net/url"
)

type SearchResult struct {
	Input   string        `json:"input"`
	Matches []SearchMatch `json:"searchMatches"`
}

type SearchMatch struct {
	LineID            string                 `json:"lineId"`
	Mode              string                 `json:"mode"`
	LineName          string                 `json:"lineName"`
	LineRouteSections []*LineRouteSection    `json:"lineRouteSection"`
	RouteSections     []*MatchedRouteSection `json:"matchedRouteSections"`
	Stops             []*MatchedStop         `json:"matchedStops"`
	ID                string                 `json:"id"`
	URL               string                 `json:"url"`
	Name              string                 `json:"name"`
	Latitude          float64                `json:"lat"`
	Longitude         float64                `json:"lon"`
}

type MatchedRouteSection struct {
	ID int `json:"id"`
}

type LineRouteSection struct {
	RouteID                int    `json:"routeId"`
	VehicleDestinationText string `json:"vehicleDestinationText"`
	Direction              string `json:"direction"`
	Destination            string `json:"destination"`
	FromStation            string `json:"fromStation"`
	ToStation              string `json:"toStation"`
	ServiceType            string `json:"serviceType"`
}

type MatchedStop struct {
	RouteID              int      `json:"routeId"`
	ParentID             string   `json:"parentId"`
	StationID            string   `json:"stationId"`
	IcsID                string   `json:"icsId"`
	TopMostParentID      string   `json:"topMostParentId"`
	Direction            string   `json:"direction"`
	Towards              string   `json:"towards"`
	Modes                []string `json:"modes"`
	StopType             string   `json:"stopType"`
	StopLetter           string   `json:"stopLetter"`
	Zone                 string   `json:"zone"`
	AccessibilitySummary string   `json:"accessibilitySummary"`
	HasDisruption        bool     `json:"hasDisruption"`
	Lines                []*Line  `json:"lines"`
	Status               bool     `json:"status"`
	ID                   string   `json:"id"`
	URL                  string   `json:"url"`
	Name                 string   `json:"name"`
	Latitude             float64  `json:"lat"`
	Longitude            float64  `json:"lon"`
}

// Search for lines or routes matching the query string
func (l *LineWrapper) Search(query string, modes []string, serviceType []string) (res SearchResult, err error) {
	qs := url.Values{}

	if len(serviceType) > 0 {
		for _, st := range serviceType {
			qs.Add("serviceTypes[]", st)
		}
	}

	url := fmt.Sprintf("%s/Line/Search/%s?%s", RootURL, query, qs.Encode())

	err = l.http.GetJSON(url, &res)

	return
}
