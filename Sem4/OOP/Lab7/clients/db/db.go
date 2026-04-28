package db

import (
	"Lab7/models"
	"context"
)

type Repository[T any] interface {
	Create(ctx context.Context) error
	FindByID(ctx context.Context, id uint64) (T, error)
	Update(ctx context.Context, item *T) error
	Delete(ctx context.Context, id uint64) error
}

type UserRepository interface {
	Repository[models.User]
	GetByName(ctx context.Context, name string) (models.User, error)
	GetPassHash(name string) (string, error)
	GetRole(name string) (string, error)
}

type SessionRepository interface {
	Repository[models.Session]
	GetAll() ([]models.Session, error)
	GetAllActiveSessions(ctx context.Context) ([]models.Session, error)
}

type ProblemsRepository interface {
	Repository[models.InterwieweProblem]
	GetAll() ([]models.InterwieweProblem, error)
	GetByName(ctx context.Context, name string) (models.InterwieweProblem, error)
	GetIdByName(ctx context.Context, name string) (models.InterwieweProblem, error)
}

type SlotRepository interface {
	Repository[models.AvailabilitySlot]
	GetAll() ([]models.AvailabilitySlot, error)
}
