package controllers

import (
	"main/clients"
)

type CurrentLocationController struct {
	Client clients.CoordinateDataCleint
}

func NewCurrentLocationController(client clients.CoordinateDataCleint) *CurrentLocationController {
	return &CurrentLocationController{
		Client: client,
	}
}

func (cl *CurrentLocationController) GetCurentLocation(name string) ([]clients.OpenWeatherCityInfo, error) {
	data, err := cl.Client.GetLocation(name)
	if err != nil {
		return nil, err
	}

	return data, nil
}
