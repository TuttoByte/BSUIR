package clients

import "github.com/shopspring/decimal"

type CoordinateDataCleint interface {
	GetCurrentLocation(cityName string) (lat decimal.Decimal, lon decimal.Decimal, err error)
}
