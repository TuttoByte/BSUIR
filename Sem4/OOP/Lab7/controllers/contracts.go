package controllers

import "Lab7/models"

type CandidateControllerable interface {
	GetSlostsByInterwiver(id uint64) ([]models.AvailabilitySlot, error)
	BookSlot(userId, slotId uint64) error
	UnbookSlot(userId, slotId uint64) error
}

type SessionControllerable interface {
	StartSession(id uint64) error
	EndSession(id uint64) error
	DeleteSession(id uint64) error
	GetAllSessions() ([]models.Session, error)
	AddSession(candidate models.User, interviwer models.User) error
	GetSessionById(id uint64) (models.Session, error)
	AddProblems(ids []uint64, sessionId uint64) error
	SetSessionResult(info *models.ResultIndo) error
}

type ProblemControllerable interface {
	AddNewProblem(problem *models.InterwieweProblem) error
	GetAllProblems() ([]models.InterwieweProblem, error)
	GetProblemById(id uint64) (models.InterwieweProblem, error)
	GetFileredProblems(filter models.Filtable) ([]models.InterwieweProblem, error)
}

type SlotControllerable interface {
	AddSlot(slot *models.AvailabilitySlot) error
	DeleteSlot(id uint64) error
	BookSlot(id uint64) error
	UnbookSlot(id uint64) error
	GetAllSlots() ([]models.AvailabilitySlot, error)
}
