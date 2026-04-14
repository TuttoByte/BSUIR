package main

import (
	"log"
	"main/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main/docs"
)

// @title           Wheather Example API
// @version         1.0
// @BasePath  /api/v1

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	currentWeatherHandler := api.NewCurrentWeatherHandler("open")
	currentLocationHandler := api.NewOpenWeatherCoordinatesHandler()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/weather", currentWeatherHandler.HandleGetCurrentWeather)
		v1.GET("/forecast", currentWeatherHandler.HandleGetCurrentForecast)
		v1.POST("/weather", currentWeatherHandler.HandleGetMultipleCurrentWeather)
		v1.GET("/location", currentLocationHandler.HandleGetCurrentCityCoord)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
