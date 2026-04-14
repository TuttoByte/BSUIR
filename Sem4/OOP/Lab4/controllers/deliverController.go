package controllers

import (
	cargo2 "Lab4/services/models/cargo"
	"Lab4/services/models/transport"
)

type DeliveryInfo struct {
	TransportType string
	Inner         string
	Distance      int
}

type CargoInfo struct {
	Amount int
	Type   string
}

type DeliverController struct {
	cargos []CargoInfo
	info   DeliveryInfo
}

func NewDeliveryController() *DeliverController {
	return &DeliverController{
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
		cargoType, err := cargo2.GetCargo(cargo.Type)
		if err != nil {
			return 0, 0, err
		}
		sum += (float64(cargo.Amount) * cargoType.GetMass()) / cargoType.GetCost()
	}

	deliveryType, err := transport.GetCargoFactory(d.info.TransportType, d.info.Inner)
	if err != nil {
		return 0, 0, err
	}

	total := sum + deliveryType.GetPetrolium()*float64(d.info.Distance)
	time := float64(d.info.Distance) / deliveryType.GetSpeed()
	return total, time, nil
}
