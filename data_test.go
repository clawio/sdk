package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestUpload() {
	suite.Router.HandleFunc(defaultDataBaseURL+"upload/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	})
	data := strings.NewReader("1")
	resp, err := suite.SDK.Data.Upload("test", data, "")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusCreated, resp.StatusCode)
}
func (suite *TestSuite) TestUpload_withError() {
	suite.Router.HandleFunc(defaultDataBaseURL+"upload/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"code":99, "message":""}`)
	})
	data := strings.NewReader("1")
	resp, err := suite.SDK.Data.Upload("test", data, "")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}

func (suite *TestSuite) TestDownload() {
	suite.Router.HandleFunc(defaultDataBaseURL+"download/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("1"))
	})
	reader, resp, err := suite.SDK.Data.Download("test")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	data, err := ioutil.ReadAll(reader)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), data, []byte("1"))
}

func (suite *TestSuite) TestDownload_withError() {
	suite.Router.HandleFunc(defaultDataBaseURL+"download/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"code":99, "message":""}`)
	})
	_, resp, err := suite.SDK.Data.Download("test")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}
