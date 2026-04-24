package db

import (
	"Lab7/models"
	"context"
	"gorm.io/gorm"
)

type CandidatesDB struct {
	*GormDB[models.Candidate]
}

func NewCandidatesDB(ctx context.Context) (*CandidatesDB, error) {
	newGorm, err := NewGormDB[models.Candidate]()
	if err != nil {
		return nil, err
	}
	return &CandidatesDB{
		newGorm,
	}, nil
}

func (c *CandidatesDB) GetByEmail(ctx context.Context, email string) (models.Candidate, error) {
	user, err := gorm.G[models.Candidate](c.db).Where("email = ?", email).First(ctx)
	if err != nil {
		return models.Candidate{}, err
	}
	return user, nil
}

func (c *CandidatesDB) GetByName(ctx context.Context, name string) (models.Candidate, error) {
	user, err := gorm.G[models.Candidate](c.db).Where("name = ?", name).First(ctx)
	if err != nil {
		return models.Candidate{}, err
	}
	return user, nil
}
