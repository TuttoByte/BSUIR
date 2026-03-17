package contracts

type Cargo interface {
	GetMass() float64
	GetCost() float64
}
