package sdk

import (
	"io"
	p "path"

	"github.com/clawio/codes"
)

// DataService is the interface that specifies the methods to call a data service.
type DataService interface {
	Upload(path string, r io.Reader, checksum string) (*codes.Response, error)
	Download(path string) (io.Reader, *codes.Response, error)
}

type dataService struct {
	client  *client
	baseURL string
}

func (s *dataService) Upload(path string, r io.Reader, checksum string) (*codes.Response, error) {
	path = p.Join("/", path)
	req, err := s.client.newUploadRequest("upload"+path, r)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.do(req, nil, true)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *dataService) Download(path string) (io.Reader, *codes.Response, error) {
	path = p.Join("/", path)
	req, err := s.client.newRequest("GET", "download"+path, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.do(req, nil, false)
	if err != nil {
		return nil, resp, err
	}
	return resp.Body, resp, nil
}
