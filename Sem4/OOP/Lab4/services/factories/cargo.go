package factories

import (
	"Lab4/services/contracts"
	"Lab4/services/models/cargo"
	"errors"
)

func GetCargo(ctype string) (contracts.Cargo, error) {

	switch ctype {
	case "cloth":
		return cargo.NewCloth(), nil
	case "equipment":
		return cargo.NewEquipment(), nil
	case "electro":
		return cargo.NewElectro(), nil
	case "perishable":
		return cargo.NewPerishable(), nil
	}

	return nil, errors.New("no cargo found")
}
