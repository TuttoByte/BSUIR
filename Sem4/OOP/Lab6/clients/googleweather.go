package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
	"time"
)

type googleWeatherResponse struct {
	Temperature struct {
		Degrees decimal.Decimal `json:"degrees"`
	} `json:"temperature"`
}

type googleForecastResponse struct {
	CurrentTime time.Time `json:"currentTime"`
	TimeZone    struct {
		ID string `json:"id"`
	} `json:"timeZone"`
	IsDaytime        bool `json:"isDaytime"`
	WeatherCondition struct {
		IconBaseURI string `json:"iconBaseUri"`
		Description struct {
			Text         string `json:"text"`
			LanguageCode string `json:"languageCode"`
		} `json:"description"`
		Type string `json:"type"`
	} `json:"weatherCondition"`
	Temperature struct {
		Degrees decimal.Decimal `json:"degrees"`
		Unit    string          `json:"unit"`
	} `json:"temperature"`
	FeelsLikeTemperature struct {
		Degrees decimal.Decimal `json:"degrees"`
		Unit    string          `json:"unit"`
	} `json:"feelsLikeTemperature"`
	DewPoint struct {
		Degrees decimal.Decimal `json:"degrees"`
		Unit    string          `json:"unit"`
	} `json:"dewPoint"`
	HeatIndex struct {
		Degrees decimal.Decimal `json:"degrees"`
		Unit    string          `json:"unit"`
	} `json:"heatIndex"`
	WindChill struct {
		Degrees decimal.Decimal `json:"degrees"`
		Unit    string          `json:"unit"`
	} `json:"windChill"`
	RelativeHumidity int `json:"relativeHumidity"`
	UvIndex          int `json:"uvIndex"`
	Precipitation    struct {
		Probability struct {
			Percent int    `json:"percent"`
			Type    string `json:"type"`
		} `json:"probability"`
		Qpf struct {
			Quantity int    `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"qpf"`
	} `json:"precipitation"`
	ThunderstormProbability int `json:"thunderstormProbability"`
	AirPressure             struct {
		MeanSeaLevelMillibars decimal.Decimal `json:"meanSeaLevelMillibars"`
	} `json:"airPressure"`
	Wind struct {
		Direction struct {
			Degrees  int    `json:"degrees"`
			Cardinal string `json:"cardinal"`
		} `json:"direction"`
		Speed struct {
			Value decimal.Decimal `json:"value"`
			Unit  string          `json:"unit"`
		} `json:"speed"`
		Gust struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"gust"`
	} `json:"wind"`
	Visibility struct {
		Distance int    `json:"distance"`
		Unit     string `json:"unit"`
	} `json:"visibility"`
	CloudCover               int `json:"cloudCover"`
	CurrentConditionsHistory struct {
		TemperatureChange struct {
			Degrees decimal.Decimal `json:"degrees"`
			Unit    string          `json:"unit"`
		} `json:"temperatureChange"`
		MaxTemperature struct {
			Degrees decimal.Decimal `json:"degrees"`
			Unit    string          `json:"unit"`
		} `json:"maxTemperature"`
		MinTemperature struct {
			Degrees decimal.Decimal `json:"degrees"`
			Unit    string          `json:"unit"`
		} `json:"minTemperature"`
		Qpf struct {
			Quantity int    `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"qpf"`
	} `json:"currentConditionsHistory"`
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

	url := fmt.Sprintf("%s/v1/currentConditions:lookup?key=%s&location.latitude=%s&location.longitude=%s",
		g.baseURL, g.apiKey, lat.String(), lon.String())

	resp, err := http.Get(url)
	if err != nil {
		return decimal.Zero, errors.New("could not get current weather")
	}

	if resp.StatusCode != http.StatusOK {
		return decimal.Zero, errors.New(fmt.Sprintf("google wether returned bad status code: %d", resp.StatusCode))
	}

	var data googleWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return decimal.Zero, errors.New(fmt.Sprintf("could not decode json response: %s", err.Error()))
	}
	return data.Temperature.Degrees, err
}

func (g *GoogleWeatherClient) LocationCurrentForcast(lat decimal.Decimal, lon decimal.Decimal) (ForecastResponse, error) {
	return ForecastResponse{}, nil
}
func (g *GoogleWeatherClient) LocationCurrentForcatByCity(cityName string) (ForecastResponse, error) {
	return ForecastResponse{}, nil
}
