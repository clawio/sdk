package sdk

import (
	"path"

	"github.com/clawio/authentication/service"
	"github.com/clawio/codes"
	"github.com/clawio/entities"
)

type (
	// AuthService is the interface that deals with the calls to an authentication service.
	AuthService interface {
		Authenticate(username, password string) (string, *codes.Response, error)
		Verify(token string) (entities.User, *codes.Response, error)
		Invalidate(token string) (*codes.Response, error)
	}

	authService struct {
		client  *client
		baseURL string
	}

	user struct {
		Username    string `json:"username"`
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
	}
)

func (u *user) GetUsername() string    { return u.Username }
func (u *user) GetEmail() string       { return u.Email }
func (u *user) GetDisplayName() string { return u.DisplayName }

// Authenticate authenticates a user using a username and a password.
func (s *authService) Authenticate(username, password string) (string, *codes.Response, error) {
	authNRequest := &service.AuthenticateRequest{
		Username: username,
		Password: password}
	req, err := s.client.newRequest("POST", "authenticate", authNRequest)
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

// Verify verifies if an issued authn token is valid. If it is valid returns
// the user obtained from it.
func (s *authService) Verify(token string) (entities.User, *codes.Response, error) {
	token = path.Join("/", token)
	req, err := s.client.newRequest("GET", "verify"+token, nil)
	if err != nil {
		return nil, nil, err
	}
	u := &user{}
	resp, err := s.client.do(req, u, true)
	if err != nil {
		return nil, resp, err
	}
	return u, resp, nil
}

// Invalidate invalidates a token.
func (s *authService) Invalidate(token string) (*codes.Response, error) {
	req, err := s.client.newRequest("DELETE", "invalidate/"+token, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.do(req, nil, true)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
