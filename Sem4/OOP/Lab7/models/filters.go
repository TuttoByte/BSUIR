package models

type Filtable interface {
	Filer(problem InterwieweProblem) bool
}

type FilterByCategory struct {
	Category string
}

func NewFilterByCategory(category string) *FilterByCategory {
	return &FilterByCategory{Category: category}
}

func (f *FilterByCategory) Filer(problem InterwieweProblem) bool {
	if problem.Theme == f.Category {
		return true
	}
	return false
}

type FilterByLevel struct {
	Level int
}

func NewFilterByLevel(level int) *FilterByLevel {
	return &FilterByLevel{Level: level}
}
func (f *FilterByLevel) Filer(problem InterwieweProblem) bool {
	if problem.Grade == f.Level {
		return true
	}
	return false
}
