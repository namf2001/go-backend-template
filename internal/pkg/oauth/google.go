package oauth

import (
	"github.com/namf2001/go-backend-template/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	OauthStateString  = "random-string"
	Scopes            = []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}
)

func Init() {
	cfg := config.GetConfig()
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  cfg.GetString("GOOGLE_REDIRECT_URL"),
		ClientID:     cfg.GetString("GOOGLE_CLIENT_ID"),
		ClientSecret: cfg.GetString("GOOGLE_CLIENT_SECRET"),
		Scopes:       Scopes,
		Endpoint:     google.Endpoint,
	}
}
