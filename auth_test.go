package sdk

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestAuthenticate() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"access_token":"testtoken"}`)
	})
	token, res, err := suite.SDK.Auth.Authenticate("", "")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "testtoken", token)
	require.Equal(suite.T(), http.StatusOK, res.StatusCode)
}

func (suite *TestSuite) TestAuthenticate_withInvalidJSON() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[1,2,3]`)
	})
	_, res, err := suite.SDK.Auth.Authenticate("", "")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, res.StatusCode)
}
