package cargo

type Electro struct {
	ExplicitCargo
}

func NewElectro() *Electro {
	return &Electro{
		ExplicitCargo: ExplicitCargo{
			Mass: 1.5,
			Cost: 50,
		},
	}
}
