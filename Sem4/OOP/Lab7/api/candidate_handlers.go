package api

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// DeleteCandidateHandler (в защищённой группе /api/candidates/:id)
// @Summary Удалить кандидата (требует PASETO токен)
// @Security PASETOAuth
// @Description Удаляет кандидата по ID в защищённой зоне
// @Tags candidates
// @Produce json
// @Param id path int true "Candidate ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/candidates/{id} [delete]
func (a *App) DeleteCandidateHandler(c fiber.Ctx) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	err = a.db.Users.Delete(ctx, uint64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

// GetAvalableSlotsHandler возвращает доступные слоты интервьюера
// @Summary Получить доступные слоты интервьюера
// @Description Возвращает список доступных слотов для записи по ID интервьюера
// @Tags slots
// @Produce json
// @Param id path int true "Interviewer ID"
// @Success 200 {array} []models.AvailabilitySlot
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/candidates/{id}/slots [get]
func (a *App) GetAvalableSlotsHandler(c fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	slots, err := a.candidates.GetSlostsByInterwiver(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(slots)
}

// CandidateBookSlotHandler бронирует слот для кандидата
// @Summary Забронировать слот
// @Security PASETOAuth
// @Description Кандидат бронирует слот по его ID
// @Tags slots
// @Produce json
// @Param CandidateId path int true "Candidate ID"
// @Param SlotId path int true "Slot ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/candidates/{candidateId}/slots/{slotId}/book [post]
func (a *App) CandidateBookSlotHandler(c fiber.Ctx) error {

	userId, err := strconv.Atoi(c.Params("CandidateId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	slotId, err := strconv.Atoi(c.Params("SlotId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	err = a.candidates.BookSlot(uint64(userId), uint64(slotId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON("slot is booked")
}

// CandidateUnbookSlotHandler отменяет бронирование слота
// @Summary Отменить бронирование слота
// @Security PASETOAuth
// @Description Кандидат отменяет бронирование слота по его ID
// @Tags slots
// @Produce json
// @Param CandidateId path int true "Candidate ID"
// @Param SlotId path int true "Slot ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/candidates/{candidateId}/slots/{slotId}/unbook [post]
func (a *App) CandidateUnbookSlotHandler(c fiber.Ctx) error {

	userId, err := strconv.Atoi(c.Params("CandidateId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	slotId, err := strconv.Atoi(c.Params("SlotId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	err = a.candidates.UnbookSlot(uint64(userId), uint64(slotId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON("slot is booked")
}
