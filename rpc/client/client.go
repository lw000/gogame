package rpcclient

import "context"

type AuthItem struct {
	Username string
	Password string
}

// GetRequestMetadata gets the current request metadata
func (a *AuthItem) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"username": a.Username,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *AuthItem) RequireTransportSecurity() bool {
	return true
}
