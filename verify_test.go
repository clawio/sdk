package sdk

/*
import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestVerify() {
	suite.Router.HandleFunc(defaultAuthBaseURL+"/{token}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"identity": {"username": "test", "email": "test@test.com", "display_name":"Test"}}`)
	})
	identity, _, err := suite.SDK.Auth.Verify("test")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "test", identity.Username)
	require.Equal(suite.T(), "test@test.com", identity.Email)
	require.Equal(suite.T(), "Test", identity.DisplayName)
}

func (suite *TestSuite) TestVerifyInvalidJSONBody() {
	suite.Router.HandleFunc("/clawio/auth/v1/verify/{token}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, ``)
	})
	_, _, err := suite.SDK.Auth.Verify("testtoken")
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestVerifyAPIError() {
	suite.Router.HandleFunc("/clawio/auth/v1/verify/{token}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "some api error", "code": 99}`)
	})
	_, _, err := suite.SDK.Auth.Verify("testtoken")
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestVerifyInvalidAPIError() {
	suite.Router.HandleFunc("/clawio/auth/v1/verify/{token}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{{}}`)
	})
	_, _, err := suite.SDK.Auth.Verify("testtoken")
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestVerifyNetworkError() {
	suite.Server.Close()
	_, _, err := suite.SDK.Auth.Verify("")
	require.NotNil(suite.T(), err)
}
*/
