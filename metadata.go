package sdk

import (
	"path"

	"github.com/clawio/clawiod/codes"
	"github.com/clawio/clawiod/entities"
)

// MetaDataService is the interface that specifies the methods to call a metaData service.
type MetaDataService interface {
	Init() (*codes.Response, error)
	CreateTree(pathSpec string) (*codes.Response, error)
	ExamineObject(pathSpec string) (*entities.ObjectInfo, *codes.Response, error)
	ListTree(pathSpec string) ([]*entities.ObjectInfo, *codes.Response, error)
	DeleteObject(pathSpec string) (*codes.Response, error)
	MoveObject(sourcePathSpec, targetPathSpec string) (*codes.Response, error)
}

type metaDataService struct {
	client  *client
	baseURL string
}

func (s *metaDataService) Init() (*codes.Response, error) {
	req, err := s.client.newRequest("POST", "init", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.do(req, nil, true)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *metaDataService) ExamineObject(pathSpec string) (*entities.ObjectInfo, *codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newRequest("GET", "examine"+pathSpec, nil)
	if err != nil {
		return nil, nil, err
	}
	o := &entities.ObjectInfo{}
	resp, err := s.client.do(req, o, true)
	if err != nil {
		return o, resp, err
	}
	return o, resp, nil
}

func (s *metaDataService) ListTree(pathSpec string) ([]*entities.ObjectInfo, *codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newRequest("GET", "list"+pathSpec, nil)
	if err != nil {
		return nil, nil, err
	}
	var oinfos []*entities.ObjectInfo
	resp, err := s.client.do(req, &oinfos, true)
	if err != nil {
		return nil, resp, err
	}
	return oinfos, resp, nil
}

func (s *metaDataService) DeleteObject(pathSpec string) (*codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newRequest("DELETE", "delete"+pathSpec, nil)
	if err != nil {
		return nil, err
	}
	return s.client.do(req, nil, true)
}

func (s *metaDataService) CreateTree(pathSpec string) (*codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newRequest("POST", "createtree"+pathSpec, nil)
	if err != nil {
		return nil, err
	}
	return s.client.do(req, nil, true)
}

func (s *metaDataService) MoveObject(sourcePathSpec, targetPathSpec string) (*codes.Response, error) {
	sourcePathSpec = path.Join("/", sourcePathSpec)
	targetPathSpec = path.Join("/", targetPathSpec)
	req, err := s.client.newRequest("POST", "move"+sourcePathSpec, nil)
	if err != nil {
		return nil, err
	}
	values := req.URL.Query()
	values.Set("target", targetPathSpec)
	req.URL.RawQuery = values.Encode()
	return s.client.do(req, nil, true)
}
