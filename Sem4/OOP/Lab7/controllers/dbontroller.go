package controllers

import (
	"Lab7/clients/db"
	"context"
)

type Dbontroller struct {
	Candidates   *db.CandidatesDB
	Interwiewers *db.InterwiewerDB
	Problems     *db.ProblemsDB
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

	return &Dbontroller{
		Candidates:   candidate,
		Interwiewers: interwiewer,
		Problems:     problems,
	}, nil

}
