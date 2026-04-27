package controllers

import (
	"Lab7/clients/db"
	"Lab7/models"
	"context"
)

type ProblemsController struct {
	Problems *db.ProblemsDB
}

func NewProblemsController(problems *db.ProblemsDB) *ProblemsController {
	return &ProblemsController{
		Problems: problems,
	}
}

func (p *ProblemsController) AddNewProblem(problem *models.InterwieweProblem) error {
	ctx := context.Background()
	err := p.Problems.Update(ctx, problem)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemsController) DeleteProblem(id uint64) error {
	ctx := context.Background()
	err := p.Problems.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemsController) GetAllProblems() ([]models.InterwieweProblem, error) {
	return p.Problems.GetAll()
}

func (p *ProblemsController) GetProblemById(id uint64) (models.InterwieweProblem, error) {
	ctx := context.Background()
	problem, err := p.Problems.FindByID(ctx, id)
	if err != nil {
		return models.InterwieweProblem{}, err
	}
	return problem, nil
}

func (p *ProblemsController) GetFileredProblems(filter models.Filtable) ([]models.InterwieweProblem, error) {
	problems, err := p.GetAllProblems()
	if err != nil {
		return nil, err
	}

	filteredProblems := []models.InterwieweProblem{}
	for _, problem := range problems {
		if filter.Filer(problem) {
			filteredProblems = append(filteredProblems, problem)
		}
	}
	return filteredProblems, nil
}
