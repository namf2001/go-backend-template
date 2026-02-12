package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	ctrlAuth "github.com/namf2001/go-backend-template/internal/controller/auth"
	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
	"github.com/namf2001/go-backend-template/internal/pkg/oauth"
)

const (
	GoogleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// GoogleLoginResponse represents the response for Google login
type GoogleLoginResponse struct {
	URL string `json:"url"`
}

// GoogleCallbackResponse represents the response for Google callback
type GoogleCallbackResponse struct {
	Token string `json:"token"`
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

// GoogleLogin handles google login
// @Summary      Google login
// @Description  Get Google login URL
// @Tags         auth
// @Produce      json
// @Success      200  {object} auth.GoogleLoginResponse
// @Router       /auth/google/login [get]
func (h *Handler) GoogleLogin() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		url := oauth.GoogleOauthConfig.AuthCodeURL(oauth.OauthStateString)
		httpserv.RespondJSON(r.Context(), w, GoogleLoginResponse{URL: url})
		return nil
	})
}

// GoogleCallback handles google callback
// @Summary      Google callback
// @Description  Handle Google OAuth callback and return token
// @Tags         auth
// @Produce      json
// @Param        state query string true "OAuth state"
// @Param        code  query string true "OAuth code"
// @Success      200  {object} auth.GoogleCallbackResponse
// @Failure      400  {object} httpserv.Error
// @Failure      500  {object} httpserv.Error
// @Router       /auth/google/callback [get]
func (h *Handler) GoogleCallback() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		state := r.FormValue("state")
		if state != oauth.OauthStateString {
			return webErrInvalidOAuthState
		}

		code := r.FormValue("code")
		token, err := oauth.GoogleOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			return webErrCodeExchangeFailed
		}

		resp, err := http.Get(GoogleUserInfoURL + "?access_token=" + token.AccessToken)
		if err != nil {
			return webErrGetUserInfoFailed
		}
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return webErrGetUserInfoFailed
		}

		var userInfo GoogleUserInfo
		if err := json.Unmarshal(content, &userInfo); err != nil {
			return webErrGetUserInfoFailed
		}

		input := ctrlAuth.OAuthInput{
			Provider:          model.ProviderGoogle,
			ProviderAccountID: userInfo.ID,
			Type:              "oauth",
			AccessToken:       token.AccessToken,
			RefreshToken:      token.RefreshToken,
			ExpiresAt:         token.Expiry.Unix(),
			TokenType:         token.TokenType,

			Name:          userInfo.Name,
			Email:         userInfo.Email,
			Image:         userInfo.Picture,
			EmailVerified: userInfo.VerifiedEmail,
		}

		authToken, err := h.ctrl.OAuthLogin(r.Context(), input)
		if err != nil {
			return err
		}

		httpserv.RespondJSON(r.Context(), w, GoogleCallbackResponse{Token: authToken})
		return nil
	})
}
