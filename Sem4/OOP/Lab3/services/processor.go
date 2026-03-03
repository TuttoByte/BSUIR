package services

import (
	"Lab3/services/classes"
	"Lab3/services/conn"
	infastruct "Lab3/services/infrastruct"
	"Lab3/services/managers"
	"Lab3/services/monitor"
	"errors"
	"fmt"
)

// =========================================================
// Файл: processor.go
// Описание: Основная бизнес-логика.
// =========================================================

type OrderProcessor struct {
	dbManager   *managers.DbManager
	connManager *managers.ConnectionManager
	logger      *monitor.CustomLogger
	calcManager *managers.Calculator
}

func NewOrderProcessor(logger *monitor.CustomLogger) *OrderProcessor {
	return &OrderProcessor{
		dbManager:   managers.NewDbManager(infastruct.NewMySQLDatabase()),
		connManager: managers.NewConnectionManager(conn.NewTelegramConnection("admin"), conn.NewSMTPConnection("10.0.0.1")),
		logger:      logger,
		calcManager: managers.NewCalculator(),
	}
}

func (op *OrderProcessor) validate(order classes.Order) error {

	if len(order.Items) == 0 {
		return errors.New("order must have at least one item")
	}
	if order.Destination.City == "" {
		return errors.New("destination city is required")
	}

	return nil
}

func (op *OrderProcessor) Process(order classes.Order) error {
	fmt.Printf("--- Processing Order %s ---\n", order.ID)

	// 1. Логика валидации
	err := op.validate(order)
	if err != nil {
		op.logger.Error("|Process|", err)
		return err
	}

	// 2. Логика расчета суммы
	total := op.calcManager.SumCalculate(order)

	// 3. Логика скидок и налогов
	total, err = order.Type.Calculate(order, total)
	if err != nil {
		op.logger.Error("|Process|", err)
		return err
	}

	// 4. Логика сохранения
	_, err = op.dbManager.AddOrderToDataBase(order, total, op.logger)
	if err != nil {
		op.logger.Error("|Process|", err)
	}

	// 5. Логика уведомлений
	msgBody := fmt.Sprintf("<h1>Your order %s is confirmed!</h1><p>Total: %.2f</p>", order.ID, total)
	op.connManager.SendOverConnection(msgBody, op.logger)

	return nil
}
