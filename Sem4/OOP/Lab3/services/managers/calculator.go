package managers

import "Lab3/services/classes"

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) SumCalculate(order classes.Order) float64 {
	var total float64
	for _, item := range order.Items {
		total += item.Price
	}
	return total
}
