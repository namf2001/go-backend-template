package model

import "time"

// Session represents a user session
type Session struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"userId" db:"userId"`
	Expires      time.Time `json:"expires" db:"expires"`
	SessionToken string    `json:"sessionToken" db:"sessionToken"`
}
