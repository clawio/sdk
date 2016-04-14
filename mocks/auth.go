package mocks

import (
	"github.com/clawio/codes"
	"github.com/clawio/entities"
	"github.com/stretchr/testify/mock"
)

// MockAuthService mocks an AuthService.
type MockAuthService struct {
	mock.Mock
}

// Authenticate mocks the Authenticate call.
func (m *MockAuthService) Authenticate(username, password string) (string, *codes.Response, error) {
	args := m.Called(username, password)
	return args.String(0), args.Get(1).(*codes.Response), args.Error(2)
}

// Verify mocks the Verify call.
func (m *MockAuthService) Verify(token string) (entities.User, *codes.Response, error) {
	args := m.Called(token)
	return args.Get(0).(entities.User), args.Get(1).(*codes.Response), args.Error(2)
}

// Invalidate mocks the Invalidate call.
func (m *MockAuthService) Invalidate(token string) (*codes.Response, error) {
	args := m.Called(token)
	return args.Get(0).(*codes.Response), args.Error(1)
}
