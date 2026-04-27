package api

import (
	models2 "Lab7/models"
	"Lab7/shared/tockens/models"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"net/http"
	"strconv"
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
// @Summary Получить все сессии (требует PASETO токен)
// @Security PASETOAuth
// @Description Получить все сессии в защищённой зоне
// @Tags session
// @Accept json
// @Produce json
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/session/all [get]
func (a *App) GetAllSessions(c fiber.Ctx) error {
	sessions, err := a.db.Sessions.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.JSON(sessions)
}

// GetSessionById (в защищённой группе /api/session)
// @Summary Получить сессию(требует PASETO токен)
// @Security PASETOAuth
// @Description Получить сесси.ю в защищённой зоне
// @Tags session
// @Accept json
// @Produce json
// @Param body body models.IdSetter true "Session id"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/session/get [post]
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

// DeleteSessionById (в защищённой группе /api/candidates/:id)
// @Summary Удалить сессию (требует PASETO токен)
// @Security PASETOAuth
// @Description Удаляет сессию  по ID в защищённой зоне
// @Tags session
// @Produce json
// @Param id path int true "Candidate ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/session/{id} [delete]
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

// StartSession (в защищённой группе /api/session)
// @Summary Начать сессию
// @Security PASETOAuth
// @Tags session
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/session/{id}/start [post]
func (a *App) StartSession(c fiber.Ctx) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid id")
	}

	session, err := a.db.Sessions.FindByID(ctx, uint64(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid session id type")
	}
	if session.IsActive() {
		return c.Status(http.StatusInternalServerError).JSON("session is already started")
	}
	session.Start()
	return c.Status(http.StatusOK).JSON(fmt.Sprintf("Session is startes id = %d", session.ID))
}

// StopSession (в защищённой группе /api/session)
// @Summary Остановить сессию
// @Security PASETOAuth
// @Tags session
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/session/{id}/stop [post]
func (a *App) StopSession(c fiber.Ctx) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid id")
	}

	session, err := a.db.Sessions.FindByID(ctx, uint64(id))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid session id type")
	}
	if !session.IsActive() {
		return c.Status(http.StatusInternalServerError).JSON("session is already stopped")
	}
	session.End()
	return c.Status(http.StatusOK).JSON(fmt.Sprintf("Session is ended id = %d", session.ID))
}
