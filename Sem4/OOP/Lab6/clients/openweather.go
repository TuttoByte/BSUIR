package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

type openWeatherResponse struct {
	Main struct {
		Temp decimal.Decimal `json:"temp"`
	} `json:"main"`
}

type openForecastResponce struct {
	Coord struct {
		Lon decimal.Decimal `json:"lon"`
		Lat decimal.Decimal `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      decimal.Decimal `json:"temp"`
		FeelsLike decimal.Decimal `json:"feels_like"`
		TempMin   decimal.Decimal `json:"temp_min"`
		TempMax   decimal.Decimal `json:"temp_max"`
		Pressure  decimal.Decimal `json:"pressure"`
		Humidity  int             `json:"humidity"`
		SeaLevel  int             `json:"sea_level"`
		GrndLevel int             `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed decimal.Decimal `json:"speed"`
		Deg   int             `json:"deg"`
		Gust  decimal.Decimal `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH decimal.Decimal `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type OpenWeatherClient struct {
	apiKey  string
	baseURL string
}

func NewOpenWeatherClient(apiKey string, baseURL string) *OpenWeatherClient {
	return &OpenWeatherClient{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// Implementation of WeatherDataClient
func (o *OpenWeatherClient) LocationCurrentTemperature(lat decimal.Decimal, lon decimal.Decimal) (decimal.Decimal, error) {
	url := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=metric",
		o.baseURL, lat.String(), lon.String(), o.apiKey)

	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to call openweather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return decimal.Zero, fmt.Errorf("openweather returned bad status: %d", resp.StatusCode)
	}

	var data openWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return decimal.Zero, fmt.Errorf("failed to decode response: %w", err)
	}

	return data.Main.Temp, nil
}

func (o *OpenWeatherClient) LocationCurrentForcast(lat decimal.Decimal, lon decimal.Decimal) (ForecastResponse, error) {

	url := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=metric",
		o.baseURL, lat.String(), lon.String(), o.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return ForecastResponse{}, fmt.Errorf("failed to call openweather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ForecastResponse{}, fmt.Errorf("openweather returned bad status: %d", resp.StatusCode)
	}

	var data openForecastResponce
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ForecastResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	forcResp := OpenResponseToForecast(data)
	return forcResp, nil
}
func (o *OpenWeatherClient) LocationCurrentForcatByCity(cityName string) (ForecastResponse, error) {
	return ForecastResponse{}, nil
}
