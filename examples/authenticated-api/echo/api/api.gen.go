// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/KosyanMedia/oapi-codegen/v2 version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/KosyanMedia/oapi-codegen/v2/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	// Error code
	Code int32 `json:"code" validate:"required"`

	// Error message
	Message string `json:"message" validate:"required"`
}

// Thing defines model for Thing.
type Thing struct {
	Name string `json:"name" validate:"required"`
}

// ThingWithID defines model for ThingWithID.
type ThingWithID struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// AddThingJSONBody defines parameters for AddThing.
type AddThingJSONBody = Thing

// AddThingJSONRequestBody defines body for AddThing for application/json ContentType.
type AddThingJSONRequestBody = AddThingJSONBody

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
	// ListThings request
	ListThings(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AddThing request with any body
	AddThingWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AddThing(ctx context.Context, body AddThingJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListThings(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListThingsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddThingWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddThingRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddThing(ctx context.Context, body AddThingJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddThingRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListThingsRequest generates requests for ListThings
func NewListThingsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/things")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddThingRequest calls the generic AddThing builder with application/json body
func NewAddThingRequest(server string, body AddThingJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddThingRequestWithBody(server, "application/json", bodyReader)
}

// NewAddThingRequestWithBody generates requests for AddThing with any type of body
func NewAddThingRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/things")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
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
	// ListThings request
	ListThingsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ClientListThingsResponse, error)

	// AddThing request with any body
	AddThingWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ClientAddThingResponse, error)

	AddThingWithResponse(ctx context.Context, body AddThingJSONRequestBody, reqEditors ...RequestEditorFn) (*ClientAddThingResponse, error)
}

type ClientListThingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      []ThingWithID
}

// Status returns HTTPResponse.Status
func (r ClientListThingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ClientListThingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ResponseBody returns HTTPResponse.Body as byte array
func (r ClientListThingsResponse) ResponseBody() []byte {
	return r.Body
}

// RawResponse returns pointer to the raw http.Response
func (r ClientListThingsResponse) RawResponse() *http.Response {
	return r.HTTPResponse
}

type ClientAddThingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ThingWithID
}

// Status returns HTTPResponse.Status
func (r ClientAddThingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ClientAddThingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ResponseBody returns HTTPResponse.Body as byte array
func (r ClientAddThingResponse) ResponseBody() []byte {
	return r.Body
}

// RawResponse returns pointer to the raw http.Response
func (r ClientAddThingResponse) RawResponse() *http.Response {
	return r.HTTPResponse
}

// ListThingsWithResponse request returning *ListThingsResponse
func (c *ClientWithResponses) ListThingsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ClientListThingsResponse, error) {
	rsp, err := c.ListThings(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListThingsResponse(rsp)
}

// AddThingWithBodyWithResponse request with arbitrary body returning *AddThingResponse
func (c *ClientWithResponses) AddThingWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ClientAddThingResponse, error) {
	rsp, err := c.AddThingWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddThingResponse(rsp)
}

func (c *ClientWithResponses) AddThingWithResponse(ctx context.Context, body AddThingJSONRequestBody, reqEditors ...RequestEditorFn) (*ClientAddThingResponse, error) {
	rsp, err := c.AddThing(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddThingResponse(rsp)
}

// ParseListThingsResponse parses an HTTP response from a ListThingsWithResponse call
func ParseListThingsResponse(rsp *http.Response) (*ClientListThingsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ClientListThingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []ThingWithID
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = dest

	}

	return response, nil
}

// ParseAddThingResponse parses an HTTP response from a AddThingWithResponse call
func ParseAddThingResponse(rsp *http.Response) (*ClientAddThingResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ClientAddThingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ThingWithID
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /things)
	ListThings(ctx echo.Context) (resp *ListThingsResponse, err error)

	// (POST /things)
	AddThing(ctx echo.Context, requestBody AddThingJSONBody) (resp *AddThingResponse, err error)
}

type ListThingsResponse struct {
	Code    int
	JSON200 []ThingWithID
}

type AddThingResponse struct {
	Code    int
	JSON201 *ThingWithID
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListThings converts echo context to params.
func (w *ServerInterfaceWrapper) ListThings(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.ListThings(ctx)

	if err != nil {
		return err
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// AddThing converts echo context to params.
func (w *ServerInterfaceWrapper) AddThing(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{"things:w"})

	var requestBody AddThingJSONBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.AddThing(ctx, requestBody)

	if err != nil {
		return err
	}

	if response.JSON201 != nil {
		if response.Code == 0 {
			response.Code = 201
		}
		return ctx.JSON(response.Code, response.JSON201)
	}
	return ctx.NoContent(response.Code)
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/things", wrapper.ListThings, m...)
	router.POST(baseURL+"/things", wrapper.AddThing, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUQWsbPRD9K8N8H/Sy2E5SetibQ1JwKLS0hhziQJTV2Kt2LSmj2bgm7H8vkry2m03T",
	"FnqxV9LozZs3b/SElVt7Z8lKwPIJQ1XTWqXPS2bH8cOz88RiKG1XTlP81xQqNl6Ms1jmYEhnBS4dr5Vg",
	"icbK2SkWKFtPeUkrYuwKXFMIavVLoP54fzUIG7vCriuQ6aE1TBrLG9wl7MNvuwLndQwc0LZqnbK9jpei",
	"9ijXRurZRbylmubjEsubJ/yfaYkl/jc+6DbeiTbOqbvieW6j4++xKu/evqDKMy5G4213G3cDVS0b2X6J",
	"eTLkOSkmnrZSx9V9Wr3vE1xdz7HIrYwJ8ukhYS3isYvAxi7dsAVTC/RdrX1DMP00g01tqhraQAEyEoj7",
	"RhZC5TwFUFbD1fUcVORSoBhpYpJIjayYSgnphHOZMbHAR+KQU52MJqNJ9IPzZJU3WOJZ2irQK6lTqWOJ",
	"sqbPFcmQ7meSlm0ABY0JAm4J+cIIzqlSbaC4DkBWe2esgHYU7BsB90jMRsdjWthV4+5VA73UBRiBXTci",
	"dKxw6ThVuSvLODtaWEzcOS1nGkv8YILMM+PYz+CdDblnp5NJHiArZFMhyvtmBzX+GmI1/QQm2wit08Xf",
	"em5n1G7fYsWstrnHP4v1XKQc4114Qdip1rH0FAjiok4DiefH0oa9piEFZ00XthcVsiWTZY60vctg5eYu",
	"ewqMBcc6GQ08cRwcUAu7YSP0kuRTrfPo5QGiIOdOb/9K6z8Y66GY04M2xgZiGUFvxlh+3iO9i9oYqUFZ",
	"mF3g8aALt9QNnHLyb9n3BhnWMD8mOj8iCq01Dy1BunZ4g9IjePz63GDfvvxcvRobI34EAAD//590s8Nz",
	"BgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
