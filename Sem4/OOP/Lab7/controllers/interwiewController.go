package controllers

import (
	"Lab7/models"
	"context"
)

type InterwieController struct {
	sessions []*models.Session
	db       *Dbontroller
}

func NewInterwieController(ctx context.Context) (*InterwieController, error) {

	db, err := NewDbontroller(ctx)
	if err != nil {
		return nil, err
	}

	return &InterwieController{
		sessions: make([]*models.Session, 0),
		db:       db,
	}, nil
}

func (i *InterwieController) GetSessions() []*models.Session {
	return i.sessions
}

func (i *InterwieController) GetActiveSessions() []*models.Session {
	activeSessions := make([]*models.Session, 0)
	for _, session := range i.sessions {
		if session.IsStarted() {
		}
		activeSessions = append(activeSessions, session)
	}
	return activeSessions
}

func (i *InterwieController) AddSession(info models.SessionInfo, ctx context.Context) error {
	candidate, err := i.db.Candidates.GetByName(ctx, info.Candidate)
	if err != nil {
		return err

	}
	interwiewer, err := i.db.Interwiewers.GetByName(ctx, info.Username)

	session := models.NewSession(candidate, interwiewer)
	i.sessions = append(i.sessions, session)
	return nil
}
