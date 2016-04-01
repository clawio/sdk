package sdk

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	defaultBaseURL     = "/clawio/v1/"
	defaultAuthBaseURL = defaultBaseURL + "auth/"
	defaultDataBaseURL = defaultBaseURL + "data/"
)

type TestSuite struct {
	suite.Suite
	SDK    *SDK
	Router *mux.Router
	Server *httptest.Server
}

func Test(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	router := mux.NewRouter()
	server := httptest.NewServer(router)
	// Make a transport that reroutes all traffic to the example server
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	urls := &ServiceEndpoints{}
	urls.AuthServiceBaseURL = server.URL + defaultAuthBaseURL
	urls.DataServiceBaseURL = server.URL + defaultDataBaseURL
	httpClient := &http.Client{Transport: transport}
	sdk := New(urls, httpClient)
	suite.Router = router
	suite.Server = server
	suite.SDK = sdk
}

func (suite *TestSuite) TeardownTest() {
	defer suite.Server.Close()
}

func (suite *TestSuite) TestNew() {
	urls := &ServiceEndpoints{}
	urls.AuthServiceBaseURL = defaultAuthBaseURL
	urls.DataServiceBaseURL = defaultBaseURL
	sdk := New(urls, nil)
	require.NotNil(suite.T(), sdk)
}
