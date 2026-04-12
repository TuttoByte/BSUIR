package clients

import "github.com/shopspring/decimal"

type googleWeatherResponse struct {
	Temperature struct {
		Degrees decimal.Decimal `json:"degrees"`
	} `json:"temperature"`
}

type GoogleWeatherClient struct {
	apiKey  string
	baseURL string
}

func NewGoogleWetherClient(apiKey string, baseURL string) *GoogleWeatherClient {
	return &GoogleWeatherClient{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (g *GoogleWeatherClient) LocationCurrentTemperature(lat decimal.Decimal, lon decimal.Decimal) (temperature decimal.Decimal, err error) {
	return decimal.Decimal{}, err
}
