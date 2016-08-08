package sdk

import (
	"github.com/clawio/clawiod/codes"
	"github.com/clawio/clawiod/services/authentication"
)

type (
	// AuthService is the interface that deals with the calls to an authentication authentication.
	AuthService interface {
		Token(username, password string) (string, *codes.Response, error)
	}

	authService struct {
		client  *client
		baseURL string
	}
)

// Token gets an access token after authenticating the user with username and password.
func (s *authService) Token(username, password string) (string, *codes.Response, error) {
	tokenRequest := &authentication.TokenRequest{
		Username: username,
		Password: password}
	req, err := s.client.newRequest("POST", "token", tokenRequest)
	if err != nil {
		return "", nil, err
	}
	tokenResponse := &authentication.TokenResponse{}
	resp, err := s.client.do(req, tokenResponse, true)
	if err != nil {
		return "", resp, err
	}
	return tokenResponse.AccessToken, resp, nil
}
