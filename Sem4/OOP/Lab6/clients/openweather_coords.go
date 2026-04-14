package clients

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
)

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
	if len(infos) == 0 {
		return openWeatherCoordResponce{
			Infos: make([]openWeatherCityInfo, 0),
		}

	}
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

	url := fmt.Sprintf("%s/geo/1.0/direct?q=%s&limit=5&appid=%s", o.baseURL, cityName, o.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return decimal.Zero, decimal.Zero, fmt.Errorf("failed to fetch location: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return decimal.Zero, decimal.Zero, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data := newOpenWeatherCoordResponce(nil)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return decimal.Zero, decimal.Zero, fmt.Errorf("failed to decode response: %w", err)
	}

	return data.Infos[0].Lat, data.Infos[0].Lon, nil
}
