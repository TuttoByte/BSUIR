package controllers

import (
	"Lab7/clients/db"
	"Lab7/models"
	"context"
)

type AvalTimeController struct {
	avalTime *db.AvalabilytyDB
}

func NewAvalTimeController(avalTime *db.AvalabilytyDB) *AvalTimeController {
	return &AvalTimeController{
		avalTime: avalTime,
	}
}

func (a *AvalTimeController) AddSlot(slot *models.AvailabilitySlot) error {
	ctx := context.Background()

	err := a.avalTime.Update(ctx, slot)
	if err != nil {
		return err
	}
	return nil
}

func (a *AvalTimeController) DeleteSlot(id uint64) error {
	ctx := context.Background()

	err := a.avalTime.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AvalTimeController) BookSlot(id uint64) error {
	ctx := context.Background()
	slot, err := a.avalTime.FindByID(ctx, id)
	if err != nil {
		return err
	}
	slot.IsBooked = true
	return nil
}

func (a *AvalTimeController) UnbookSlot(id uint64) error {
	ctx := context.Background()
	slot, err := a.avalTime.FindByID(ctx, id)
	if err != nil {
		return err
	}
	slot.IsBooked = false
	return nil
}

func (a *AvalTimeController) GetAllSlots() ([]models.AvailabilitySlot, error) {
	slots, err := a.avalTime.GetAll()
	if err != nil {
		return nil, err
	}
	return slots, nil
}
