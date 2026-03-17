package transport

import "Lab4/services/contracts"

type Truck struct {
	MaxSpeed       float64
	PetrolCons     float64
	RubberCapacity float64
}

func NewTruck() contracts.GroundTransport {
	return &Truck{
		MaxSpeed:       80,
		PetrolCons:     15,
		RubberCapacity: 500,
	}
}

func (t *Truck) GetSpeed() float64 {
	return t.MaxSpeed
}
func (t *Truck) GetPetrolium() float64 {
	return t.PetrolCons
}

func (t *Truck) GetDurability() float64 {
	return t.RubberCapacity
}
