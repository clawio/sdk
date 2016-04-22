package sdk

import (
	"io"
	"path"

	"github.com/clawio/codes"
)

// DataService is the interface that specifies the methods to call a data service.
type DataService interface {
	Upload(pathSpec string, r io.Reader, checksum string) (*codes.Response, error)
	Download(pathSpec string) (io.Reader, *codes.Response, error)
}

type dataService struct {
	client  *client
	baseURL string
}

func (s *dataService) Upload(pathSpec string, r io.Reader, checksum string) (*codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newUploadRequest("upload"+pathSpec, r)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.do(req, nil, true)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *dataService) Download(pathSpec string) (io.Reader, *codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newRequest("GET", "download"+pathSpec, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.do(req, nil, false)
	if err != nil {
		return nil, resp, err
	}
	return resp.Body, resp, nil
}
