package controllers

import (
	"main/clients"
	"main/models/location"
)

type CurrentLocationController struct {
	Client clients.CoordinateDataCleint
}

func NewCurrentLocationController(client clients.CoordinateDataCleint) *CurrentLocationController {
	return &CurrentLocationController{
		Client: client,
	}
}

func (cl *CurrentLocationController) GetLocation(name string) (location.CurrentLocation, error) {
	lat, lon, err := cl.Client.GetCurrentLocation(name)
	if err != nil {
		return location.CurrentLocation{}, err
	}

	return location.CurrentLocation{
		Latitude:  lat,
		Longitude: lon,
	}, nil

}
