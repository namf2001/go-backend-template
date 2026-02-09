package model

// Account represents an OAuth account connection
type Account struct {
	ID                int64  `json:"id" db:"id"`
	UserID            int64  `json:"userId" db:"userId"`
	Type              string `json:"type" db:"type"`
	Provider          string `json:"provider" db:"provider"`
	ProviderAccountID string `json:"providerAccountId" db:"providerAccountId"`
	RefreshToken      string `json:"refresh_token" db:"refresh_token"`
	AccessToken       string `json:"access_token" db:"access_token"`
	ExpiresAt         int64  `json:"expires_at" db:"expires_at"`
	IDToken           string `json:"id_token" db:"id_token"`
	Scope             string `json:"scope" db:"scope"`
	SessionState      string `json:"session_state" db:"session_state"`
	TokenType         string `json:"token_type" db:"token_type"`
}
