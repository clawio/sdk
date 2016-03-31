package sdk

import (
	"io"

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
	return "", nil, nil
}

func (s *dataService) Download(path string) (io.Reader, *codes.Response, error) {
	return nil, nil, nil
}
