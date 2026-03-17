package cargo

import (
	"Lab4/services/contracts"
	"errors"
)

func GetCargo(ctype string) (contracts.Cargo, error) {

	switch ctype {
	case "cloth":
		return NewCloth(), nil
	case "equipment":
		return NewEquipment(), nil
	case "electro":
		return NewElectro(), nil
	case "perishable":
		return NewPerishable(), nil
	}

	return nil, errors.New("no cargo found")
}
