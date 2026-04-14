package controllers

import (
	"github.com/shopspring/decimal"
	"main/clients"
	"main/models/weather"
)

type CurrentWeatherController struct {
	Client clients.WeatherDataClient
}

func NewCurrentWeatherController(client clients.WeatherDataClient) *CurrentWeatherController {
	return &CurrentWeatherController{
		Client: client,
	}
}

func (c *CurrentWeatherController) GetCurrentWeather(lat decimal.Decimal, lon decimal.Decimal) (weather.CurrentWeather, error) {

	temperature, err := c.Client.LocationCurrentTemperature(lat, lon)
	if err != nil {
		return weather.CurrentWeather{}, err
	}

	return weather.CurrentWeather{
		Temperature: temperature,
	}, nil
}

func (c *CurrentWeatherController) GetCurrentForecast() {

}
