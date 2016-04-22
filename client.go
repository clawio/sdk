package sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/clawio/codes"
)

// A client manages communication with the ClawIO API.
type client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
}

// Newclient returns a new client.
func newClient(baseURL string, httpClient *http.Client) *client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	u, _ := url.Parse(baseURL)
	return &client{httpClient: httpClient, baseURL: u}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the baseURL of the client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// TODO(labkode) Check for trailing slash and allow it to access home directories
	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if e := json.NewEncoder(buf).Encode(body); e != nil {
			return nil, e
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.userAgent != "" {
		req.Header.Add("User-Agent", c.userAgent)
	}
	return req, nil
}

// NewUploadRequest creates an upload request. A relative URL can be provided in
// urlStr, in which case it is resolved relative to the baseURL of the client.
// Relative URLs should always be specified without a preceding slash.
func (c *client) newUploadRequest(urlStr string, reader io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)
	req, err := http.NewRequest("PUT", u.String(), reader)
	if err != nil {
		return nil, err
	}
	//req.ContentLength = size

	//if len(mediaType) == 0 {
	//	mediaType = defaultMediaType
	//}
	//req.Header.Add("Content-Type", mediaType)
	//req.Header.Add("Accept", mediaTypeV3)
	if c.userAgent != "" {
		req.Header.Add("User-Agent", c.userAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *client) do(req *http.Request, v interface{}, close bool) (*codes.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if close == true {
			// Drain and close the body to let the Transport reuse the connection
			// This is valid when the response is decoded into v
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
	}()

	response := codes.NewResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}
	return response, err
}

// checkResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &codes.ErrorResponse{Response: r, Err: &codes.Err{}}
	if r.StatusCode == http.StatusBadRequest {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorResponse.Err = codes.NewErr(codes.Internal, "error reading body for errorResponse")
			return errorResponse
		}
		if e := json.Unmarshal(data, errorResponse.Err); e != nil {
			errorResponse.Err = codes.NewErr(codes.Internal, "error response is not valid JSON")
		}
		return errorResponse
	}
	errorResponse.Err = codes.NewErr(codes.Internal, "")
	return errorResponse
}
