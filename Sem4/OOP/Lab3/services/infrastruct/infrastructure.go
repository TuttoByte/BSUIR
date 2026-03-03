package infastruct

import (
	"fmt"
	"os"
	"time"
)

// =========================================================
// Файл: infrastructure.go
// Описание: Имитация работы с БД и внешними сервисами.
// =========================================================

// RandomSQLDatabase - имитация тяжелой базы данных
type RandomSQLDatabase struct {
	ConnectionString string
}

func NewMySQLDatabase() *RandomSQLDatabase {
	return &RandomSQLDatabase{ConnectionString: "random://root:password@localhost:228/shop"}
}

// Сохранение заказа в "базу данных"
func (db *RandomSQLDatabase) SaveOrder(order interface{}, total float64) error {
	fmt.Println("Connecting to RandomSQL at", db.ConnectionString, "...")
	time.Sleep(500 * time.Millisecond) // Имитация задержки сети

	file, err := os.OpenFile("orders_db.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	record := fmt.Sprintf("[%s] ID: %s | Type: %s | Total: %.2f\n", time.Now().Format(time.RFC3339), order, total)
	if _, err := file.WriteString(record); err != nil {
		return err
	}
	fmt.Println("Order saved successfully.")
	return nil
}
