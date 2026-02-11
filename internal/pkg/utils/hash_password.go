package utils

import (
	"golang.org/x/crypto/bcrypt"
	pkgerrors "github.com/pkg/errors"
)

const defaultCost = 10

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		return "", pkgerrors.WithStack(err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword verifies a password against a hashed password
func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	return nil
}