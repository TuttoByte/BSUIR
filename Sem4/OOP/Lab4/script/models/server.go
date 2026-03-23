package models

import (
	"Lab4/controllers"
	"fmt"
	"strconv"
)

type Server struct {
	dcontrol *controllers.DeliverController
}

func NewServer() *Server {
	return &Server{
		dcontrol: controllers.NewDeliveryController(),
	}
}

func (s *Server) Print(msg string) {
	fmt.Println(msg)
}
func (s *Server) PrintWithWait(msg string) string {
	var input string
	fmt.Println(msg)
	_, _ = fmt.Scanln(&input)
	return input
}

func (s *Server) Show() {

	cargoAmount := s.PrintWithWait("Enter amount of batches ")
	i, err := strconv.Atoi(cargoAmount)
	if err != nil {
		fmt.Println(err)
	}

	for range i {
		s.Print(fmt.Sprintf("Cargo %d", i))
		cargoNumber := s.PrintWithWait("Enter number of cargos")
		cargoType := s.PrintWithWait("Enter type of cargo")

		amount, _ := strconv.Atoi(cargoNumber)
		s.dcontrol.AddCargo(controllers.CargoInfo{
			amount,
			cargoType,
		})

	}

	distance := s.PrintWithWait("Enter distance to the point of the destination")
	transportType := s.PrintWithWait("Enter type of transport")
	inner := s.GetTransportInfo(transportType)

	distanceFloat, _ := strconv.ParseFloat(distance, 64)

	s.dcontrol.SetDeliveryInfo(controllers.DeliveryInfo{
		transportType,
		inner, distanceFloat,
	})

	s.Print("The cost of delivery: ")

	total, time, err := s.dcontrol.GetDeliveyResult()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Total delivery cost: ", total)
	fmt.Println("Time delivery cost: ", time)

}

func (s *Server) GetTransportInfo(trtype string) string {
	fmt.Println("Print one of the variant: ")
	switch trtype {
	case "whater":
		fmt.Println("1. tanker")
	case "air":
		fmt.Println("1. plane")
		fmt.Println("2. helicopter")
	case "ground":
		fmt.Println("1. train")
		fmt.Println("2. truck")
	default:
		fmt.Println("Non type in")
		return ""
	}

	var input string
	_, _ = fmt.Scanln(&input)
	return input
}
