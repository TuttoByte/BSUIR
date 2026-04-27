package api

import (
	"Lab7/models"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// AddProblemHandler
// @Summary Добавить задачу (требует PASETO токен)
// @Security PASETOAuth
// @Description Создаёт новую задачу для интервью
// @Tags problems
// @Accept json
// @Produce json
// @Param body body models.InterwieweProblem true "Problem data"
// @Success 201 {string} string "Problem added"
// @Failure 400 {object} map[string]string
// @Failure 500 {string} string
// @Router /api/problems/add [post]
func (a *App) AddProblemHandler(c fiber.Ctx) error {
	problem := new(models.InterwieweProblem)

	if err := c.Bind().JSON(problem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := a.promlems.AddNewProblem(problem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Unable to add new problem")
	}

	return c.Status(fiber.StatusCreated).JSON("Problem added")
}

// DeleteProblemHandler
// @Summary Удалить задачу (требует PASETO токен)
// @Security PASETOAuth
// @Description Удаляет задачу по ID
// @Tags problems
// @Produce json
// @Param id path int true "Problem ID"
// @Success 200 {string} string "Problem deleted"
// @Failure 400 {object} map[string]string
// @Failure 500 {string} string
// @Router /api/problems/{id} [delete]
func (a *App) DeleteProblemHandler(c fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = a.promlems.DeleteProblem(uint64(idInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Unable to delete problem")
	}
	return c.Status(fiber.StatusOK).JSON("Problem deleted")
}

// GetAllProblems
// @Summary Получить все задачи (требует PASETO токен)
// @Security PASETOAuth
// @Description Возвращает список всех задач
// @Tags problems
// @Produce json
// @Success 200 {array} models.InterwieweProblem
// @Failure 500 {string} string
// @Router /api/problems/all [get]
func (a *App) GetAllProblems(c fiber.Ctx) error {
	problems, err := a.promlems.GetAllProblems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Unable to get all problems")
	}
	return c.Status(fiber.StatusOK).JSON(problems)
}

// GetProblemByIdHandler
// @Summary Получить задачу по ID (требует PASETO токен)
// @Security PASETOAuth
// @Description Возвращает задачу по её ID
// @Tags problems
// @Produce json
// @Param id path int true "Problem ID"
// @Success 200 {object} models.InterwieweProblem
// @Failure 400 {object} map[string]string
// @Failure 500 {string} string
// @Router /api/problems/{id} [get]
func (a *App) GetProblemByIdHandler(c fiber.Ctx) error {
	problemId := c.Params("id")
	problemIdInt, err := strconv.Atoi(problemId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	problem, err := a.promlems.GetProblemById(uint64(problemIdInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Unable to get problem")
	}
	return c.Status(fiber.StatusOK).JSON(problem)
}
