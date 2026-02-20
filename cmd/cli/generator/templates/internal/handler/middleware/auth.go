package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
	"github.com/namf2001/go-backend-template/internal/pkg/jwt"
)

type contextKey string

const contextKeyUserID contextKey = "userID"

var (
	webErrMissingAuth  = &httpserv.Error{Status: http.StatusUnauthorized, Code: "missing_auth", Desc: "Missing authorization header"}
	webErrInvalidAuth  = &httpserv.Error{Status: http.StatusUnauthorized, Code: "invalid_auth", Desc: "Invalid authorization header format"}
	webErrInvalidToken = &httpserv.Error{Status: http.StatusUnauthorized, Code: "invalid_token", Desc: "Invalid or expired token"}
)

// RequireAuth middleware verifies JWT token
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpserv.RespondJSON(r.Context(), w, webErrMissingAuth)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			httpserv.RespondJSON(r.Context(), w, webErrInvalidAuth)
			return
		}

		tokenString := headerParts[1]
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			httpserv.RespondJSON(r.Context(), w, webErrInvalidToken)
			return
		}

		// Add UserID to context
		ctx := context.WithValue(r.Context(), contextKeyUserID, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
