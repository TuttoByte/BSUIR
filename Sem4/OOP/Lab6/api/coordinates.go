package api

import (
	"github.com/gin-gonic/gin"
	"main/clients"
	"main/controllers"
	"main/shared/responses"
	"main/shared/utils"
)

type CoordintaesHandler struct {
	Controller controllers.CurrentLocationController
}

func NewOpenWeatherCoordinatesHandler() *CoordintaesHandler {
	return &CoordintaesHandler{Controller: *controllers.NewCurrentLocationController(
		clients.NewOpenWeatherCoords(utils.GetEnv("OPENWEATHER_API_KEY", ""), utils.GetEnv("OPENWEATHER_BASE_UNDER_URL", "")))}
}

// HandleGetCurrentCityCoord godoc
// @Summary      Get city coordinates
// @Description  Returns coordinates for a city by name
// @Tags         coordinates
// @Accept       json
// @Produce      json
//
// @Param        name   query   string  true  "City name"  example(London)
//
// @Success      200  {object}  responses.SuccessResponse[[]clients.OpenWeatherCityInfo]
// @Failure      500  {object}  responses.StatusResponse
//
// @Router       /location [get]
func (h *CoordintaesHandler) HandleGetCurrentCityCoord(c *gin.Context) {
	name := c.Query("name")
	result, err := h.Controller.GetCurentLocation(name)
	if err != nil {
		c.JSON(500, responses.StatusResponse{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, responses.SuccessResponse[[]clients.OpenWeatherCityInfo]{Code: 200, Message: "Success", Data: result})
}
