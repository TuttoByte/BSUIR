package clients

import "github.com/shopspring/decimal"

type openWeatherCityInfo struct {
	Name    string          `json:"name"`
	Lat     decimal.Decimal `json:"lat"`
	Lon     decimal.Decimal `json:"lon"`
	Country string          `json:"country"`
}

type openWeatherCoordResponce struct {
	Infos []openWeatherCityInfo
}

func newOpenWeatherCoordResponce(infos []openWeatherCityInfo) openWeatherCoordResponce {
	return openWeatherCoordResponce{
		Infos: infos,
	}
}

type OpenWeatherCoords struct {
	apiKey  string
	baseURL string
}

func NewOpenWeatherCoords(apiKey string, baseUrl string) *OpenWeatherCoords {
	return &OpenWeatherCoords{
		apiKey:  apiKey,
		baseURL: baseUrl,
	}
}

func (o *OpenWeatherCoords) GetCurrentLocation(cityName string) (decimal.Decimal, decimal.Decimal, error) {
	return decimal.Decimal{}, decimal.Decimal{}, nil
}
