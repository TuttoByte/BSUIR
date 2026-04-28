package controllers

import (
	"Lab7/clients/db"
	"context"
)

type Dbontroller struct {
	Candidates   *db.CandidatesDB
	Interwiewers *db.InterwiewerDB
	Problems     *db.ProblemsDB
	Sessions     *db.SessionDB
	Avalavility  *db.AvalabilytyDB
}

func NewDbontroller(ctx context.Context) (*Dbontroller, error) {
	candidate, err := db.NewCandidatesDB(ctx)
	if err != nil {
		return nil, err
	}

	interwiewer, err := db.NewInterwiewerDB(ctx)
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

	err = candidate.Create(ctx)
	if err != nil {
		return nil, err
	}
	err = interwiewer.Create(ctx)
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
		Candidates:   candidate,
		Interwiewers: interwiewer,
		Problems:     problems,
		Sessions:     sessions,
		Avalavility:  aval,
	}, nil

}
