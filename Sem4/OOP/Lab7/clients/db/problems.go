package db

import (
	"Lab7/models"
	"context"
	"gorm.io/gorm"
)

type ProblemsDB struct {
	*GormDB[models.InterwieweProblem]
}

func NewProblemsDB(ctx context.Context) (*ProblemsDB, error) {

	newGorm, err := NewGormDB[models.InterwieweProblem]()
	if err != nil {
		return nil, err
	}
	err = newGorm.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &ProblemsDB{
		newGorm,
	}, nil
}

func (p *ProblemsDB) GetByName(ctx context.Context, name string) (models.InterwieweProblem, error) {

	user, err := gorm.G[models.InterwieweProblem](p.db).Where("name = ?", name).First(ctx)
	if err != nil {
		return models.InterwieweProblem{}, err
	}
	return user, nil
}
func (p *ProblemsDB) GetIdByName(ctx context.Context, name string) (models.InterwieweProblem, error) {
	user, err := gorm.G[models.InterwieweProblem](p.db).Where("name = ?", name).First(ctx)
	if err != nil {
		return models.InterwieweProblem{}, err
	}
	return user, nil
}
