package location

import "github.com/shopspring/decimal"

type CurrentLocation struct {
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
}
