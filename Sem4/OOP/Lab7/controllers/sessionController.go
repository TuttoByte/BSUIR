package controllers

import (
	"Lab7/clients/db"
	"Lab7/models"
	"context"
	"errors"
)

type SessionController struct {
	session  *db.SessionDB
	problems *db.ProblemsDB
}

func NewSessionController(sessinons *db.SessionDB, problems *db.ProblemsDB) *SessionController {
	return &SessionController{
		session:  sessinons,
		problems: problems,
	}
}

func (s *SessionController) StartSession(id uint64) error {
	ctx := context.Background()
	session, err := s.session.FindByID(ctx, id)
	if err != nil {
		return err
	}
	session.Start()

	err = s.session.Update(ctx, &session)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionController) EndSession(id uint64) error {
	ctx := context.Background()
	session, err := s.session.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if session.IsActive() == false {
		return errors.New("session is not active")
	}
	session.End()
	err = s.session.Update(ctx, &session)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionController) DeleteSession(id uint64) error {
	ctx := context.Background()
	err := s.session.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionController) GetAllSessions() ([]models.Session, error) {
	sessions, err := s.session.GetAll()
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *SessionController) AddSession(candidate models.Candidate, interviwer models.Interviewer) error {
	ctx := context.Background()
	session := models.NewSession(candidate, interviwer)
	err := s.session.Update(ctx, session)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionController) GetSessionById(id uint64) (models.Session, error) {
	ctx := context.Background()
	session, err := s.session.FindByID(ctx, id)
	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func (s *SessionController) AddProblems(ids []uint64, sessionId uint64) error {
	if len(ids) == 0 {
		return errors.New("no index of problems is provided")
	}
	ctx := context.Background()
	problems := make([]models.InterwieweProblem, len(ids))
	for i, id := range ids {
		problem, err := s.problems.FindByID(ctx, id)
		if err != nil {
			return err
		}
		problems[i] = problem
	}

	session, err := s.GetSessionById(sessionId)
	if err != nil {
		return err
	}
	session.Problems = problems
	err = s.session.Update(ctx, &session)
	if err != nil {
		return err
	}
	return nil
}
