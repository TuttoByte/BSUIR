package db

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type GormDB[T any] struct {
	db  *gorm.DB
	dns string
}

func NewGormDB[T any]() (*GormDB[T], error) {
	dsn := os.Getenv("DNS")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &GormDB[T]{
		db:  db,
		dns: dsn,
	}, nil
}

func (g *GormDB[T]) Create(ctx context.Context) error {
	var item T
	err := g.db.AutoMigrate(item)
	if err != nil {
		return err
	}
	return nil
}
func (g *GormDB[T]) FindByID(ctx context.Context, id uint64) (T, error) {
	var empty T
	user, err := gorm.G[T](g.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return empty, err
	}
	return user, nil
}
func (g *GormDB[T]) Update(ctx context.Context, item *T) error {
	g.db.Save(item)
	return nil
}

func (g *GormDB[T]) Delete(ctx context.Context, id uint64) error {
	_, err := gorm.G[T](g.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
