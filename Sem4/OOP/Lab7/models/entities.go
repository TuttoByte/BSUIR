package models

type Interviewer struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string `json:"name"`
	HashPass string `json:"hash_pass"`
	Email    string `json:"email"`
}

type Candidate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	InterviewGrade
}

type InterviewGrade struct {
	SoftSkill string `json:"soft_skill"`
	HardSkill string `json:"hard_skill"`
	ToHire    bool   `json:"to_hire"`
}

type InterwieweProblem struct {
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
