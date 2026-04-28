package api

import (
	"Lab7/controllers"
	models2 "Lab7/models"
	tokenmodels "Lab7/shared/tockens/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSessionService struct {
	mock.Mock
}

func (m *MockSessionService) AddSession(candidate models2.User, interviewer models2.User) error {
	args := m.Called(candidate, interviewer)
	return args.Error(0)
}

func (m *MockSessionService) GetAllSessions() ([]models2.Session, error) {
	args := m.Called()
	return args.Get(0).([]models2.Session), args.Error(1)
}

func (m *MockSessionService) GetSessionById(id uint64) (models2.Session, error) {
	args := m.Called(id)
	return args.Get(0).(models2.Session), args.Error(1)
}

func (m *MockSessionService) DeleteSession(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSessionService) StartSession(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSessionService) EndSession(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSessionService) AddProblems(problems []uint64, sessionId uint64) error {
	args := m.Called(problems, sessionId)
	return args.Error(0)
}

func (m *MockSessionService) SetSessionResult(info *models2.ResultIndo) error {
	args := m.Called(info)
	return args.Error(0)
}

type MockUserRepository struct {
	mock.Mock
}

// Repository[models.User]
func (m *MockUserRepository) Create(ctx context.Context) error {
	return m.Called(ctx).Error(0)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id uint64) (models2.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models2.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, item *models2.User) error {
	return m.Called(ctx, item).Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id uint64) error {
	return m.Called(ctx, id).Error(0)
}

// UserRepository
func (m *MockUserRepository) GetByName(ctx context.Context, name string) (models2.User, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(models2.User), args.Error(1)
}

func (m *MockUserRepository) GetPassHash(name string) (string, error) {
	args := m.Called(name)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) GetRole(name string) (string, error) {
	args := m.Called(name)
	return args.String(0), args.Error(1)
}

// ===== ФИКСТУРЫ =====

func mockUser(id uint64, name string) models2.User {
	return models2.User{
		ID:   id,
		Name: name,
		Role: "candidate",
	}
}

func mockSession() models2.Session {
	return models2.Session{
		ID:            1,
		CandidateId:   1,
		InterviewerId: 2,
		TimeInfo: models2.TimeInfo{
			StartTime: time.Now(),
		},
	}
}

// ===== ХЕЛПЕР =====

func setupSessionApp(sessionSvc *MockSessionService, userDB *MockUserRepository) *fiber.App {
	app := fiber.New()
	a := &App{
		session: sessionSvc,
		db: &controllers.Dbontroller{
			Users: userDB,
		},
	}

	app.Post("/api/session/create", a.CreateSession)
	app.Get("/api/session/all", a.GetAllSessions)
	app.Post("/api/session/get", a.GetSessionById)
	app.Delete("/api/session/:id", a.DeleteSessionById)
	app.Post("/api/session/:id/start", a.StartSession)
	app.Post("/api/session/:id/stop", a.StopSession)
	app.Post("/api/session/problems", a.AddProblemsToSession)
	app.Post("/api/session/result", a.SetResultHandler)

	return app
}

func toJSON(t *testing.T, v any) *bytes.Buffer {
	b, err := json.Marshal(v)
	assert.NoError(t, err)
	return bytes.NewBuffer(b)
}

// ===== CreateSession =====

func TestCreateSession_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	candidate := mockUser(1, "candidate")
	interviewer := mockUser(2, "interviewer")

	userDB.On("FindByID", mock.Anything, uint64(1)).Return(candidate, nil)
	userDB.On("FindByID", mock.Anything, uint64(2)).Return(interviewer, nil)
	sessionSvc.On("AddSession", candidate, interviewer).Return(nil)

	app := setupSessionApp(sessionSvc, userDB)

	body := toJSON(t, tokenmodels.SessionRegisterInfo{CandidateId: 1, InterwiwerId: 2})
	req := httptest.NewRequest(http.MethodPost, "/api/session/create", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	sessionSvc.AssertExpectations(t)
}

