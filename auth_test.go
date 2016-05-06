package sdk

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestAuthenticate() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"authenticate", func(w http.ResponseWriter, r *http.Request) {
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
	suite.Router.HandleFunc(defaultAuthBaseURL+"authenticate", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[1,2,3]`)
	})
	_, res, err := suite.SDK.Auth.Authenticate("", "")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, res.StatusCode)
}

func (suite *TestSuite) TestVerify() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"verify/testtoken", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"username": "test", "email":"test@test.com", "display_name": "Test"}`)
	})
	user, resp, err := suite.SDK.Auth.Verify("testtoken")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	require.Equal(suite.T(), "test", user.GetUsername())
	require.Equal(suite.T(), "test@test.com", user.GetEmail())
	require.Equal(suite.T(), "Test", user.GetDisplayName())
}
func (suite *TestSuite) TestVerify_withInvalidJSON() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"verify/testtoken", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[1,2,3]`)
	})
	_, resp, err := suite.SDK.Auth.Verify("testtoken")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}
func (suite *TestSuite) TestInvalidate() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"invalidate/testtoken", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
	resp, err := suite.SDK.Auth.Invalidate("testtoken")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusNoContent, resp.StatusCode)
}
func (suite *TestSuite) TestVerify_withBadInput() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"invalidate/testtoken", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"code": 99, "message": "test"}`)
	})
	resp, err := suite.SDK.Auth.Invalidate("testtoken")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}
