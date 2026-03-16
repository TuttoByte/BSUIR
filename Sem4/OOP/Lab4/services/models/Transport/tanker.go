package Transport

import "Lab4/services/contracts"

type Tanker struct {
	MaxSpeed        float64
	PetrolCons      float64
	MaxHoldCapacity float64
}

func NewTanker() contracts.WaterTransport {
	return &Tanker{
		35,
		3,
		2000,
	}
}

func (t *Tanker) GetSpeed() float64 {
	return t.MaxSpeed
}
func (t *Tanker) GetPetrolium() float64 {
	return t.PetrolCons
}
func (t *Tanker) GetMaxHoldSize() float64 {
	return t.MaxHoldCapacity
}
