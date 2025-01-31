// Package issue_52 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/KosyanMedia/oapi-codegen/v2 version (devel) DO NOT EDIT.
package issue_52

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ArrayValue defines model for ArrayValue.
type ArrayValue []Value

// Document defines model for Document.
type Document struct {
	Fields map[string]Value `json:"fields,omitempty"`
}

// Value defines model for Value.
type Value struct {
	ArrayValue  ArrayValue `json:"arrayValue,omitempty"`
	StringValue *string    `json:"stringValue,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /example)
	ExampleGet(ctx echo.Context) (resp *ExampleGetResponse, err error)
}

type ExampleGetResponse struct {
	Code    int
	JSON200 *Document
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ExampleGet converts echo context to params.
func (w *ServerInterfaceWrapper) ExampleGet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.ExampleGet(ctx)

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

	router.GET(baseURL+"/example", wrapper.ExampleGet, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/5RSwU7rMBD8lWjfO0ZJXt/NNyQQQgjBiROXxd42Lo5t2ZuKqsq/o3XaQgQCcYo92Zmd",
	"He8BdBhi8OQ5gzpA1j0NWI4XKeH+Ed1IcrNMQ4H/JlqDgj/tO7E9stq5eqqB95FAAYqE3C+DHgfyLAIx",
	"hUiJLRW5tSVnygmNsWyDR/ewqPhNw/C8Jc0wfUZqOI+yNICLMb9r9iGQqYbMyfrNmXhsN6NfGRDI+nWQ",
	"YkNZJxtlWlBwhy9U5TFRxT1ylUiPKdsdVaKQK0xU9eiNI1PN1t3+yUMNbNlJB3rFITqCGnaU8qzZNV3z",
	"T2yGSB6jBQX/m65ZQQ0RuS+TtyeiOsCGytuIOoqtGwMKrub/18RQQ6Icg89zaKuuk48Ono+vijE6qwu3",
	"3WbxcFqmn2I970aJaBnN/a2g0zS9BQAA//+hEzLlqAIAAA==",
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
