package sdk

import (
	"path"

	"github.com/clawio/codes"
	"github.com/clawio/entities"
)

// MetaDataService is the interface that specifies the methods to call a metaData service.
type MetaDataService interface {
	Init() (*codes.Response, error)
	ExamineObject(pathSpec string) (entities.ObjectInfo, *codes.Response, error)
	ListTree(pathSpec string) ([]entities.ObjectInfo, *codes.Response, error)
}

type objectInfo struct {
	PathSpec string              `json:"pathspec"`
	Type     entities.ObjectType `json:"type"`
	Size     uint64              `json:"size"`
	MimeType string              `json:"mimetype"`
	Checksum string              `json:"checksum"`
}

func (o *objectInfo) GetPathSpec() string          { return o.PathSpec }
func (o *objectInfo) GetType() entities.ObjectType { return o.Type }
func (o *objectInfo) GetSize() uint64              { return o.Size }
func (o *objectInfo) GetMimeType() string          { return o.MimeType }
func (o *objectInfo) GetChecksum() string          { return o.Checksum }

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

func (s *metaDataService) ExamineObject(pathSpec string) (entities.ObjectInfo, *codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newRequest("GET", "examine"+pathSpec, nil)
	if err != nil {
		return nil, nil, err
	}
	o := &objectInfo{}
	resp, err := s.client.do(req, o, true)
	if err != nil {
		return o, resp, err
	}
	return o, resp, nil
}

func (s *metaDataService) ListTree(pathSpec string) ([]entities.ObjectInfo, *codes.Response, error) {
	pathSpec = path.Join("/", pathSpec)
	req, err := s.client.newRequest("GET", "listtree"+pathSpec, nil)
	if err != nil {
		return nil, nil, err
	}
	var oinfos []entities.ObjectInfo
	var infos []*objectInfo
	resp, err := s.client.do(req, &infos, true)
	if err != nil {
		return nil, resp, err
	}
	for _, info := range infos {
		oinfos = append(oinfos, info)
	}
	return oinfos, resp, nil
}