func TestCreateSession_InvalidBody(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)
	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/create", bytes.NewBufferString("invalid"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateSession_InvalidCandidateID(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	userDB.On("FindByID", mock.Anything, uint64(99)).Return(models2.User{}, errors.New("not found"))

	app := setupSessionApp(sessionSvc, userDB)

	body := toJSON(t, tokenmodels.SessionRegisterInfo{CandidateId: 99, InterwiwerId: 2})
	req := httptest.NewRequest(http.MethodPost, "/api/session/create", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// ===== GetAllSessions =====

func TestGetAllSessions_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessions := []models2.Session{mockSession()}
	sessionSvc.On("GetAllSessions").Return(sessions, nil)

	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodGet, "/api/session/all", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var result []models2.Session
	json.Unmarshal(body, &result)
	assert.Len(t, result, 1)
	assert.Equal(t, uint64(1), result[0].ID)
}

func TestGetAllSessions_ServiceError(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("GetAllSessions").Return([]models2.Session{}, errors.New("db error"))

	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodGet, "/api/session/all", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

// ===== GetSessionById =====

func TestGetSessionById_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("GetSessionById", uint64(1)).Return(mockSession(), nil)

	app := setupSessionApp(sessionSvc, userDB)

	body := toJSON(t, tokenmodels.IdSetter{Id: 1})
	req := httptest.NewRequest(http.MethodPost, "/api/session/get", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetSessionById_NotFound(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("GetSessionById", uint64(99)).Return(models2.Session{}, errors.New("not found"))

	app := setupSessionApp(sessionSvc, userDB)

	body := toJSON(t, tokenmodels.IdSetter{Id: 99})
	req := httptest.NewRequest(http.MethodPost, "/api/session/get", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// ===== DeleteSessionById =====

func TestDeleteSessionById_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("DeleteSession", uint64(1)).Return(nil)

	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodDelete, "/api/session/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteSessionById_InvalidID(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)
	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodDelete, "/api/session/abc", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	sessionSvc.AssertNotCalled(t, "DeleteSession")
}

// ===== StartSession =====

func TestStartSession_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("StartSession", uint64(1)).Return(nil)

	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/1/start", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestStartSession_InvalidID(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)
	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/abc/start", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	sessionSvc.AssertNotCalled(t, "StartSession")
}

func TestStartSession_ServiceError(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("StartSession", uint64(1)).Return(errors.New("already started"))

	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/1/start", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

// ===== StopSession =====

func TestStopSession_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("EndSession", uint64(1)).Return(nil)

	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/1/stop", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestStopSession_ServiceError(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("EndSession", uint64(1)).Return(errors.New("not started"))

	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/1/stop", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestAddProblemsToSession_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("AddProblems", []uint64{1, 2}, uint64(1)).Return(nil)

	app := setupSessionApp(sessionSvc, userDB)

	body := toJSON(t, models2.AddProblemsInfo{SessionId: 1, Problems: []uint64{1, 2}})
	req := httptest.NewRequest(http.MethodPost, "/api/session/problems", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAddProblemsToSession_InvalidBody(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)
	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/problems", bytes.NewBufferString("invalid"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestSetResultHandler_Success(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	info := &models2.ResultIndo{
		SessionId: 1,
		Result: models2.Result{
			HardSkils: "good",
			SoftSkils: "great",
			ToHire:    "yes",
		},
	}
	sessionSvc.On("SetSessionResult", info).Return(nil)

	app := setupSessionApp(sessionSvc, userDB)

	body := toJSON(t, info)
	req := httptest.NewRequest(http.MethodPost, "/api/session/result", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSetResultHandler_InvalidBody(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)
	app := setupSessionApp(sessionSvc, userDB)

	req := httptest.NewRequest(http.MethodPost, "/api/session/result", bytes.NewBufferString("invalid"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestSetResultHandler_ServiceError(t *testing.T) {
	sessionSvc := new(MockSessionService)
	userDB := new(MockUserRepository)

	sessionSvc.On("SetSessionResult", mock.AnythingOfType("*models.ResultIndo")).Return(errors.New("db error"))

	app := setupSessionApp(sessionSvc, userDB)

	body := toJSON(t, &models2.ResultIndo{SessionId: 1})
	req := httptest.NewRequest(http.MethodPost, "/api/session/result", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
