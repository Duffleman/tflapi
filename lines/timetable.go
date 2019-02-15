package lines

import (
	"fmt"
	"net/url"
)

type Timetable struct {
	LineID          string      `json:"lineId"`
	LineName        string      `json:"lineName"`
	DepartureStopID string      `json:"departureStopId"`
	Direction       string      `json:"direction"`
	Stations        []*Station  `json:"stations"`
	Stops           interface{} `json:"stops"`
	Timetable       interface{} `json:"timetable"`
}

// GetTimetableForLine gets the timetable for a specified station on the given
// line
func (l *LineWrapper) GetTimetableForLine(stationID, line, direction string) (tt *Timetable, err error) {
	qs := url.Values{
		"direction": []string{direction},
	}

	url := fmt.Sprintf("%s/Line/%s/Timetable/%s?%s", RootURL, line, stationID, qs.Encode())

	err = l.http.GetJSON(url, &tt)

	return
}
