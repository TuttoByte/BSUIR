package factories

import (
	"Lab4/services/contracts"
	"Lab4/services/models/Transport"
	"errors"
)

func GetAirTransport(atype string) (contracts.AirTransport, error) {
	switch atype {
	case "helicopter":
		return &Transport.Helicopter{}, nil
	case "plane":
		return &Transport.Plane{}, nil
	}
	return nil, errors.New("invalid air transport")
}

func GetWaterTransport(atype string) (contracts.WaterTransport, error) {
	switch atype {
	case "tanker":
		return &Transport.Tanker{}, nil
	}

	return nil, errors.New("invalid whater transport")
}

func GetGroundTransport(atype string) (contracts.GroundTransport, error) {
	switch atype {
	case "train":
		return &Transport.Train{}, nil
	case "truck":
		return &Transport.Truck{}, nil
	}
	return nil, errors.New("invalid ground transport")
}
