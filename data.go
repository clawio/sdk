package sdk

import (
	"io"
	p "path"

	"github.com/clawio/codes"
)

type DataService interface {
	Upload(path string, r io.Reader, checksum string) (string, *codes.Response, error)
	Download(path string) (io.Reader, *codes.Response, error)
}

type dataService struct {
	client  *Client
	baseURL string
}

func (s *dataService) Upload(path string, r io.Reader, checksum string) (string, *codes.Response, error) {
	path = p.Join("/", path)
	req, err := s.client.NewUploadRequest("upload"+path, r)
	if err != nil {
		return "", nil, err
	}
	resp, err := s.client.Do(req, nil, true)
	if err != nil {
		return "", resp, err
	}
	return resp.Header.Get("checksum"), resp, nil
}

func (s *dataService) Download(path string) (io.Reader, *codes.Response, error) {
	path = p.Join("/", path)
	req, err := s.client.NewRequest("GET", "download"+path, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.Do(req, nil, false)
	if err != nil {
		return nil, resp, err
	}
	return resp.Body, resp, nil
}
