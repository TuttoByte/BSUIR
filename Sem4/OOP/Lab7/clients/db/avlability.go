package db

import (
	"Lab7/models"
)

type AvalabilytyDB struct {
	*GormDB[models.AvailabilitySlot]
}

func NewAvalabilytyDB() (*AvalabilytyDB, error) {
	newGorm, err := NewGormDB[models.AvailabilitySlot]()
	if err != nil {
		return nil, err
	}
	return &AvalabilytyDB{
		newGorm,
	}, nil
}

func (a *AvalabilytyDB) GetAll() ([]models.AvailabilitySlot, error) {
	var slots []models.AvailabilitySlot
	a.db.Find(&slots)
	return slots, nil
}
