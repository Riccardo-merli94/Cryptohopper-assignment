package exchange

import (
	"time"
)

type Exchange interface {
	Chart(market string, startTime int64, endTime int64, period string) ([]Item, error)
}

type Item struct {
	Open        float64   `json:"Open"`
	High        float64   `json:"High"`
	Low         float64   `json:"Low"`
	Close       float64   `json:"Close"`
	BaseVolume  float64   `json:"BaseVolume"`
	QuoteVolume float64   `json:"QuoteVolume"`
	OpenTime    time.Time `json:"OpenTime"`
}
