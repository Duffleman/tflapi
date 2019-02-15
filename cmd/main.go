package main

import (
	"github.com/duffleman/tflapi"

	"github.com/kr/pretty"
)

const (
	// Whitechapel is the NaptanCode for Whitechapel Underground Station
	Whitechapel = "940GZZLUWPL"
)

func main() {
	client := tflapi.NewClient("", "", nil)

	stations, err := client.Line.ListStationsOnLine("district", false)
	if err != nil {
		panic(err)
	}

	for _, station := range stations {
		if station.ID == Whitechapel {
			pretty.Println(station)
		}
	}
}
