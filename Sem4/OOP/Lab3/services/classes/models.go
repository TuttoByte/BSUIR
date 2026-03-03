package classes

import infastruct "Lab3/services/infrastruct"

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
	Discount              infastruct.Discout
}
