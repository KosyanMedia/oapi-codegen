// Package operations provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/KosyanMedia/oapi-codegen/v2 version (devel) DO NOT EDIT.
package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/KosyanMedia/oapi-codegen/v2/pkg/runtime"
)

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName" validate:"required"`
	Role      string `json:"role" validate:"required"`
}

// FirstParam defines model for FirstParam.
type FirstParam = string

// SecondParam defines model for SecondParam.
type SecondParam = string

// LeavePostOnlyPostParams defines parameters for LeavePostOnlyPost.
type LeavePostOnlyPostParams struct {
	First  *FirstParam  `form:"first,omitempty" json:"first,omitempty"`
	Second *SecondParam `form:"second,omitempty" json:"second,omitempty"`
}

// ShouldHaveBothParams defines parameters for ShouldHaveBoth.
type ShouldHaveBothParams struct {
	First  *FirstParam  `form:"first,omitempty" json:"first,omitempty"`
	Second *SecondParam `form:"second,omitempty" json:"second,omitempty"`
}

// ShouldHaveSecondParams defines parameters for ShouldHaveSecond.
type ShouldHaveSecondParams struct {
	Second *SecondParam `form:"second,omitempty" json:"second,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// LeavePostOnlyPost request
	LeavePostOnlyPost(ctx context.Context, params LeavePostOnlyPostParams) (*http.Response, error)

	// ShouldHaveBoth request
	ShouldHaveBoth(ctx context.Context, params ShouldHaveBothParams) (*http.Response, error)

	// ShouldHaveSecond request
	ShouldHaveSecond(ctx context.Context, params ShouldHaveSecondParams) (*http.Response, error)
}

func (c *Client) LeavePostOnlyPost(ctx context.Context, params LeavePostOnlyPostParams) (*http.Response, error) {
	req, err := NewLeavePostOnlyPostRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ShouldHaveBoth(ctx context.Context, params ShouldHaveBothParams) (*http.Response, error) {
	req, err := NewShouldHaveBothRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ShouldHaveSecond(ctx context.Context, params ShouldHaveSecondParams) (*http.Response, error) {
	req, err := NewShouldHaveSecondRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewLeavePostOnlyPostRequest generates requests for LeavePostOnlyPost
func NewLeavePostOnlyPostRequest(server string, params LeavePostOnlyPostParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/leave_post_only")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.First != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", false, "first", runtime.ParamLocationQuery, *params.First); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Second != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", false, "second", runtime.ParamLocationQuery, *params.Second); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewShouldHaveBothRequest generates requests for ShouldHaveBoth
func NewShouldHaveBothRequest(server string, params ShouldHaveBothParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/should_have_both")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.First != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", false, "first", runtime.ParamLocationQuery, *params.First); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Second != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", false, "second", runtime.ParamLocationQuery, *params.Second); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewShouldHaveSecondRequest generates requests for ShouldHaveSecond
func NewShouldHaveSecondRequest(server string, params ShouldHaveSecondParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/should_have_second")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Second != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", false, "second", runtime.ParamLocationQuery, *params.Second); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// LeavePostOnlyPost request
	LeavePostOnlyPostWithResponse(ctx context.Context, params LeavePostOnlyPostParams) (*ClientLeavePostOnlyPostResponse, error)

	// ShouldHaveBoth request
	ShouldHaveBothWithResponse(ctx context.Context, params ShouldHaveBothParams) (*ClientShouldHaveBothResponse, error)

	// ShouldHaveSecond request
	ShouldHaveSecondWithResponse(ctx context.Context, params ShouldHaveSecondParams) (*ClientShouldHaveSecondResponse, error)
}

type ClientLeavePostOnlyPostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SchemaObject
}

// Status returns HTTPResponse.Status
func (r ClientLeavePostOnlyPostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ClientLeavePostOnlyPostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ResponseBody returns HTTPResponse.Body as byte array
func (r ClientLeavePostOnlyPostResponse) ResponseBody() []byte {
	return r.Body
}

// RawResponse returns pointer to the raw http.Response
func (r ClientLeavePostOnlyPostResponse) RawResponse() *http.Response {
	return r.HTTPResponse
}

type ClientShouldHaveBothResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SchemaObject
}

// Status returns HTTPResponse.Status
func (r ClientShouldHaveBothResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ClientShouldHaveBothResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ResponseBody returns HTTPResponse.Body as byte array
func (r ClientShouldHaveBothResponse) ResponseBody() []byte {
	return r.Body
}

// RawResponse returns pointer to the raw http.Response
func (r ClientShouldHaveBothResponse) RawResponse() *http.Response {
	return r.HTTPResponse
}

type ClientShouldHaveSecondResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SchemaObject
}

// Status returns HTTPResponse.Status
func (r ClientShouldHaveSecondResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ClientShouldHaveSecondResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ResponseBody returns HTTPResponse.Body as byte array
func (r ClientShouldHaveSecondResponse) ResponseBody() []byte {
	return r.Body
}

// RawResponse returns pointer to the raw http.Response
func (r ClientShouldHaveSecondResponse) RawResponse() *http.Response {
	return r.HTTPResponse
}

// LeavePostOnlyPostWithResponse request returning *LeavePostOnlyPostResponse
func (c *ClientWithResponses) LeavePostOnlyPostWithResponse(ctx context.Context, params LeavePostOnlyPostParams) (*ClientLeavePostOnlyPostResponse, error) {
	rsp, err := c.LeavePostOnlyPost(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseLeavePostOnlyPostResponse(rsp)
}

// ShouldHaveBothWithResponse request returning *ShouldHaveBothResponse
func (c *ClientWithResponses) ShouldHaveBothWithResponse(ctx context.Context, params ShouldHaveBothParams) (*ClientShouldHaveBothResponse, error) {
	rsp, err := c.ShouldHaveBoth(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseShouldHaveBothResponse(rsp)
}

// ShouldHaveSecondWithResponse request returning *ShouldHaveSecondResponse
func (c *ClientWithResponses) ShouldHaveSecondWithResponse(ctx context.Context, params ShouldHaveSecondParams) (*ClientShouldHaveSecondResponse, error) {
	rsp, err := c.ShouldHaveSecond(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseShouldHaveSecondResponse(rsp)
}

// ParseLeavePostOnlyPostResponse parses an HTTP response from a LeavePostOnlyPostWithResponse call
func ParseLeavePostOnlyPostResponse(rsp *http.Response) (*ClientLeavePostOnlyPostResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ClientLeavePostOnlyPostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SchemaObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseShouldHaveBothResponse parses an HTTP response from a ShouldHaveBothWithResponse call
func ParseShouldHaveBothResponse(rsp *http.Response) (*ClientShouldHaveBothResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ClientShouldHaveBothResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SchemaObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseShouldHaveSecondResponse parses an HTTP response from a ShouldHaveSecondWithResponse call
func ParseShouldHaveSecondResponse(rsp *http.Response) (*ClientShouldHaveSecondResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ClientShouldHaveSecondResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SchemaObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
