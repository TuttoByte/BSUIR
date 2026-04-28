package api

import (
	"Lab7/models"
	"context"
	"github.com/gofiber/fiber/v3"
	"net/http"
	"strconv"
)

// BookSlotsHandler (в защищённой группе /api/slots)
// @Summary Забронировать время
// @Security PASETOAuth
// @Tags slot
// @Produce json
// @Param id path int true "session ID"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/slots/{id}/book [post]
func (a *App) BookSlotsHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid id")
	}
	err = a.slots.BookSlot(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON("Booked")
}

// UnookSlotsHandler (в защищённой группе /api/slots)
// @Summary Освободить время
// @Security PASETOAuth
// @Tags slot
// @Produce json
// @Param id path int true "session ID"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/slots/{id}/unbook [post]
func (a *App) UnookSlotsHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid id")
	}
	err = a.slots.UnbookSlot(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON("Unbooked")
}

// DeleteSlotsByIdHandler (в защищённой группе /api/slots/:id)
// @Summary Удалить бронь времени (требует PASETO токен)
// @Security PASETOAuth
// @Tags slot
// @Produce json
// @Param id path int true "Slot ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/slots/{id} [delete]
func (a *App) DeleteSlotsByIdHandler(c fiber.Ctx) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	err = a.db.Avalavility.Delete(ctx, uint64(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return nil
}

// AddSlotHandler  (в защищённой группе /api)
// @Summary  Создать времной слот (требует PASETO токен)
// @Security PASETOAuth
// @Tags slot
// @Accept json
// @Produce json
// @Param body body models.AvailabilitySlot true "session info"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /api/slots/add [post]
func (a *App) AddSlotHandler(c fiber.Ctx) error {
	slotInfo := new(models.AvailabilitySlot)

	if err := c.Bind().JSON(slotInfo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	err := a.slots.AddSlot(slotInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON("Slot added")
}
