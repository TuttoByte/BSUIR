package cargo

import "Lab4/services/contracts"

type Cloth struct {
	ExplicitCargo
}

func NewCloth() contracts.Cargo {
	return &Cloth{
		ExplicitCargo: ExplicitCargo{
			Mass: 0.8,
			Cost: 20,
		},
	}
}
