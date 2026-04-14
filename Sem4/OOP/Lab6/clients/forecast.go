package clients

import (
	"github.com/shopspring/decimal"
)

type Forecast interface {
	Temperature() decimal.Decimal
	FeelsLikeTemperature() decimal.Decimal
	Humidity() int
	WindSpeed() decimal.Decimal
	Pressure() decimal.Decimal
	Visibility() int
	Clouds() int
}

// ForecastResponse общая модель для Swagger
// swagger:model forecast
type ForecastResponse struct {
	Temperature decimal.Decimal `json:"temperature" example:"15.5"`
	FeelsLike   decimal.Decimal `json:"feelsLikeTemperature" example:"14.0"`
	Humidity    int             `json:"humidity" example:"65"`
	WindSpeed   decimal.Decimal `json:"windSpeed" example:"10"`
	Pressure    decimal.Decimal `json:"pressure" example:"1013"`
	Visibility  int             `json:"visibility" example:"10000"`
	Clouds      int             `json:"clouds" example:"50"`
}

func GoogleResponseToForecast(response googleForecastResponse) ForecastResponse {
	return ForecastResponse{
		Temperature: response.Temperature.Degrees,
		FeelsLike:   response.FeelsLikeTemperature.Degrees,
		Humidity:    response.RelativeHumidity,
		WindSpeed:   response.Wind.Speed.Value,
		Pressure:    response.AirPressure.MeanSeaLevelMillibars,
		Visibility:  response.Visibility.Distance,
		Clouds:      response.CloudCover,
	}
}

func OpenResponseToForecast(response openForecastResponce) ForecastResponse {
	return ForecastResponse{
		Temperature: response.Main.Temp,
		FeelsLike:   response.Main.FeelsLike,
		Humidity:    response.Main.Humidity,
		WindSpeed:   response.Wind.Speed,
		Pressure:    response.Main.Pressure,
		Visibility:  response.Visibility,
		Clouds:      response.Clouds.All,
	}
}
