package users

import "errors"

var (
	ErrUserExited = errors.New("user with this email already exists")
)