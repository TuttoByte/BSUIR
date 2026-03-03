package services

import "fmt"

// =========================================================
// Файл: staff.go
// Описание: Система управления персоналом склада.
// =========================================================

type WarehouseWorker interface {
	DoWork()
}

type ProccessWorker interface {
	ProcessOrder()
}

type ProccessAttendor interface {
	AttendMeeting()
}

type ProccessSwinging interface {
	SwingingTheLead()
}

type ProcessResting interface {
	GetRest()
}

// HumanManager - Человек
type HumanManager struct{}

func (h HumanManager) ProcessOrder() {
	fmt.Println("Manager is processing logic...")
}

func (h HumanManager) AttendMeeting() {
	fmt.Println("Manager is boring at the meeting...")
}

func (h HumanManager) GetRest() {
	fmt.Println("Manager is taking a break...")
}

func (h HumanManager) SwingingTheLead() {
	fmt.Println("Manager is watching reels...")
}
func (h HumanManager) DoWork() {
	h.ProcessOrder()
	h.AttendMeeting()
	h.SwingingTheLead()
	h.GetRest()
}

// RobotPacker - Робот
type RobotPacker struct {
	Model string
}

func (r RobotPacker) ProcessOrder() {
	fmt.Println("Robot " + r.Model + " is packing boxes...")
}

func (r RobotPacker) GetRest() {
	fmt.Println("Robot was taken for maintenance")
}

func (r RobotPacker) DoWork() {
	r.ProcessOrder()
	r.GetRest()
}

// ManageWarehouse - функция, которая работает со списком работников
func ManageWarehouse(workers []WarehouseWorker) {
	fmt.Println("\n--- Warehouse Shift Started ---")
	for _, worker := range workers {
		worker.DoWork()
	}
}
