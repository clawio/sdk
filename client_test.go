package sdk

import (
	"bytes"
	"io/ioutil"
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

func (suite *TestSuite) TestNewRequest() {
	c := NewClient("", nil)
	_, err := c.NewRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
}

func (suite *TestSuite) TestNewRequest_withUserAgent() {
	c := NewClient("", nil)
	c.UserAgent = "custom"
	_, err := c.NewRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
}

func (suite *TestSuite) TestNewUploadRequest_withUserAgent() {
	c := NewClient("", nil)
	c.UserAgent = "custom"
	_, err := c.NewUploadRequest("/", nil)
	require.Nil(suite.T(), err)
}

func (suite *TestSuite) TestNewUploadRequest_withBadURL() {
	c := NewClient("", nil)
	c.UserAgent = "custom"
	_, err := c.NewUploadRequest(":", nil)
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestNewClient() {
	c := NewClient("", nil)
	c.UserAgent = "custom"
	_, err := c.NewRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
}

func (suite *TestSuite) TestCopyBody() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("1"))
	})
	buf := new(bytes.Buffer)
	c := NewClient("", nil)
	req, err := c.NewRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	_, err = c.Do(req, buf, false)
	require.Nil(suite.T(), err)
	data, err := ioutil.ReadAll(buf)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "1", string(data))
}
func (suite *TestSuite) TestCopyEmptyBody() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	c := NewClient("", nil)
	req, err := c.NewRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	type T struct{}
	t := &T{}
	_, err = c.Do(req, t, false)
	require.Nil(suite.T(), err)
}
