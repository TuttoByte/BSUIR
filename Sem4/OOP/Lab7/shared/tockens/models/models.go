package models

import (
	"github.com/vk-rv/pvx"
	"time"
)

type AdditionalClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Footer struct {
	MetaData string `json:"meta_data"`
}

type ServiceClaims struct {
	pvx.RegisteredClaims
	AdditionalClaims
	Footer //Not crypto
}

type TockenData struct {
	Subject  string
	Duration time.Duration
	AdditionalClaims
	Footer
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type RegisterInfo struct {
	Credentials
	Email string `json:"email"`
}

type AnyUserInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Maybe fix
type SessionRegisterInfo struct {
	CandidateId  uint64 `json:"candidate_id"`
	InterwiwerId uint64 `json:"interwiwer_id"`
}

type IdSetter struct {
	Id uint64 `json:"id"`
}
