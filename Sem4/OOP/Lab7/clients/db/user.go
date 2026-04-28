package db

import (
	"Lab7/models"
	"context"
	"gorm.io/gorm"
)

type UserDB struct {
	*GormDB[models.User]
}

func NewUsersDB(ctx context.Context) (*UserDB, error) {
	newGorm, err := NewGormDB[models.User]()
	if err != nil {
		return nil, err
	}
	err = newGorm.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &UserDB{
		newGorm,
	}, nil
}

func (i *UserDB) GetPassHash(name string) (string, error) {
	var password string
	i.db.Model(&models.User{}).Where("name = ?", name).Pluck("hash_pass", &password)
	return password, nil
}

func (i *UserDB) GetByName(ctx context.Context, name string) (models.User, error) {

	user, err := gorm.G[models.User](i.db).Where("name = ?", name).First(ctx)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (i *UserDB) GetRole(name string) (string, error) {
	var role string
	i.db.Model(&models.User{}).Where("name = ?", name).Pluck("role", &role)
	return role, nil
}
