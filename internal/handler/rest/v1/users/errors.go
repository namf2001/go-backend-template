package users

import (
	"errors"
	"net/http"

	ctrlUsers "github.com/namf2001/go-backend-template/internal/controller/users"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
	repoUsers "github.com/namf2001/go-backend-template/internal/repository/users"
)

var (
	webErrInvalidID = &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_id", Desc: "Invalid user ID"}

	webErrValidationFailed = &httpserv.Error{Status: http.StatusBadRequest, Code: "validation_failed", Desc: "Validation failed"}
	webErrUserExists       = &httpserv.Error{Status: http.StatusConflict, Code: "user_exists", Desc: "User with this email already exists"}
	webErrUserNotFound     = &httpserv.Error{Status: http.StatusNotFound, Code: "user_not_found", Desc: "User not found"}
)

func convertError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, ctrlUsers.ErrUserExited):
		return webErrUserExists
	case errors.Is(err, repoUsers.ErrNotFound):
		return webErrUserNotFound
	default:
		return err
	}
}
