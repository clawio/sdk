package sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestNewRequest() {
	c := NewClient("", nil)
	_, err := c.NewRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
}
func (suite *TestSuite) TestNewRequest_withBadURL() {
	c := NewClient(":", nil)
	_, err := c.NewRequest("GET", ":", nil)
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestNewRequest_withInvalidJSON() {
	c := NewClient("", nil)
	type T struct {
		A map[int]interface{}
	}
	_, err := c.NewRequest("GET", "/", &T{})
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestNewRequest_withBadMethod() {
	c := NewClient("", nil)
	_, err := c.NewRequest(":", "/", nil)
	require.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestNewRequest_withUserAgent() {
	c := NewClient("", nil)
	c.UserAgent = "custom"
	r, err := c.NewRequest("GET", "/", nil)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), r.Header.Get("user-agent"), "custom")
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

func (suite *TestSuite) TestDo_withClose() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("1"))
	})
	c := NewClient("", nil)
	req, err := c.NewRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	_, err = c.Do(req, nil, true)
	require.Nil(suite.T(), err)
}
func (suite *TestSuite) TestDo_withWriter() {
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
func (suite *TestSuite) TestDo_withEmptyBody() {
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

func (suite *TestSuite) TestDo_withError() {
	c := NewClient("", nil)
	req, err := c.NewRequest("GET", "", nil)
	require.Nil(suite.T(), err)
	_, err = c.Do(req, nil, false)
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestDo_withResponseError() {
	suite.Router.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("1"))
	})
	c := NewClient("", nil)
	req, err := c.NewRequest("GET", suite.Server.URL+"/dummy", nil)
	require.Nil(suite.T(), err)
	_, err = c.Do(req, nil, true)
	require.NotNil(suite.T(), err)
}
