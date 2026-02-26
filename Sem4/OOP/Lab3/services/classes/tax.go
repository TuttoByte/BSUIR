package classes

type Tax interface {
	Calculate(float64) float64
}

type StandardTax struct {
}

func (s *StandardTax) Calculate(amount float64) float64 {
	return amount * 1.2
}
