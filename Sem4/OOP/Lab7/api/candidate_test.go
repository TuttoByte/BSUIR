package api

import (
	"Lab7/models"
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ===== МОКИ =====

type MockCandidateService struct {
	mock.Mock
}

func (m *MockCandidateService) GetSlostsByInterwiver(id uint64) ([]models.AvailabilitySlot, error) {
	args := m.Called(id)
	return args.Get(0).([]models.AvailabilitySlot), args.Error(1)
}

func (m *MockCandidateService) BookSlot(userId, slotId uint64) error {
	args := m.Called(userId, slotId)
	return args.Error(0)
}

func (m *MockCandidateService) UnbookSlot(userId, slotId uint64) error {
	args := m.Called(userId, slotId)
	return args.Error(0)
}

// ===== ФИКСТУРЫ =====

func mockSlots() []models.AvailabilitySlot {
	return []models.AvailabilitySlot{
		{
			ID:            1,
			InterviewerID: 1,
			CandidateID:   0,
			StartTime:     time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			EndTime:       time.Date(2025, 1, 1, 11, 0, 0, 0, time.UTC),
			IsBooked:      false,
		},
		{
			ID:            2,
			InterviewerID: 1,
			CandidateID:   0,
			StartTime:     time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			EndTime:       time.Date(2025, 1, 1, 13, 0, 0, 0, time.UTC),
			IsBooked:      false,
		},
	}
}

// ===== ХЕЛПЕР =====

func setupApp(candidates *MockCandidateService) *fiber.App {
	app := fiber.New()
	a := &App{candidates: candidates}

	app.Get("/api/candidates/:id/slots", a.GetAvalableSlotsHandler)
	app.Post("/api/candidates/:CandidateId/slots/:SlotId/book", a.CandidateBookSlotHandler)
	app.Post("/api/candidates/:CandidateId/slots/:SlotId/unbook", a.CandidateUnbookSlotHandler)

	return app
}

// ===== GetAvalableSlotsHandler =====

func TestGetAvalableSlotsHandler_Success(t *testing.T) {
	mockSvc := new(MockCandidateService)
	slots := mockSlots()
	mockSvc.On("GetSlostsByInterwiver", uint64(1)).Return(slots, nil)

	app := setupApp(mockSvc)

	req := httptest.NewRequest("GET", "/api/candidates/1/slots", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Проверяем тело ответа
	body, _ := io.ReadAll(resp.Body)
	var result []models.AvailabilitySlot
	json.Unmarshal(body, &result)
	assert.Len(t, result, 2)
	assert.Equal(t, uint64(1), result[0].InterviewerID)
	assert.False(t, result[0].IsBooked)

	mockSvc.AssertExpectations(t)
}

func TestGetAvalableSlotsHandler_InvalidID(t *testing.T) {
	mockSvc := new(MockCandidateService)
	app := setupApp(mockSvc)

	req := httptest.NewRequest("GET", "/api/candidates/abc/slots", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockSvc.AssertNotCalled(t, "GetSlostsByInterwiver")
}

func TestGetAvalableSlotsHandler_ServiceError(t *testing.T) {
	mockSvc := new(MockCandidateService)
	mockSvc.On("GetSlostsByInterwiver", uint64(1)).Return([]models.AvailabilitySlot{}, errors.New("db error"))

	app := setupApp(mockSvc)

	req := httptest.NewRequest("GET", "/api/candidates/1/slots", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var result map[string]string
	json.Unmarshal(body, &result)
	assert.Equal(t, "db error", result["error"])
}

// ===== CandidateBookSlotHandler =====

func TestCandidateBookSlotHandler_Success(t *testing.T) {
	mockSvc := new(MockCandidateService)
	mockSvc.On("BookSlot", uint64(1), uint64(2)).Return(nil)

	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/1/slots/2/book", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockSvc.AssertExpectations(t)
}

func TestCandidateBookSlotHandler_InvalidCandidateID(t *testing.T) {
	mockSvc := new(MockCandidateService)
	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/abc/slots/2/book", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockSvc.AssertNotCalled(t, "BookSlot")
}

func TestCandidateBookSlotHandler_InvalidSlotID(t *testing.T) {
	mockSvc := new(MockCandidateService)
	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/1/slots/abc/book", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockSvc.AssertNotCalled(t, "BookSlot")
}

func TestCandidateBookSlotHandler_ServiceError(t *testing.T) {
	mockSvc := new(MockCandidateService)
	mockSvc.On("BookSlot", uint64(1), uint64(2)).Return(errors.New("slot unavailable"))

	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/1/slots/2/book", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var result map[string]string
	json.Unmarshal(body, &result)
	assert.Equal(t, "slot unavailable", result["error"])
}

// ===== CandidateUnbookSlotHandler =====

func TestCandidateUnbookSlotHandler_Success(t *testing.T) {
	mockSvc := new(MockCandidateService)
	mockSvc.On("UnbookSlot", uint64(1), uint64(2)).Return(nil)

	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/1/slots/2/unbook", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockSvc.AssertExpectations(t)
}

func TestCandidateUnbookSlotHandler_InvalidCandidateID(t *testing.T) {
	mockSvc := new(MockCandidateService)
	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/abc/slots/2/unbook", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockSvc.AssertNotCalled(t, "UnbookSlot")
}

func TestCandidateUnbookSlotHandler_InvalidSlotID(t *testing.T) {
	mockSvc := new(MockCandidateService)
	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/1/slots/abc/unbook", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockSvc.AssertNotCalled(t, "UnbookSlot")
}

func TestCandidateUnbookSlotHandler_ServiceError(t *testing.T) {
	mockSvc := new(MockCandidateService)
	mockSvc.On("UnbookSlot", uint64(1), uint64(2)).Return(errors.New("slot not booked"))

	app := setupApp(mockSvc)

	req := httptest.NewRequest("POST", "/api/candidates/1/slots/2/unbook", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var result map[string]string
	json.Unmarshal(body, &result)
	assert.Equal(t, "slot not booked", result["error"])
}
