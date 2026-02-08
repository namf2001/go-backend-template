package model

import "time"

// Account represents an account in the system
type Account struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validates account data
func (a *Account) Validate() error {
	if a.Username == "" {
		return ErrInvalidUsername
	}
	if a.Password == "" {
		return ErrInvalidPassword
	}
	return nil
}
