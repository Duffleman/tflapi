package lines

import "time"

type RouteSection struct {
	Name            string    `json:"name"`
	ValidTo         time.Time `json:"validTo"`
	ValidFrom       time.Time `json:"validFrom"`
	Originator      string    `json:"originator"`
	Destination     string    `json:"destination"`
	ServiceType     string    `json:"serviceType"`
	Direction       string    `json:"direction"`
	OriginationName string    `json:"originationName"`
	DestinationName string    `json:"destinationName"`
}
