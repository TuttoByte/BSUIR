package classes

// =========================================================
// Файл: models.go
// Описание: Модели данных системы.
// =========================================================

// Item - товар в заказе
type Item struct {
	ID    string
	Name  string
	Price float64
}

// Address - адрес доставки
type Address struct {
	City   string
	Street string
	Zip    string
}

// Order - заказ
type Order struct {
	ID                    string
	Items                 []Item
	Type                  Tax
	ClientEmail           string
	ClientTelegramTockern string
	Destination           Address
	Discount              Discout
}

type NewBye struct {
	discount float64
}

func (n *NewBye) Get(amount float64) float64 {
	return n.discount * amount
}

func NewNewBye() *NewBye {
	return &NewBye{discount: 0}
}

type GoldenCard struct {
	discount float64
}

func (g *GoldenCard) Get(amount float64) float64 {
	return g.discount * amount
}

func NewGolden() *GoldenCard {
	return &GoldenCard{discount: 0.15}
}

type SilverCard struct {
	discount float64
}

func (s *SilverCard) Get(amount float64) float64 {
	return s.discount * amount
}

func NewSilverCArd() *SilverCard {
	return &SilverCard{discount: 0.1}
}

type Discout interface {
	Get(amount float64) float64
}
