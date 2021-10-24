package coinbase

import (
	"app/exchanges"
	"encoding/json"
	"fmt"
	"net/http"
)

type Coinbase struct {
	Url string
}

func (self Coinbase) Chart(market string, startTime int64, endTime int64, period string) ([]exchange.Item, error) {
	fullUrl := self.makeURL(market, startTime, endTime, period)

	resp, err := http.Get(fullUrl)
	if err != nil || resp.StatusCode != 200{
		return []exchange.Item{}, err
	}

	var data []exchange.Item
	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return []exchange.Item{}, err
	}
	return data, err
}

func (self Coinbase) makeURL(market string, startTime int64, endTime int64, period string) string {
	return fmt.Sprintf("%s?pair=%s&start=%d&end=%d&period=%s", self.Url, market, startTime, endTime, period)
}

func New() Coinbase{
	return Coinbase{Url:"http://cryptohopper-ticker-frontend.us-east-1.elasticbeanstalk.com/v1/coinbasepro/candles"}
}