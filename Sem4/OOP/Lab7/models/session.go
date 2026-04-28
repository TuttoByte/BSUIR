package models

import (
	"time"
)

const (
	StandartProblemsAmount = 2
	ExtendeProblemsAmount  = 4
)

type Session struct {
	ID            uint64 `gorm:"primaryKey"`
	CandidateId   uint64
	InterviewerId uint64
	Candidate     Candidate `json:"candidate" gorm:"foreignKey:CandidateId"`

	// Исправлено: Interwiewer -> Interviewer
	Interviewer Interviewer `json:"interviewer" gorm:"foreignKey:InterviewerId"`

	// Поле называется Problems, значит и прелоадить нужно "Problems"
	Problems []InterwieweProblem `json:"problems" gorm:"many2many:session_problems"`

	TimeInfo   TimeInfo `json:"time_info" gorm:"serializer:json"`
	IsStarted  bool
	isExtended bool

	Result Result `gorm:"serializer:json"`
}

type TimeInfo struct {
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	Duration  time.Duration `json:"duration"`
}

type Result struct {
	HardSkils string `json:"hard_skils"`
	SoftSkils string `json:"soft_skils"`
	ToHire    string `json:"to_hire"`
}

func NewSession(candidate Candidate, interwiewer Interviewer) *Session {
	return &Session{
		Candidate:   candidate,
		Interviewer: interwiewer,
		Problems:    make([]InterwieweProblem, 0),
		IsStarted:   false,
		isExtended:  false,
	}
}

func (s *Session) Start() {
	if s.IsStarted {
		return
	}
	s.IsStarted = true
	s.TimeInfo.StartTime = time.Now()
	return
}

func (s *Session) End() {
	if !s.IsStarted {
		return
	}
	s.TimeInfo.EndTime = time.Now()
	s.TimeInfo.Duration = time.Since(s.TimeInfo.StartTime)
}

func (s *Session) IsActive() bool {
	return s.IsStarted
}

func (s *Session) IsExtended() bool {
	return s.isExtended
}

func (s *Session) AddProblems(additional []InterwieweProblem) {
	if len(additional) > StandartProblemsAmount {
		s.isExtended = true
	}
	s.Problems = append(s.Problems, additional...)
}

func (s *Session) SetResult(result Result) {
	s.Result = result
}
func (s *Session) GetResult() Result {
	return s.Result
}
