package api

import (
	models2 "Lab7/models"
	"Lab7/shared/tockens/models"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

// CreateSession (в защищённой группе /api)
// @Summary создать сессию (требует PASETO токен)
// @Security PASETOAuth
// @Description Создаёт сессию в защищённой зоне
// @Tags session
// @Accept json
// @Produce json
// @Param body body models.SessionRegisterInfo true "Session info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/session/create [post]
func (a *App) CreateSession(c fiber.Ctx) error {
	ctx := context.Background()
	regInfo := new(models.SessionRegisterInfo)

	if err := c.Bind().JSON(regInfo); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid info")
	}

	candidate, err := a.db.Candidates.FindByID(ctx, regInfo.CandidateId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid candidate id")
	}
	interwiewer, err := a.db.Interwiewers.FindByID(ctx, regInfo.InterwiwerId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid interwiewer id")
	}

	newSession := models2.NewSession(candidate, interwiewer)
	err = a.db.Sessions.Update(ctx, newSession)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("New session id = %d", newSession.ID))
}

// GetAllSessions (в защищённой группе /api/session)
// @Summary Добавить кандидата (требует PASETO токен)
// @Security PASETOAuth
// @Description Создаёт кандидата в защищённой зоне
// @Tags session
// @Accept json
// @Produce json
// @Param body body models.AnyUserInfo true "Candidate info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/session/getall [get]
func (a *App) GetAllSessions(c fiber.Ctx) error {
	ctx := context.Background()
	sessions, err := a.db.Sessions.GetAllActiveSessions(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.JSON(sessions)
}

// GetSessionById (в защищённой группе /api/session)
// @Summary Добавить кандидата (требует PASETO токен)
// @Security PASETOAuth
// @Description Создаёт кандидата в защищённой зоне
// @Tags session
// @Accept json
// @Produce json
// @Param body body models.AnyUserInfo true "Candidate info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/candidates [post]
func (a *App) GetSessionById(c fiber.Ctx) error {
	ctx := context.Background()
	idSetter := new(models.IdSetter)
	if err := c.Bind().JSON(idSetter); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid id type")
	}

	session, err := a.db.Sessions.FindByID(ctx, idSetter.Id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid session id")
	}
	return c.JSON(session)
}

// DeleteSessionById (в защищённой группе /api)
// @Summary Добавить кандидата (требует PASETO токен)
// @Security PASETOAuth
// @Description Создаёт кандидата в защищённой зоне
// @Tags session
// @Accept json
// @Produce json
// @Param body body models.AnyUserInfo true "Candidate info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/candidates [post]
func (a *App) DeleteSessionById(c fiber.Ctx) error {
	ctx := context.Background()
	idSetter := new(models.IdSetter)
	if err := c.Bind().JSON(idSetter); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid id type")
	}

	err := a.db.Sessions.Delete(ctx, idSetter.Id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return nil
}
