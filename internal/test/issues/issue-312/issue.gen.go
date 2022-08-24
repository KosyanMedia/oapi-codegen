// Package issue_312 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/KosyanMedia/oapi-codegen/v2 version (devel) DO NOT EDIT.
package issue_312

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/KosyanMedia/oapi-codegen/v2/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// Error defines model for Error.
type Error struct {
	// Error code
	Code int32 `json:"code" validate:"required"`

	// Error message
	Message string `json:"message" validate:"required"`
}

// Pet defines model for Pet.
type Pet struct {
	// The name of the pet.
	Name string `json:"name" validate:"required"`
}

// PetNames defines model for PetNames.
type PetNames struct {
	// The names of the pets.
	Names []string `json:"names" validate:"required"`
}

// ValidatePetsJSONBody defines parameters for ValidatePets.
type ValidatePetsJSONBody = PetNames

// ValidatePetsJSONRequestBody defines body for ValidatePets for application/json ContentType.
type ValidatePetsJSONRequestBody = ValidatePetsJSONBody

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
	// GetPet request
	GetPet(ctx context.Context, petId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ValidatePets request with any body
	ValidatePetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ValidatePets(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetPet(ctx context.Context, petId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPetRequest(c.Server, petId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ValidatePetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewValidatePetsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ValidatePets(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewValidatePetsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetPetRequest generates requests for GetPet
func NewGetPetRequest(server string, petId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "petId", runtime.ParamLocationPath, petId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
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

// NewValidatePetsRequest calls the generic ValidatePets builder with application/json body
func NewValidatePetsRequest(server string, body ValidatePetsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewValidatePetsRequestWithBody(server, "application/json", bodyReader)
}

// NewValidatePetsRequestWithBody generates requests for ValidatePets with any type of body
func NewValidatePetsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets:validate")
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
	// GetPet request
	GetPetWithResponse(ctx context.Context, petId string, reqEditors ...RequestEditorFn) (*ClientGetPetResponse, error)

	// ValidatePets request with any body
	ValidatePetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ClientValidatePetsResponse, error)

	ValidatePetsWithResponse(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*ClientValidatePetsResponse, error)
}

type ClientGetPetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
}

// Status returns HTTPResponse.Status
func (r ClientGetPetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ClientGetPetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ResponseBody returns HTTPResponse.Body as byte array
func (r ClientGetPetResponse) ResponseBody() []byte {
	return r.Body
}

type ClientValidatePetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      []Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r ClientValidatePetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ClientValidatePetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ResponseBody returns HTTPResponse.Body as byte array
func (r ClientValidatePetsResponse) ResponseBody() []byte {
	return r.Body
}

// GetPetWithResponse request returning *GetPetResponse
func (c *ClientWithResponses) GetPetWithResponse(ctx context.Context, petId string, reqEditors ...RequestEditorFn) (*ClientGetPetResponse, error) {
	rsp, err := c.GetPet(ctx, petId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPetResponse(rsp)
}

// ValidatePetsWithBodyWithResponse request with arbitrary body returning *ValidatePetsResponse
func (c *ClientWithResponses) ValidatePetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ClientValidatePetsResponse, error) {
	rsp, err := c.ValidatePetsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseValidatePetsResponse(rsp)
}

func (c *ClientWithResponses) ValidatePetsWithResponse(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*ClientValidatePetsResponse, error) {
	rsp, err := c.ValidatePets(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseValidatePetsResponse(rsp)
}

// ParseGetPetResponse parses an HTTP response from a GetPetWithResponse call
func ParseGetPetResponse(rsp *http.Response) (*ClientGetPetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ClientGetPetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseValidatePetsResponse parses an HTTP response from a ValidatePetsWithResponse call
func ParseValidatePetsResponse(rsp *http.Response) (*ClientValidatePetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ClientValidatePetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get pet given identifier.
	// (GET /pets/{petId})
	GetPet(ctx echo.Context, petId string) (resp *GetPetResponse, err error)
	// Validate pets
	// (POST /pets:validate)
	ValidatePets(ctx echo.Context, requestBody ValidatePetsJSONBody) (resp *ValidatePetsResponse, err error)
}

type GetPetResponse struct {
	Code    int
	JSON200 *Pet
}

type ValidatePetsResponse struct {
	Code        int
	JSON200     []Pet
	JSONDefault *Error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPet converts echo context to params.
func (w *ServerInterfaceWrapper) GetPet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetPet(ctx, petId)

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

// ValidatePets converts echo context to params.
func (w *ServerInterfaceWrapper) ValidatePets(ctx echo.Context) error {
	var err error

	var requestBody ValidatePetsJSONBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.ValidatePets(ctx, requestBody)

	if err != nil {
		return err
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	if response.JSONDefault != nil {
		return ctx.JSON(response.Code, response.JSONDefault)
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

	router.GET(baseURL+"/pets/:petId", wrapper.GetPet, m...)
	router.POST(baseURL+"/pets:validate", wrapper.ValidatePets, m...)

}
