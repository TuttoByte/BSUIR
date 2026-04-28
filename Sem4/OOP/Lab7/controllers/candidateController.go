package controllers

import (
	"Lab7/clients/db"
	"Lab7/models"
	"context"
)

type CandidateController struct {
	users db.UserRepository
	slots db.SlotRepository
}

func NewCandidateController(users db.UserRepository, slots db.SlotRepository) *CandidateController {
	return &CandidateController{users, slots}
}

func (c *CandidateController) GetSlostsByInterwiver(id uint64) ([]models.AvailabilitySlot, error) {
	slots, err := c.slots.GetAll()
	if err != nil {
		return nil, err
	}

	var avalableSlots []models.AvailabilitySlot
	for _, slot := range slots {
		if slot.InterviewerID == id {
			avalableSlots = append(avalableSlots, slot)
		}
	}
	return avalableSlots, nil
}

func (c *CandidateController) BookSlot(userId, slotId uint64) error {
	ctx := context.Background()
	_, err := c.users.FindByID(ctx, userId)
	if err != nil {
		return err
	}
	slot, err := c.slots.FindByID(ctx, slotId)
	if err != nil {
		return err
	}
	slot.IsBooked = true
	slot.CandidateID = userId

	err = c.slots.Update(ctx, &slot)
	if err != nil {
		return err
	}
	return nil
}

func (c *CandidateController) UnbookSlot(userId, slotId uint64) error {
	ctx := context.Background()
	_, err := c.users.FindByID(ctx, userId)
	if err != nil {
		return err
	}
	slot, err := c.slots.FindByID(ctx, slotId)
	if err != nil {
		return err
	}
	slot.IsBooked = false
	slot.CandidateID = 0

	err = c.slots.Update(ctx, &slot)
	if err != nil {
		return err
	}
	return nil
}
