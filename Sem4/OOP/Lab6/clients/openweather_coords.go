package clients

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
)

type OpenWeatherCityInfo struct {
	Name    string          `json:"name"`
	Lat     decimal.Decimal `json:"lat"`
	Lon     decimal.Decimal `json:"lon"`
	Country string          `json:"country"`
}

type openWeatherCoordResponce struct {
	Infos []OpenWeatherCityInfo
}

func newOpenWeatherCoordResponce(infos []OpenWeatherCityInfo) openWeatherCoordResponce {
	if len(infos) == 0 {
		return openWeatherCoordResponce{
			Infos: make([]OpenWeatherCityInfo, 0),
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

func (o *OpenWeatherCoords) GetLocation(cityName string) ([]OpenWeatherCityInfo, error) {

	url := fmt.Sprintf("%s/geo/1.0/direct?q=%s&limit=5&appid=%s", o.baseURL, cityName, o.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch location: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	var data []OpenWeatherCityInfo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return data, nil
}
