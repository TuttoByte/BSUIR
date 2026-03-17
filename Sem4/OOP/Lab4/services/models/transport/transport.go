package transport

import (
	"Lab4/services/contracts"
	"errors"
)

func GetAirTransport(atype string) (contracts.AirTransport, error) {
	switch atype {
	case "helicopter":
		return NewHelicopter(), nil
	case "plane":
		return NewPlane(), nil
	}
	return nil, errors.New("invalid air transport")
}

func GetWaterTransport(atype string) (contracts.WaterTransport, error) {
	switch atype {
	case "tanker":
		return NewTanker(), nil
	}

	return nil, errors.New("invalid whater transport")
}

func GetGroundTransport(atype string) (contracts.GroundTransport, error) {
	switch atype {
	case "train":
		return NewTrain(), nil
	case "truck":
		return NewTruck(), nil
	}
	return nil, errors.New("invalid ground transport")
}

func GetCargoFactory(trType string, inner string) (contracts.ITransportFactory, error) {
	switch trType {

	case "air":
		return GetAirTransport(inner)

	case "whater":
		return GetWaterTransport(inner)

	case "ground":
		return GetGroundTransport(inner)
	}

	return nil, errors.New(trType + " is not a cargo transport")
}
