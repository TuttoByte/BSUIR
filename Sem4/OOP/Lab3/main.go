package main

import (
	"Lab3/services"
	"Lab3/services/classes"
	"Lab3/services/monitor"
	"fmt"
	"log"
)

// =========================================================
// Файл: main.go
// Описание: Точка входа в приложение.
// =========================================================

func main() {

	logger := monitor.NewCustomLogger()
	// 1. Создание заказа
	order := classes.Order{
		ID:   "ORD-256-X",
		Type: &classes.PremiumTax{},
		Items: []classes.Item{
			{ID: "1", Name: "Thermal Clips", Price: 1500},
			{ID: "2", Name: "UNATCO Pass Card", Price: 50},
		},
		ClientEmail: "jeevacation@gmail.com",
		Destination: classes.Address{City: "Agartha", Street: "33 Thomas Street", Zip: "[REDACTED]"},
		Discount:    classes.NewGolden(),
	}

	// 2. Инициализация процессора
	processor := services.NewOrderProcessor(logger)

	// 3. Обработка заказа
	if err := processor.Process(order); err != nil {
		log.Fatalf("Failed to process order: %v", err)
	}

	// 4. Работа с обслуживанием
	fmt.Println("\nTesting Warehouse Stuff:")
	workers := []services.WarehouseWorker{
		services.HumanManager{},
		services.RobotPacker{Model: "George Droid"},
	}

	services.ManageWarehouse(workers)
}
