package transport

import "Lab4/services/contracts"

type Helicopter struct {
	MaxSpeed   float64
	PetrolCons float64
	MaxHight   float64
}

// Fabric implementaton

func NewHelicopter() contracts.AirTransport {
	return &Helicopter{
		250,
		200,
		6000,
	}
}

func (h *Helicopter) GetSpeed() float64 {
	return h.MaxSpeed
}
func (h *Helicopter) GetPetrolium() float64 {
	return h.PetrolCons
}

func (h *Helicopter) GetMaxHight() float64 {
	return h.MaxHight
}
