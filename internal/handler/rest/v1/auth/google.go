package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/namf2001/go-backend-template/internal/controller/auth"
	"github.com/namf2001/go-backend-template/internal/pkg/oauth"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
)

const (
	GoogleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

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
// @Success      200  {object} map[string]string
// @Router       /auth/google/login [get]
func (h *Handler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth.GoogleOauthConfig.AuthCodeURL(oauth.OauthStateString)
	response.Success(w, map[string]string{"url": url})
}

// GoogleCallback handles google callback
// @Summary      Google callback
// @Description  Handle Google OAuth callback and return token
// @Tags         auth
// @Produce      json
// @Param        state query string true "OAuth state"
// @Param        code  query string true "OAuth code"
// @Success      200  {object} map[string]string
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /auth/google/callback [get]
func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauth.OauthStateString {
		response.Error(w, fmt.Errorf("invalid oauth state"))
		return
	}

	code := r.FormValue("code")
	token, err := oauth.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		response.Error(w, fmt.Errorf("code exchange failed: %s", err.Error()))
		return
	}

	resp, err := http.Get(GoogleUserInfoURL + "?access_token=" + token.AccessToken)
	if err != nil {
		response.Error(w, fmt.Errorf("failed to get user info: %s", err.Error()))
		return
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		response.Error(w, fmt.Errorf("failed to read response body: %s", err.Error()))
		return
	}

	var userInfo GoogleUserInfo

	if err := json.Unmarshal(content, &userInfo); err != nil {
		response.Error(w, fmt.Errorf("failed to parse user info: %s", err.Error()))
		return
	}

	input := auth.OAuthInput{
		Provider:          "google",
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
		response.Error(w, err)
		return
	}

	// For now, return token as JSON. In real app might define how to return it (cookie vs json)
	response.Success(w, map[string]string{"token": authToken})
}
