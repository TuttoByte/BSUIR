package Transport

type Train struct {
	MaxSpeed   float64
	PetrolCons float64
	Stops      int
}

func (t *Train) GetSpeed() float64 {
	return t.MaxSpeed
}
func (t *Train) GetPetrolium() float64 {
	return t.PetrolCons
}
func (t *Train) GetDurability() float64 {
	return float64(t.Stops * 1000)
}
