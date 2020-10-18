package tfl

import (
	"time"
)

const (
	DLR          string = "dlr"
	Overground          = "overground"
	Tube                = "tube"
	TFLRail             = "tflrail"
	Bus                 = "bus"
	NationalRail        = "national-rail"
)

type Line struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	ModeName   string    `json:"modeName"`
	CreatedAt  time.Time `json:"created"`
	ModifiedAt time.Time `json:"modified"`
}

type Stop struct {
	ID             string          `json:"id"`
	Status         bool            `json:"status"`
	Name           string          `json:"commonName"`
	Modes          []string        `json:"modes"`
	ICSCode        string          `json:"icsCode"`
	StationNaptan  string          `json:"stationNaptan"`
	LineModeGroups []LineModeGroup `json:"lineModeGroups"`
	Lat            float64         `json:"lat"`
	Lon            float64         `json:"lon"`
}

type LineModeGroup struct {
	ModeName       string   `json:"modeName"`
	LineIdentifier []string `json:"lineIdentifier"`
}
