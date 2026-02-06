package response

import (
	"encoding/json"
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

	if errors, ok := err.(*apperrors.AppError); ok {
		appErr = errors
		errorInfo = ErrorInfo{
			Code:    appErr.Code,
			Message: appErr.Message,
		}

		// Map error types to HTTP status codes
		switch appErr.Code {
		case "NOT_FOUND":
			statusCode = http.StatusNotFound
		case "ALREADY_EXISTS":
			statusCode = http.StatusConflict
		case "INVALID_INPUT":
			statusCode = http.StatusBadRequest
		case "UNAUTHORIZED":
			statusCode = http.StatusUnauthorized
		case "FORBIDDEN":
			statusCode = http.StatusForbidden
		default:
			statusCode = http.StatusInternalServerError
		}
	} else {
		statusCode = http.StatusInternalServerError
		errorInfo = ErrorInfo{
			Code:    "INTERNAL_ERROR",
			Message: "An unexpected error occurred",
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
