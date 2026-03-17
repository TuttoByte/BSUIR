package cargo

import "Lab4/services/contracts"

type Perishable struct {
	ExplicitCargo
}

func NewPerishable() contracts.Cargo {
	return &Perishable{
		ExplicitCargo: ExplicitCargo{
			Mass: 10,
			Cost: 100,
		},
	}
}
