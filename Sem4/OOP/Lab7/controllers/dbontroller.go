package controllers

import (
	"Lab7/clients/db"
	"context"
)

type Dbontroller struct {
	Users       db.UserRepository
	Problems    db.ProblemsRepository
	Sessions    db.SessionRepository
	Avalavility db.SlotRepository
}

func NewDbontroller(ctx context.Context) (*Dbontroller, error) {

	users, err := db.NewUsersDB(ctx)
	if err != nil {
		return nil, err
	}

	problems, err := db.NewProblemsDB(ctx)
	if err != nil {
		return nil, err
	}

	sessions, err := db.NewSessionDB()
	if err != nil {
		return nil, err
	}

	aval, err := db.NewAvalabilytyDB()
	if err != nil {
		return nil, err
	}

	err = users.Create(ctx)
	if err != nil {
		return nil, err
	}
	err = problems.Create(ctx)
	if err != nil {
		return nil, err
	}
	err = sessions.Create(ctx)
	if err != nil {
		return nil, err
	}
	err = aval.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &Dbontroller{
		Users:       users,
		Problems:    problems,
		Sessions:    sessions,
		Avalavility: aval,
	}, nil

}
