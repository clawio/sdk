package sdk

import (
	"net/http"
)

// ServiceEndpoints contains the url of the different services of ClawIO.
type ServiceEndpoints struct {
	AuthServiceBaseURL string
	DataServiceBaseURL string
}

// SDK contains services used for talking to different parts of the ClawIO API.
type SDK struct {
	Auth AuthService
	Data DataService
}

// New creates a new SDK.
func New(urls *ServiceEndpoints, httpClient *http.Client) *SDK {
	sdk := &SDK{}
	sdk.Auth = &authService{client: newClient(urls.AuthServiceBaseURL, httpClient)}
	sdk.Data = &dataService{client: newClient(urls.DataServiceBaseURL, httpClient)}
	return sdk
}
