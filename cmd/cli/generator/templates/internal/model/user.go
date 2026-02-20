package model

import "time"

// User represents a user in the system
type User struct {
	ID            int64      `json:"id" db:"id"`
	Email         string     `json:"email" db:"email"`
	Name          string     `json:"name" db:"name"`
	EmailVerified *time.Time `json:"emailVerified" db:"emailVerified"`
	Image         string     `json:"image" db:"image"`
	Password      string     `json:"-" db:"password"` // Stored in users now
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

// Validate validates user data
func (u *User) Validate() error {
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if u.Name == "" {
		return ErrInvalidName
	}
	return nil
}
