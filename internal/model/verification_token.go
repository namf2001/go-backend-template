package model

import "time"

// VerificationToken represents a verification token
type VerificationToken struct {
	Identifier string    `json:"identifier" db:"identifier"`
	Expires    time.Time `json:"expires" db:"expires"`
	Token      string    `json:"token" db:"token"`
}
