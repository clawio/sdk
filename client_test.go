package sdk

import (
	"net/http"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestClientNewRequestBadURL() {
	c := NewClient(":", nil)
	_, err := c.NewRequest("GET", ":", nil)
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestNewRequestInvalidJSON() {
	c := NewClient("", nil)
	type T struct {
		A map[int]interface{}
	}
	_, err := c.NewRequest("GET", "/", &T{})
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestNewClient() {
	c := NewClient("", nil)
	c.UserAgent = "custom"
	_, err := c.NewRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
}
func (suite *TestSuite) TestEmptyBody() {
	suite.Router.HandleFunc("doesnotexists", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
	})
	_, _, err := suite.SDK.Auth.Authenticate("", "")
	require.NotNil(suite.T(), err)
}
