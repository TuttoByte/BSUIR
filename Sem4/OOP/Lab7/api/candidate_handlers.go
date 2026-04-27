package api

import (
	models2 "Lab7/models"
	"Lab7/shared/tockens/models"
	"fmt"
	"github.com/gofiber/fiber/v3"
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
// @Router /api/candidates/create/ [post]
func (a *App) AddCandidateHandler(c fiber.Ctx) error {
	creds := new(models.AnyUserInfo)
	if err := c.Bind().JSON(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}
	fmt.Println(creds)
	fmt.Println(creds)
	fmt.Println(creds)
	fmt.Println(creds)

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
