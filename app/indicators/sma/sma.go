package sma

import (
	"app/exchanges"
	"time"
)

type Sma struct {
	Exchange exchange.Exchange
	Market string
	Period string
	periodInSeconds int64
}

type Signal struct {
	Action string
}

const samplesStart = 8
const samplesEnd = 55

func (self Sma) Calculate() (Signal, error){
	offset := samplesStart
	//calculate start timestamp based on the number of sample (55) * the period in seconds
	start := time.Now().Unix() - (self.periodInSeconds * (samplesEnd))
	//end is the most recent
	end := time.Now().Unix()

	//get the chart for the selected Exchange
	chart, err := self.Exchange.Chart(self.Market, start , end, self.Period)

	if err != nil{
		return Signal{}, err
	}

	//calculate the first comparison term sma(8)
	avgStart := self.sma(samplesStart, chart, 0)
	//calculate the second comparison term sma(8)
	avgEnd := self.sma(samplesEnd, chart, 0)
	//calculate the previous sma(8)
	prevAvgStart := self.sma(samplesStart, chart, offset)

	println(avgStart, avgEnd, prevAvgStart)

	if avgStart < avgEnd && avgStart >= prevAvgStart{
		return Signal{Action: "sell"}, nil
	} else if avgStart > avgEnd && avgStart <= prevAvgStart {
		return Signal{Action: "buy"}, nil
	} else {
		return Signal{Action: "neutral"}, nil
	}
}

func (self Sma) sma(length int, data []exchange.Item, offset int) float64{
	var sum float64
	for _, item := range data[len(data) - (length + offset):len(data) - 1 - offset] {
		sum += item.Close
	}
	return sum/float64(length)
}

func New (e exchange.Exchange, market string, period string) (Sma, error) {
	duration, err := time.ParseDuration(period)

	if err != nil && period == "1d" {
		duration, err = time.ParseDuration("24h")
	}

	seconds := duration.Seconds()
	return Sma{Exchange: e, Market: market, Period: period, periodInSeconds: int64(seconds)}, err
}