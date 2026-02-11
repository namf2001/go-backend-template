package model

type Provider string

const (
	// Credentials is for manual registration
	ProviderCredentials Provider = "Credentials"
	// Google is for Google OAuth
	ProviderGoogle Provider = "google"
)

// String converts to string value
func (p Provider) String() string {
	return string(p)
}
	
// IsValid checks if the provider is valid
func (p Provider) IsValid() bool {
	return p == ProviderCredentials || p == ProviderGoogle
}
