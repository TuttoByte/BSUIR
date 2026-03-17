package controllers

import (
	"Lab4/services/contracts"
	"Lab4/services/factories"
)

type DeliveryInfo struct {
	TransportType string
	distance      float64
}

type CargoInfo struct {
	Amount int
	Type   string
}

type DeliverController struct {
	cargos []CargoInfo
	info   DeliveryInfo
}

func (d *DeliverController) NewDeliveryController() DeliverController {
	return DeliverController{
		cargos: make([]CargoInfo, 0),
	}
}

func (d *DeliverController) AddCargo(cargo CargoInfo) {
	d.cargos = append(d.cargos, cargo)
}

func (d *DeliverController) SetDeliveryInfo(info DeliveryInfo) {
	d.info = info
}

func (d *DeliverController) GetDeliveyResult() (float64, float64, error) {
	sum := 0.0

	for _, cargo := range d.cargos {
		cargoType, err := factories.GetCargo(cargo.Type)
		if err != nil {
			return 0, 0, err
		}
		sum += (float64(cargo.Amount) * cargoType.GetMass()) / cargoType.GetCost()
	}

	deliveryType, err := contracts.GetCargoFactory(d.info.TransportType, "")
	if err != nil {
		return 0, 0, err
	}

	total := sum + deliveryType.GetPetrolium()*d.info.distance
	time := d.info.distance / deliveryType.GetSpeed()
	return total, time, nil
}
