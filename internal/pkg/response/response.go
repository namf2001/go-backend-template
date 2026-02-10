package response

import (
	"encoding/json"
	"errors"
	"net/http"

	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error information in response
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// JSON sends a JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: statusCode < 400,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

// Error sends an error response
func Error(w http.ResponseWriter, err error) {
	var appErr *apperrors.AppError
	var statusCode int
	var errorInfo ErrorInfo

	if errors.As(err, &appErr) {
		errorInfo = ErrorInfo{
			Code:    appErr.Code,
			Message: appErr.Message,
		}
		statusCode = mapCodeToStatus(appErr.Code)
	} else {
		// Handle sentinel errors
		var code string
		if errors.Is(err, apperrors.ErrNotFound) {
			code = "NOT_FOUND"
			statusCode = http.StatusNotFound
		} else if errors.Is(err, apperrors.ErrAlreadyExists) {
			code = "ALREADY_EXISTS"
			statusCode = http.StatusConflict
		} else if errors.Is(err, apperrors.ErrInvalidInput) {
			code = "INVALID_INPUT"
			statusCode = http.StatusBadRequest
		} else if errors.Is(err, apperrors.ErrUnauthorized) {
			code = "UNAUTHORIZED"
			statusCode = http.StatusUnauthorized
		} else if errors.Is(err, apperrors.ErrForbidden) {
			code = "FORBIDDEN"
			statusCode = http.StatusForbidden
		} else if errors.Is(err, apperrors.ErrConflict) {
			code = "CONFLICT"
			statusCode = http.StatusConflict
		} else {
			code = "INTERNAL_ERROR"
			statusCode = http.StatusInternalServerError
		}

		msg := err.Error()
		if statusCode == http.StatusInternalServerError {
			msg = "An unexpected error occurred"
		}

		errorInfo = ErrorInfo{
			Code:    code,
			Message: msg,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: false,
		Error:   &errorInfo,
	}

	json.NewEncoder(w).Encode(response)
}

func mapCodeToStatus(code string) int {
	switch code {
	case "NOT_FOUND":
		return http.StatusNotFound
	case "ALREADY_EXISTS":
		return http.StatusConflict
	case "INVALID_INPUT":
		return http.StatusBadRequest
	case "UNAUTHORIZED":
		return http.StatusUnauthorized
	case "FORBIDDEN":
		return http.StatusForbidden
	case "CONFLICT":
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// Success sends a success response
func Success(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, data)
}

// Created sends a created response
func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, data)
}

// NoContent sends a no content response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
