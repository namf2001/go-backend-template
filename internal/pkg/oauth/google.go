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

func Init(cfg config.GoogleConfig) {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  cfg.RedirectURL,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       Scopes,
		Endpoint:     google.Endpoint,
	}
}
