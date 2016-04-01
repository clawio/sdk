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
		w.Header().Set("checksum", "type:value")
		w.WriteHeader(201)
		fmt.Fprintln(w, "1")
	})
	data := strings.NewReader("1")
	checksum, _, err := suite.SDK.Data.Upload("test", data, "")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "type:value", checksum)
}

func (suite *TestSuite) TestUpload_withFailConnection() {
	_, _, err := suite.SDK.Data.Upload("test", nil, "")
	require.NotNil(suite.T(), err)
}
func (suite *TestSuite) TestDownload() {
	suite.Router.HandleFunc(defaultDataBaseURL+"download/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "1")
	})
	reader, _, err := suite.SDK.Data.Download("test")
	require.Nil(suite.T(), err)
	data, err := ioutil.ReadAll(reader)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "1", string(data))
}
func (suite *TestSuite) TestDownload_withFailConnection() {
	_, _, err := suite.SDK.Data.Download("test")
	require.NotNil(suite.T(), err)
}
