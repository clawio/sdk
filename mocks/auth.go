package mocks

import (
	"github.com/clawio/codes"
	"github.com/clawio/service-auth/server/spec"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Authenticate(username, password string) (string, *codes.Response, error) {
	args := m.Called(username, password)
	return args.String(0), args.Get(1).(*codes.Response), args.Error(2)
}
func (m *MockAuthService) Verify(token string) (*spec.Identity, *codes.Response, error) {
	args := m.Called(token)
	return args.Get(0).(*spec.Identity), args.Get(1).(*codes.Response), args.Error(2)
}
