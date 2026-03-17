package transport

import "Lab4/services/contracts"

type Plane struct {
	MaxSpeed   float64
	PetrolCons float64
	MaxHight   float64
}

// Fabric implementaton

func NewPlane() contracts.AirTransport {
	return &Plane{
		MaxSpeed:   850,
		PetrolCons: 150,
		MaxHight:   10000,
	}
}

func (p *Plane) GetSpeed() float64 {
	return p.MaxSpeed
}
func (p *Plane) GetPetrolium() float64 {
	return p.PetrolCons
}

func (p *Plane) GetMaxHight() float64 {
	return p.MaxHight
}
