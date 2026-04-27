package db

import (
	"Lab7/models"
	"context"
	"gorm.io/gorm"
)

type SessionDB struct {
	*GormDB[models.Session]
}

func NewSessionDB() (*SessionDB, error) {

	newGorm, err := NewGormDB[models.Session]()
	if err != nil {
		return nil, err
	}
	return &SessionDB{
		newGorm,
	}, nil
}

func (s *SessionDB) GetAllActiveSessions(ctx context.Context) ([]models.Session, error) {
	sessions, err := gorm.G[models.Session](s.db).Where("is_started = ?", true).Find(ctx)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *SessionDB) GetAll() ([]models.Session, error) {
	var sessions []models.Session
	s.db.Preload("Candidate").Preload("Interviewer").Find(&sessions)
	return sessions, nil
}
