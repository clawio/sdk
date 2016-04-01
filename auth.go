package sdk

import (
	"github.com/clawio/codes"
	"github.com/clawio/service-auth/server/spec"
)

type AuthService interface {
	Authenticate(username, password string) (string, *codes.Response, error)
	Verify(token string) (*spec.Identity, *codes.Response, error)
}

type authService struct {
	client  *Client
	baseURL string
}

// Authenticate authenticates a user using a username and a password.
func (s *authService) Authenticate(username, password string) (string, *codes.Response, error) {
	authNRequest := &spec.AuthNRequest{
		Username: username,
		Password: password}
	req, err := s.client.NewRequest("POST", "authenticate", authNRequest)
	if err != nil {
		return "", nil, err
	}
	authNResponse := &spec.AuthNResponse{}
	resp, err := s.client.Do(req, authNResponse, true)
	if err != nil {
		return "", resp, err
	}
	return authNResponse.Token, resp, nil
}

// Verify verifies if an issued authn token is valid. If it is valid returns
// the identity obtained from it.
func (s *authService) Verify(token string) (*spec.Identity, *codes.Response, error) {
	req, err := s.client.NewRequest("GET", "verify/"+token, nil)
	if err != nil {
		return nil, nil, err
	}
	verifyResponse := &spec.VerifyResponse{}
	resp, err := s.client.Do(req, verifyResponse, true)
	if err != nil {
		return nil, resp, err
	}
	return verifyResponse.Identity, resp, nil
}
