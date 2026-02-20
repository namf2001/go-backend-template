package auth

import (
	"net/http"

	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
)

var (
	webErrValidationFailed   = &httpserv.Error{Status: http.StatusBadRequest, Code: "validation_failed", Desc: "Validation failed"}
	webErrInvalidCredentials = &httpserv.Error{Status: http.StatusUnauthorized, Code: "invalid_credentials", Desc: "Invalid email or password"}
	webErrInvalidOAuthState  = &httpserv.Error{Status: http.StatusBadRequest, Code: "invalid_oauth_state", Desc: "Invalid OAuth state"}
	webErrCodeExchangeFailed = &httpserv.Error{Status: http.StatusBadRequest, Code: "code_exchange_failed", Desc: "OAuth code exchange failed"}
	webErrGetUserInfoFailed  = &httpserv.Error{Status: http.StatusInternalServerError, Code: "get_user_info_failed", Desc: "Failed to get user info from provider"}
)
