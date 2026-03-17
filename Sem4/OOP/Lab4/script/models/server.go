package models

import (
	"fmt"
	"strconv"
)

type Server struct {
	cargos
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

	cargoAmount := s.PrintWithWait("Enter amount of cargo")
	i, err := strconv.Atoi(cargoAmount)
	if err != nil {
		fmt.Println(err)
	}

	for _ := range i {
		s.Print(fmt.Sprintf("Cargo %d", i))
		cargoType := s.PrintWithWait("Enter type of cargo")
		cargoType := s.PrintWithWait("Enter type of cargo")

	}

	cargoType := s.PrintWithWait("Enter type of cargo")
	distance := s.PrintWithWait("Enter distance to the point of the destination")
	transportType := s.PrintWithWait("Enter type of transport")

}
