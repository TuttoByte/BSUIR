package contracts

import (
	"Lab7/models"
	"context"
)

type Repository[T any] interface {
	Create(ctx context.Context, item *T) error
	FindByID(ctx context.Context, id uint64) (*T, error)
	Update(ctx context.Context, item *T) error
	Delete(ctx context.Context, id uint64) error
}

type InterwiewerRepository interface {
	Repository[models.Interviewer]
	GetPassHash(context.Context) (string, error)
}

type ProblemsRepository interface {
	Repository[models.InterwieweProblem]
	GetByName(ctx context.Context, name string) (*models.InterwieweProblem, error)
	GetIdByName(ctx context.Context, name string) (uint64, error)
}

type CandidatesRepository interface {
	Repository[models.Candidate]
}
