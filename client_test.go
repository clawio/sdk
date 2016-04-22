package sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestnewRequest() {
	c := newClient("", nil)
	_, err := c.newRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
}
func (suite *TestSuite) TestnewRequest_withBadURL() {
	c := newClient(":", nil)
	_, err := c.newRequest("GET", ":", nil)
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestnewRequest_withInvalidJSON() {
	c := newClient("", nil)
	type T struct {
		A map[int]interface{}
	}
	_, err := c.newRequest("GET", "/", &T{})
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestnewRequest_withBadMethod() {
	c := newClient("", nil)
	_, err := c.newRequest(":", "/", nil)
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestnewRequest_withuserAgent() {
	c := newClient("", nil)
	c.userAgent = "custom"
	r, err := c.newRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), r.Header.Get("user-agent"), "custom")
}

func (suite *TestSuite) TestnewUploadRequest_withuserAgent() {
	c := newClient("", nil)
	c.userAgent = "custom"
	_, err := c.newUploadRequest("/", nil)
	require.Nil(suite.T(), err)
}

func (suite *TestSuite) TestnewUploadRequest_withBadURL() {
	c := newClient("", nil)
	c.userAgent = "custom"
	_, err := c.newUploadRequest(":", nil)
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestnewClient() {
	c := newClient("", nil)
	c.userAgent = "custom"
	_, err := c.newRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
}

func (suite *TestSuite) Testdo_withClose() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("1"))
	})
	c := newClient("", nil)
	req, err := c.newRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	_, err = c.do(req, nil, true)
	require.Nil(suite.T(), err)
}
func (suite *TestSuite) Testdo_withWriter() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("1"))
	})
	buf := new(bytes.Buffer)
	c := newClient("", nil)
	req, err := c.newRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	_, err = c.do(req, buf, false)
	require.Nil(suite.T(), err)
	data, err := ioutil.ReadAll(buf)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "1", string(data))
}
func (suite *TestSuite) Testdo_withEmptyBody() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	c := newClient("", nil)
	req, err := c.newRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	type T struct{}
	t := &T{}
	_, err = c.do(req, t, false)
	require.Nil(suite.T(), err)
}

func (suite *TestSuite) Testdo_withError() {
	c := newClient("", nil)
	req, err := c.newRequest("GET", "", nil)
	require.Nil(suite.T(), err)
	_, err = c.do(req, nil, false)
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) Testdo_withResponseError() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("1"))
	})
	c := newClient("", nil)
	req, err := c.newRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	_, err = c.do(req, nil, true)
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestcheckResponse_withBadJSONBody() {
	resp := &http.Response{}
	resp.Body = ioutil.NopCloser(strings.NewReader("1"))
	resp.StatusCode = http.StatusBadRequest
	err := checkResponse(resp)
	require.NotNil(suite.T(), err)
}
