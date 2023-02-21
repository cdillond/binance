package binance

import (
	"math"
	"strconv"
)

// MOVES FILTERS TO STRUCT FIELDS
func parseFilters(s Symbol) Symbol {
	for _, f := range s.RawFilters {
		ft, ok := f["filterType"]
		if !ok {
			continue
		}
		switch ft.(string) {
		case "PRICE_FILTER":
			tickSize, ok := f["tickSize"]
			if !ok {
				continue
			}
			tsf, err := strconv.ParseFloat(tickSize.(string), 64)
			if err != nil {
				continue
			}
			s.TickSize = int(math.Round(math.Abs(math.Log10(tsf))))
		case "LOT_SIZE":
			stepSize, ok := f["stepSize"]
			if !ok {
				continue
			}
			ssf, err := strconv.ParseFloat(stepSize.(string), 64)
			if err != nil {
				continue
			}
			s.StepSize = int(math.Round(math.Abs(math.Log10(ssf))))
		case "MIN_NOTIONAL":
			minNotional, ok := f["minNotional"]
			if !ok {
				continue
			}
			mmf, err := strconv.ParseFloat(minNotional.(string), 64)
			if err != nil {
				continue
			}
			s.MinNotional = mmf
		}
	}
	return s
}
