package cargo

type ExplicitCargo struct {
	Mass float64
	Cost float64
}

func (c *ExplicitCargo) GetMass() float64 {
	return c.Mass
}

func (c *ExplicitCargo) GetCost() float64 {
	return c.Cost
}
