package api

import (
	models2 "Lab7/models"
	"Lab7/shared/tockens/models"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// AddCandidateHandler (в защищённой группе /api/candidates/create)
// @Summary Добавить кандидата (требует PASETO токен)
// @Security PASETOAuth
// @Description Создаёт кандидата в защищённой зоне
// @Tags candidates
// @Accept json
// @Produce json
// @Param body body models.AnyUserInfo true "Candidate info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/candidates/add [post]
func (a *App) AddCandidateHandler(c fiber.Ctx) error {
	creds := new(models.AnyUserInfo)
	if err := c.Bind().JSON(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	candidate := &models2.Candidate{
		Name:  creds.Username,
		Email: creds.Email,
	}

	err := a.db.Candidates.Update(a.ctx, candidate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusCreated).JSON(fmt.Sprintf("User successfully added with id = %d", candidate.ID))
}

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

	err = a.db.Candidates.Delete(ctx, uint64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
