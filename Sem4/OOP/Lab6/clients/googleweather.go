package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
)

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
