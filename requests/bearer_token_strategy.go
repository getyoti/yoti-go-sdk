package requests

import (
	"errors"
)

// BearerTokenStrategy implements AuthStrategy using a Bearer token for
// central auth. It sets the Authorization header with the provided token
// and does not add any query parameters.
type BearerTokenStrategy struct {
	token string
}

// NewBearerTokenStrategy creates a new BearerTokenStrategy with the given token.
// Returns an error if the token is empty.
func NewBearerTokenStrategy(token string) (*BearerTokenStrategy, error) {
	if token == "" {
		return nil, errors.New("authentication token must not be empty")
	}
	return &BearerTokenStrategy{token: token}, nil
}

// CreateAuthHeaders returns an Authorization header with a Bearer token.
func (s *BearerTokenStrategy) CreateAuthHeaders(httpMethod, endpoint string, body []byte) (map[string][]string, error) {
	return map[string][]string{
		"Authorization": {"Bearer " + s.token},
	}, nil
}

// CreateQueryParams returns an empty map as bearer auth does not require query parameters.
func (s *BearerTokenStrategy) CreateQueryParams() (map[string]string, error) {
	return map[string]string{}, nil
}
