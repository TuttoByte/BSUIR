package clients

import "github.com/shopspring/decimal"

type WeatherDataClient interface {
	LocationCurrentTemperature(lat decimal.Decimal, lon decimal.Decimal) (temperature decimal.Decimal, err error)
	LocationCurrentForcast(lat decimal.Decimal, lon decimal.Decimal) (ForecastResponse, error)
	LocationCurrentForcatByCity(cityName string) (ForecastResponse, error)
}
