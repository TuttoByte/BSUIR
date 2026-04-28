package models

import "time"

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string `json:"name"`
	HashPass string `json:"hash_pass"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type InterviewGrade struct {
	SoftSkill string `json:"soft_skill"`
	HardSkill string `json:"hard_skill"`
	ToHire    bool   `json:"to_hire"`
}

type InterwieweProblem struct {
	ID    uint64 `gorm:"primaryKey"`
	Text  string `json:"text"`
	Grade int    `json:"grade"`
	Theme string `json:"theme"`
}

type SessionInfo struct {
	Problems          []string
	Username          string
	UserIdentificator string
	Candidate         string
}

type AvailabilitySlot struct {
	ID            uint64    `gorm:"primaryKey"`
	InterviewerID uint64    `json:"interviewer_id"`
	CandidateID   uint64    `json:"candidate_id"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	IsBooked      bool      `json:"is_booked"`
}

type AddProblemsInfo struct {
	SessionId uint64   `json:"session_id"`
	Problems  []uint64 `json:"problems"`
}

type ResultIndo struct {
	SessionId uint64 `json:"session_id"`
	Result
}
