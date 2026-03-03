package classes

import (
	"errors"
)

type Tax interface {
	Calculate(Order, float64) (float64, error)
}

type StandardTax struct {
}
type PremiumTax struct {
}
type BudjetTax struct {
}
type InternationalTax struct {
}

func (s *StandardTax) Calculate(order Order, amount float64) (float64, error) {
	return amount * 1.2, nil
}
func (p *PremiumTax) Calculate(order Order, amount float64) (float64, error) {
	return (amount * 0.9) * 1.2, nil
}
func (b *BudjetTax) Calculate(order Order, amount float64) (float64, error) {
	if len(order.Items) > 3 {
		return 0, errors.New("Budget orders cannot have more than 3 items. Skipping.")
	}
	return amount, nil
}

func (i *InternationalTax) Calculate(order Order, amount float64) (float64, error) {
	if order.Destination.City == "Nowhere" {
		return 0, errors.New("cannot ship to Nowhere")
	}
	amount = amount * 1.5
	return amount, nil
}
