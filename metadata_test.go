package sdk

import (
	"fmt"
	"net/http"

	"github.com/clawio/entities"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestInit() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"init", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	resp, err := suite.SDK.Meta.Init()
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}
func (suite *TestSuite) TestInit_withError() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"init", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"code":99, "message":""}`)
	})
	resp, err := suite.SDK.Meta.Init()
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}

func (suite *TestSuite) TestExamineObject() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"examine/myblob", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, `{"pathspec":"myblob", "size": 100, "type": 1, "mime": "", "checksum": ""}`)
	})
	info, resp, err := suite.SDK.Meta.ExamineObject("myblob")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	require.Equal(suite.T(), "myblob", info.GetPathSpec())
	fmt.Printf("%+v", info)
	require.Equal(suite.T(), uint64(100), info.GetSize())
	require.Equal(suite.T(), entities.ObjectType(1), info.GetType())
	require.Equal(suite.T(), "", info.GetMimeType())
	require.Equal(suite.T(), "", info.GetChecksum())
}
func (suite *TestSuite) TestExamineObject_withError() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"examine/myblob", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"code":99, "message":""}`)
	})
	_, resp, err := suite.SDK.Meta.ExamineObject("myblob")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}

func (suite *TestSuite) TestListTree() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"listtree/tree", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, `[{"pathspec":"myblob", "size": 100, "type": 1, "mime": "", "checksum": ""}]`)
	})
	infos, resp, err := suite.SDK.Meta.ListTree("tree")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	require.Equal(suite.T(), 1, len(infos))
}
func (suite *TestSuite) TestListTree_withError() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"listtree/tree", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"code":99, "message":""}`)
	})
	_, resp, err := suite.SDK.Meta.ListTree("tree")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}
func (suite *TestSuite) TestDeleteObject() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"delete/tree", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	resp, err := suite.SDK.Meta.DeleteObject("tree")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}
func (suite *TestSuite) TestDeleteObject_withError() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"delete/tree", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	resp, err := suite.SDK.Meta.DeleteObject("tree")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusInternalServerError, resp.StatusCode)
}
func (suite *TestSuite) TestMoveObject() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"move/tree", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	resp, err := suite.SDK.Meta.MoveObject("tree", "newtree")
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}
func (suite *TestSuite) TestMoveObject_withError() {
	suite.Router.HandleFunc(defaultMetaDataBaseURL+"move/tree", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	resp, err := suite.SDK.Meta.MoveObject("tree", "newtree")
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), http.StatusInternalServerError, resp.StatusCode)
}
