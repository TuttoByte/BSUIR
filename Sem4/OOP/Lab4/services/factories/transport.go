package factories

import (
	"Lab4/services/contracts"
	"Lab4/services/models/transport"
	"errors"
)

func GetAirTransport(atype string) (contracts.AirTransport, error) {
	switch atype {
	case "helicopter":
		return &transport.Helicopter{}, nil
	case "plane":
		return &transport.Plane{}, nil
	}
	return nil, errors.New("invalid air transport")
}

func GetWaterTransport(atype string) (contracts.WaterTransport, error) {
	switch atype {
	case "tanker":
		return &transport.Tanker{}, nil
	}

	return nil, errors.New("invalid whater transport")
}

func GetGroundTransport(atype string) (contracts.GroundTransport, error) {
	switch atype {
	case "train":
		return &transport.Train{}, nil
	case "truck":
		return &transport.Truck{}, nil
	}
	return nil, errors.New("invalid ground transport")
}
