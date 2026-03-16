package contracts

import (
	"Lab4/services/factories"
	"errors"
)

type ITransportFactory interface {
	GetSpeed() float64
	GetPetrolium() float64
}

func GetCargoFactory(trType string, inner string) (ITransportFactory, error) {
	switch trType {

	case "air":
		return factories.GetAirTransport(inner)

	case "whater":
		return factories.GetWaterTransport(inner)

	case "ground":
		return factories.GetGroundTransport(inner)
	}

	return nil, errors.New(trType + " is not a cargo transport")
}
