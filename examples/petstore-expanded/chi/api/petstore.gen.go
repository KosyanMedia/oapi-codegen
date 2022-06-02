// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/KosyanMedia/oapi-codegen/v2 version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/KosyanMedia/oapi-codegen/v2/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// Error defines model for Error.
type Error struct {
	// Error code
	Code int32 `json:"code" validate:"required"`

	// Error message
	Message string `json:"message" validate:"required"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	// Name of the pet
	Name string `json:"name" validate:"required"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Unique id of the pet
	Id int64 `json:"id" validate:"required"`

	// Name of the pet
	Name string `json:"name" validate:"required"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {
	// tags to filter by
	Tags []string `form:"tags,omitempty" json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody = NewPet

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = AddPetJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams)
	// Creates a new pet
	// (POST /pets)
	AddPet(w http.ResponseWriter, r *http.Request)
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(w http.ResponseWriter, r *http.Request, id int64)
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(w http.ResponseWriter, r *http.Request, id int64)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// FindPets operation middleware
func (siw *ServerInterfaceWrapper) FindPets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams

	// ------------- Optional query parameter "tags" -------------
	if paramValue := r.URL.Query().Get("tags"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tags", r.URL.Query(), &params.Tags)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tags", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindPets(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AddPet operation middleware
func (siw *ServerInterfaceWrapper) AddPet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPet(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeletePet operation middleware
func (siw *ServerInterfaceWrapper) DeletePet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeletePet(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindPetByID operation middleware
func (siw *ServerInterfaceWrapper) FindPetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindPetByID(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/pets", wrapper.FindPets)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/pets", wrapper.AddPet)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/pets/{id}", wrapper.DeletePet)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/pets/{id}", wrapper.FindPetByID)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"W3073d3v/h0AAP//bOkdVFcSAAA=",
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
