package cargo

import "Lab4/services/contracts"

type Equipment struct {
	ExplicitCargo
}

func NewEquipment() contracts.Cargo {
	return &Equipment{
		ExplicitCargo: ExplicitCargo{
			Mass: 120,
			Cost: 15,
		},
	}
}
