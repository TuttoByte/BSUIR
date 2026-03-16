package contracts

type AirTransport interface {
	GetSpeed() float64
	GetPetrolium() float64
	GetMaxHight() float64
}

type WaterTransport interface {
	GetSpeed() float64
	GetPetrolium() float64
	GetMaxHoldSize() float64
}

type GroundTransport interface {
	GetSpeed() float64
	GetPetrolium() float64
	GetDurability() float64
}
