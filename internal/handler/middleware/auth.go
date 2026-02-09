package middleware

import (
	"context"
	"net/http"
	"strings"

	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/namf2001/go-backend-template/internal/pkg/jwt"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
	"github.com/pkg/errors"
)

type contextKey string

const contextKeyUserID contextKey = "userID"

// RequireAuth middleware verifies JWT token
// RequireAuth middleware verifies JWT token
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(w, errors.Wrap(apperrors.ErrUnauthorized, "missing authorization header"))
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.Error(w, errors.Wrap(apperrors.ErrUnauthorized, "invalid authorization header format"))
			return
		}

		tokenString := headerParts[1]
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.Error(w, errors.Wrap(apperrors.ErrUnauthorized, "invalid or expired token"))
			return
		}

		// Add UserID to context
		ctx := context.WithValue(r.Context(), contextKeyUserID, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
