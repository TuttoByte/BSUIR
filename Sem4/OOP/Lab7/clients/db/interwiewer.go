package db

import (
	"Lab7/models"
	"context"
	"gorm.io/gorm"
)

type InterwiewerDB struct {
	*GormDB[models.Interviewer]
}

func NewInterwiewerDB(ctx context.Context) (*InterwiewerDB, error) {
	newGorm, err := NewGormDB[models.Interviewer]()
	if err != nil {
		return nil, err
	}
	err = newGorm.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &InterwiewerDB{
		newGorm,
	}, nil
}

func (i *InterwiewerDB) GetPassHash(ctx context.Context, name string) (string, error) {
	var password string
	i.db.Model(&models.Interviewer{}).Where("name = ?", name).Pluck("hash_pass", &password)
	//fmt.Println(gorm.G[models.Interviewer](i.db).Select("hash_pass").Where("name = ?", name).First(ctx))
	return password, nil
}

func (i *InterwiewerDB) GetByName(ctx context.Context, name string) (models.Interviewer, error) {

	user, err := gorm.G[models.Interviewer](i.db).Where("name = ?", name).First(ctx)
	if err != nil {
		return models.Interviewer{}, err
	}
	return user, nil
}
