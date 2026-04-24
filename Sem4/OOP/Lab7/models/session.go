package models

import (
	"time"
)

const (
	StandartProblemsAmount = 2
	ExtendeProblemsAmount  = 4
)

type Session struct {
	candidate   Candidate
	interwiewer Interviewer

	problems []InterwieweProblem

	timeInfo   TimeInfo
	isStarted  bool
	isExtended bool
}

type TimeInfo struct {
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

func NewSession(candidate Candidate, interwiewer Interviewer) *Session {
	return &Session{
		candidate:   candidate,
		interwiewer: interwiewer,
		problems:    make([]InterwieweProblem, 0),
		isStarted:   false,
		isExtended:  false,
	}
}

func (s *Session) Start() {
	if s.isStarted {
		return
	}
	s.isStarted = true
	s.timeInfo.StartTime = time.Now()
	return
}

func (s *Session) End() {
	if !s.isStarted {
		return
	}
	s.timeInfo.EndTime = time.Now()
	s.timeInfo.Duration = time.Since(s.timeInfo.StartTime)
}

func (s *Session) IsStarted() bool {
	return s.isStarted
}

func (s *Session) IsExtended() bool {
	return s.isExtended
}

func (s *Session) AddProblems(additional []InterwieweProblem) {
	if len(additional) > StandartProblemsAmount {
		s.isExtended = true
	}
	s.problems = append(s.problems, additional...)
}
