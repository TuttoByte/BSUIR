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

func (c *CurrentWeatherController) GetCurrentForecast(lat decimal.Decimal, lon decimal.Decimal) (clients.ForecastResponse, error) {

	forecast, err := c.Client.LocationCurrentForcast(lat, lon)
	if err != nil {
		return clients.ForecastResponse{}, err
	}
	return forecast, nil
}

func (c *CurrentWeatherController) GetMultipleWeather(weathers []weather.CurrentWeatherRequest) ([]weather.CurrentWeatherResponse, error) {
	data := make([]weather.CurrentWeatherResponse, len(weathers))
	for i, w := range weathers {
		temp, err := c.Client.LocationCurrentTemperature(w.Lat, w.Lon)
		if err != nil {
			return nil, err
		}
		data[i].Lat = w.Lat
		data[i].Lon = w.Lon
		data[i].Temperature = temp
	}
	return data, nil
}
