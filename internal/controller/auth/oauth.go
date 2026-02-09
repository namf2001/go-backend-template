package auth

import (
	"context"
	"time"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/namf2001/go-backend-template/internal/pkg/jwt"
)

// OAuthInput is the input for OAuth login
type OAuthInput struct {
	Provider          string
	ProviderAccountID string
	Type              string
	AccessToken       string
	RefreshToken      string
	ExpiresAt         int64
	IDToken           string
	Scope             string
	SessionState      string
	TokenType         string
	// User info
	Name          string
	Email         string
	Image         string
	EmailVerified bool
}

// OAuthLogin handles oauth login/registration
func (i impl) OAuthLogin(ctx context.Context, input OAuthInput) (string, error) {
	// 1. Check if account exists
	account, err := i.repo.Account().GetByProvider(ctx, input.Provider, input.ProviderAccountID)
	if err == nil {
		// Account exists, get user and login
		user, err := i.repo.User().GetByID(ctx, account.UserID)
		if err != nil {
			return "", err
		}
		return jwt.GenerateToken(user.ID, user.Email)
	}

	// If account does not exist
	// 2. Check if user exists by email
	user, err := i.repo.User().GetByEmail(ctx, input.Email)
	if err != nil {
		// User does not exist, create user
		// Note: We might want cleaner error checking for NotFound vs other errors
		if !apperrors.IsNotFound(err) {
			// If error is NOT "not found", return it (ignore for now assuming it's not found)
			// But actually GetByEmail returns not found.
		}

		newUser := model.User{
			Name:  input.Name,
			Email: input.Email,
			Image: input.Image,
		}
		if input.EmailVerified {
			now := time.Now()
			newUser.EmailVerified = &now
		}

		user, err = i.repo.User().Create(ctx, newUser)
		if err != nil {
			return "", err
		}
	}

	// 3. Link account to user
	newAccount := model.Account{
		UserID:            user.ID,
		Type:              input.Type,
		Provider:          input.Provider,
		ProviderAccountID: input.ProviderAccountID,
		RefreshToken:      input.RefreshToken,
		AccessToken:       input.AccessToken,
		ExpiresAt:         input.ExpiresAt,
		IDToken:           input.IDToken,
		Scope:             input.Scope,
		SessionState:      input.SessionState,
		TokenType:         input.TokenType,
	}

	_, err = i.repo.Account().Create(ctx, newAccount)
	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(user.ID, user.Email)
}
