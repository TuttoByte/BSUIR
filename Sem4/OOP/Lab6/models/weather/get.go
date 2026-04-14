package weather

import "github.com/shopspring/decimal"

type CurrentWeather struct {
	Temperature decimal.Decimal `json:"temperature"`
}

type CurrentWeatherRequest struct {
	Type string          `json:"type"`
	Lat  decimal.Decimal `json:"lat"`
	Lon  decimal.Decimal `json:"lon"`
}

type CurrentWeatherResponse struct {
	Lat         decimal.Decimal `json:"lat"`
	Lon         decimal.Decimal `json:"lon"`
	Temperature decimal.Decimal `json:"temperature"`
}
