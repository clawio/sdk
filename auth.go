package sdk

import (
	"github.com/clawio/authentication/service"
	"github.com/clawio/codes"
)

type (
	// AuthService is the interface that deals with the calls to an authentication service.
	AuthService interface {
		Authenticate(username, password string) (string, *codes.Response, error)
	}

	authService struct {
		client  *client
		baseURL string
	}
)

// Authenticate authenticates a user using a username and a password.
func (s *authService) Authenticate(username, password string) (string, *codes.Response, error) {
	authNRequest := &service.AuthenticateRequest{
		Username: username,
		Password: password}
	req, err := s.client.newRequest("POST", "token", authNRequest)
	if err != nil {
		return "", nil, err
	}
	authNResponse := &service.AuthenticateResponse{}
	resp, err := s.client.do(req, authNResponse, true)
	if err != nil {
		return "", resp, err
	}
	return authNResponse.AccessToken, resp, nil
}
