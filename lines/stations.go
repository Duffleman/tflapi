package lines

import (
	"fmt"
)

// Station is a struct for a TFL station
type Station struct {
	ID                   string                  `json:"naptanId"`
	CommonName           string                  `json:"commonName"`
	PlaceType            string                  `json:"placeType"`
	PlatformName         *string                 `json:"platformName,omitempty"`
	Indicator            *string                 `json:"indicator,omitempty"`
	StopLetter           *string                 `json:"stopLetter,omitempty"`
	Modes                []string                `json:"modes"`
	ICSCode              string                  `json:"icsCode"`
	SMSCode              *string                 `json:"smsCode,omitempty"`
	StopType             string                  `json:"stopType"`
	StationNaptan        string                  `json:"stationNaptan"`
	AccessibilitySummary *string                 `json:"accessibilitySummary,omitempty"`
	HubNaptanCode        *string                 `json:"hubNaptanCode,omitempty"`
	Lines                []*StationLine          `json:"lines"`
	LineGroup            []*LineGroup            `json:"lineGroup"`
	AdditionalProperties []*AdditionalProperties `json:"additionalProperties"`
	Children             []*Station              `json:"children"`
	Latitude             float64                 `json:"lat"`
	Longitude            float64                 `json:"lon"`
}

type LineGroup struct {
	StationAtcoCode string   `json:"stationAtcoCode"`
	LineIdentifier  []string `json:"lineIdentifier"`
}

type AdditionalProperties struct {
	Category        string `json:"category"`
	Key             string `json:"key"`
	SourceSystemKey string `json:"sourceSystemKey"`
	Value           string `json:"value"`
}

// StationLine is a sub-struct from
type StationLine struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	URI             string      `json:"uri"`
	StationLineType string      `json:"type"`
	Crowding        interface{} `json:"crowding"`
	RouteType       string      `json:"routeType"`
	Status          string      `json:"status"`
}

// ListStationsOnLine lists the stations on a given line
func (l *LineWrapper) ListStationsOnLine(line string, includeNationalRail bool) (stations []*Station, err error) {
	url := fmt.Sprintf("%s/Line/%s/StopPoints?tflOperatedNationalRailStationsOnly=%t", RootURL, line, includeNationalRail)

	err = l.http.GetJSON(url, &stations)

	return
}
