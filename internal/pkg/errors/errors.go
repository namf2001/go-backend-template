package errors

import (
	"errors"
	"fmt"
)

// Common application errors
var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInternal      = errors.New("internal server error")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
)

// AppError represents an application error with additional context
type AppError struct {
	Code    string
	Message string
	Err     error
}

// Error shows the error message
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NotFound creates a not found error
func NotFound(message string) *AppError {
	return &AppError{
		Code:    "NOT_FOUND",
		Message: message,
		Err:     ErrNotFound,
	}
}

// AlreadyExists creates an already existed error
func AlreadyExists(message string) *AppError {
	return &AppError{
		Code:    "ALREADY_EXISTS",
		Message: message,
		Err:     ErrAlreadyExists,
	}
}

// InvalidInput creates an invalid input error
func InvalidInput(message string) *AppError {
	return &AppError{
		Code:    "INVALID_INPUT",
		Message: message,
		Err:     ErrInvalidInput,
	}
}

// Internal creates an internal error
func Internal(message string, err error) *AppError {
	return &AppError{
		Code:    "INTERNAL_ERROR",
		Message: message,
		Err:     err,
	}
}
