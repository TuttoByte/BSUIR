package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"main/clients"
	"main/controllers"
	"main/models/weather"
	"main/shared/responses"
	"main/shared/utils"
)

type WeatherHandler struct {
	Controller controllers.CurrentWeatherController
}

func NewCurrentWeatherHandler(wtype string) *WeatherHandler {
	switch wtype {
	case "open":
		return NewOpenWeatherHandler()
	case "google":
		return NewGoogleWeatherHandler()
	}
	return nil
}

func NewOpenWeatherHandler() *WeatherHandler {
	return &WeatherHandler{Controller: *controllers.NewCurrentWeatherController(
		clients.NewOpenWeatherClient(utils.GetEnv("OPENWEATHER_API_KEY", ""), utils.GetEnv("OPENWEATHER_BASE_URL", "")))}
}

func NewGoogleWeatherHandler() *WeatherHandler {
	return &WeatherHandler{Controller: *controllers.NewCurrentWeatherController(
		clients.NewGoogleWetherClient(utils.GetEnv("GOOGLE_API_KEY", ""), utils.GetEnv("GOOGLE_BASE_URL", "")))}
}

// HandleGetCurrentWeather godoc
// @Summary      Get Current Weather
// @Description  Returns current weather for given coordinates
// @Tags         weather
// @Produce      json
// @Param        lat   query     string  true  "Latitude"    default(18.300231990440125)
// @Param        lon   query     string  true  "Longitude"   default(-64.8251590359234)
// @Success      200   {object}  responses.SuccessResponse[weather.CurrentWeather]
// @Failure      400   {object}  responses.StatusResponse
// @Failure      500   {object}  responses.StatusResponse
// @Router       /weather [get]
func (h *WeatherHandler) HandleGetCurrentWeather(c *gin.Context) {
	lat, errLat := decimal.NewFromString(c.Query("lat"))
	lon, errLon := decimal.NewFromString(c.Query("lon"))

	if errLat != nil || errLon != nil {
		c.JSON(400, responses.StatusResponse{Code: 400, Message: "invalid coordinates"})
		return
	}

	result, err := h.Controller.GetCurrentWeather(lat, lon)

	if err != nil {

		c.JSON(500, responses.StatusResponse{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, responses.SuccessResponse[weather.CurrentWeather]{Code: 200, Message: "Success", Data: result})
}

// HandleGetCurrentForecast godoc
// @Summary      Get Current Forecast
// @Description  Returns current forecast for given coordinates
// @Tags         forecast
// @Produce      json
// @Param        lat   query     string  true  "Latitude"    default(18.300231990440125)
// @Param        lon   query     string  true  "Longitude"   default(-64.8251590359234)
// @Success      200   {object}  responses.SuccessResponse[clients.ForecastResponse]
// @Failure      400   {object}  responses.StatusResponse
// @Failure      500   {object}  responses.StatusResponse
// @Router       /forecast [get]
func (h *WeatherHandler) HandleGetCurrentForecast(c *gin.Context) {
	lat, errLat := decimal.NewFromString(c.Query("lat"))
	lon, errLon := decimal.NewFromString(c.Query("lon"))

	if errLat != nil || errLon != nil {
		c.JSON(400, responses.StatusResponse{Code: 400, Message: "invalid coordinates"})
		return
	}

	result, err := h.Controller.GetCurrentForecast(lat, lon)

	if err != nil {

		c.JSON(500, responses.StatusResponse{Code: 500, Message: err.Error()})
		return
	}

	fmt.Println(result)
	c.JSON(200, responses.SuccessResponse[clients.ForecastResponse]{Code: 200, Message: "Success", Data: result})
}

// HandleGetMultipleCurrentWeather godoc
// @Summary      Get multiple current weather
// @Description  Returns current weather for multiple coordinates
// @Tags         weather
// @Accept       json
// @Produce      json
//
// @Param        request  body  []weather.CurrentWeatherRequest  true  "Array of weather requests"
//
// @Success      200  {object}  responses.SuccessResponse[[]weather.CurrentWeatherResponse]
// @Failure      400  {object}  responses.StatusResponse
// @Failure      500  {object}  responses.StatusResponse
//
// @Router      /weather [post]
func (h *WeatherHandler) HandleGetMultipleCurrentWeather(c *gin.Context) {
	var weatherReq []weather.CurrentWeatherRequest

	err := c.ShouldBindJSON(&weatherReq)
	if err != nil {
		c.JSON(400, responses.StatusResponse{Code: 400, Message: err.Error()})
		return
	}

	result, err := h.Controller.GetMultipleWeather(weatherReq)

	if err != nil {

		c.JSON(500, responses.StatusResponse{Code: 500, Message: err.Error()})
		return
	}

	fmt.Println(result)
	c.JSON(200, responses.SuccessResponse[[]weather.CurrentWeatherResponse]{Code: 200, Message: "Success", Data: result})
}

// HandleGetCurrentWeatherByCity godoc
// @Summary      Get current weather by city
// @Description  Returns current weather for a given city and country
// @Tags         weather
// @Accept       json
// @Produce      json
//
// @Param        city     query   string  true  "City name"     example(London)
// @Param        country  query   string  true  "Country code"  example(UK)
//
// @Success      200  {object}  responses.SuccessResponse[weather.CurrentWeather]
// @Failure      400  {object}  responses.StatusResponse
// @Failure      500  {object}  responses.StatusResponse
//
// @Router       /weather/city [get]
func (h *WeatherHandler) HandleGetCurrentWeatherByCity(control *controllers.CurrentLocationController) func(c *gin.Context) {
	return func(c *gin.Context) {
		city := c.Query("city")
		country := c.Query("country")

		result, err := h.Controller.GetCurrentCityWeather(city, country, control)

		if err != nil {
			c.JSON(500, responses.StatusResponse{Code: 500, Message: err.Error()})
			return
		}
		c.JSON(200, responses.SuccessResponse[weather.CurrentWeather]{Code: 200, Message: "Success", Data: result})

	}
}

// HandleGetCurrentForecastByCity godoc
// @Summary      Get current forecast by city
// @Description  Returns current forecast for a given city and country
// @Tags         forecast
// @Accept       json
// @Produce      json
//
// @Param        city     query   string  true  "City name"     example(London)
// @Param        country  query   string  true  "Country code"  example(UK)
//
// @Success      200  {object}  responses.SuccessResponse[clients.ForecastResponse]
// @Failure      400  {object}  responses.StatusResponse
// @Failure      500  {object}  responses.StatusResponse
//
// @Router       /forecast/city [get]
func (h *WeatherHandler) HandleGetCurrentForecastByCity(control *controllers.CurrentLocationController) func(c *gin.Context) {
	return func(c *gin.Context) {
		city := c.Query("city")
		country := c.Query("country")

		result, err := h.Controller.GetCurrentCityForecast(city, country, control)

		if err != nil {
			c.JSON(500, responses.StatusResponse{Code: 500, Message: err.Error()})
			return
		}
		c.JSON(200, responses.SuccessResponse[clients.ForecastResponse]{Code: 200, Message: "Success", Data: result})

	}
}
