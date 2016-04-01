package mock

import (
	"github.com/clawio/codes"
	"github.com/clawio/service-auth/server/spec"
	"github.com/stretchr/testify/mock"
)

type MockSDK struct {
	mock.Mock
}

func (m *MockSDK) Authenticate(username, password string) (string, *codes.Response, error) {
	args := m.Called(username, password)
	return args.String(0), args.Get(1).(*codes.Response), args.Error(2)
}
func (m *MockSDK) Verify(token string) (*spec.Identity, *codes.Response, error) {
	args := m.Called(token)
	return args.Get(0).(*spec.Identity), args.Get(1).(*codes.Response), args.Error(1)
}
