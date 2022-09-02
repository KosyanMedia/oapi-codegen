// Package components provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/KosyanMedia/oapi-codegen/v2 version (devel) DO NOT EDIT.
package components

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/KosyanMedia/oapi-codegen/v2/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Defines values for Enum1.
const (
	Enum1One   Enum1 = "One"
	Enum1Three Enum1 = "Three"
	Enum1Two   Enum1 = "Two"
)

// Defines values for Enum2.
const (
	Enum2One   Enum2 = "One"
	Enum2Three Enum2 = "Three"
	Enum2Two   Enum2 = "Two"
)

// Defines values for Enum3.
const (
	Enum3Bar      Enum3 = "Bar"
	Enum3Enum1One Enum3 = "Enum1One"
	Enum3Foo      Enum3 = "Foo"
)

// Defines values for Enum4.
const (
	Cat   Enum4 = "Cat"
	Dog   Enum4 = "Dog"
	Mouse Enum4 = "Mouse"
)

// Defines values for Enum5.
const (
	N5 Enum5 = 5
	N6 Enum5 = 6
	N7 Enum5 = 7
)

// Defines values for EnumParam1.
const (
	EnumParam1Both  EnumParam1 = "both"
	EnumParam1False EnumParam1 = "false"
	EnumParam1True  EnumParam1 = "true"
)

// Defines values for EnumParam2.
const (
	EnumParam2Both  EnumParam2 = "both"
	EnumParam2False EnumParam2 = "false"
	EnumParam2True  EnumParam2 = "true"
)

// Defines values for EnumParam3.
const (
	Alice EnumParam3 = "alice"
	Bob   EnumParam3 = "bob"
	Eve   EnumParam3 = "eve"
)

// Has additional properties of type int
type AdditionalPropertiesObject1 struct {
	Id                   int            `json:"id" validate:"required"`
	Name                 string         `json:"name" validate:"required"`
	Optional             *string        `json:"optional,omitempty"`
	AdditionalProperties map[string]int `json:"-"`
}

// Does not allow additional properties
type AdditionalPropertiesObject2 struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// Allows any additional property
type AdditionalPropertiesObject3 struct {
	Name                 string                 `json:"name" validate:"required"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Has anonymous field which has additional properties
type AdditionalPropertiesObject4 struct {
	Inner                AdditionalPropertiesObject4_Inner `json:"inner" validate:"required"`
	Name                 string                            `json:"name" validate:"required"`
	AdditionalProperties map[string]interface{}            `json:"-"`
}

// AdditionalPropertiesObject4_Inner defines model for AdditionalPropertiesObject4.Inner.
type AdditionalPropertiesObject4_Inner struct {
	Name                 string                 `json:"name" validate:"required"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Has additional properties with schema for dictionaries
type AdditionalPropertiesObject5 map[string]SchemaObject

// Ensureeverythingisreferenced200JSONResponseBodySchema defines model for Ensureeverythingisreferenced200JSONResponseBodySchema.
type Ensureeverythingisreferenced200JSONResponseBodySchema struct {
	// Has additional properties with schema for dictionaries
	Five AdditionalPropertiesObject5 `json:"five,omitempty"`

	// Has anonymous field which has additional properties
	Four      *AdditionalPropertiesObject4 `json:"four,omitempty"`
	JsonField *ObjectWithJsonField         `json:"jsonField,omitempty"`

	// Has additional properties of type int
	One *AdditionalPropertiesObject1 `json:"one,omitempty"`

	// Allows any additional property
	Three *AdditionalPropertiesObject3 `json:"three,omitempty"`

	// Does not allow additional properties
	Two *AdditionalPropertiesObject2 `json:"two,omitempty"`
}

// Conflicts with Enum2, enum values need to be prefixed with type
// name.
type Enum1 string

// Conflicts with Enum1, enum values need to be prefixed with type
// name.
type Enum2 string

// Enum values conflict with Enums above, need to be prefixed
// with type name.
type Enum3 string

// No conflicts here, should have unmodified enums
type Enum4 string

// Numerical enum
type Enum5 int

// ObjectWithJsonField defines model for ObjectWithJsonField.
type ObjectWithJsonField struct {
	Name   string          `json:"name" validate:"required"`
	Value1 json.RawMessage `json:"value1" validate:"required"`
	Value2 json.RawMessage `json:"value2,omitempty"`
}

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName" validate:"required"`

	// This property is required and readOnly, so the go model should have it as a pointer,
	// as it will not be included when it is sent from client to server.
	ReadOnlyRequiredProp  *string `json:"readOnlyRequiredProp,omitempty" validate:"required"`
	Role                  string  `json:"role" validate:"required"`
	WriteOnlyRequiredProp *int    `json:"writeOnlyRequiredProp,omitempty" validate:"required"`
}

// EnumParam1 defines model for EnumParam1.
type EnumParam1 string

// EnumParam2 defines model for EnumParam2.
type EnumParam2 string

// EnumParam3 defines model for EnumParam3.
type EnumParam3 string

// a parameter
type ParameterObject string

// ResponseObject defines model for ResponseObject.
type ResponseObject struct {
	Field SchemaObject `json:"Field" validate:"required"`
}

// RequestBody defines model for RequestBody.
type RequestBody struct {
	Field SchemaObject `json:"Field" validate:"required"`
}

// ParamsWithAddPropsParams_P1 defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P1 map[string]interface{}

// ParamsWithAddPropsParams defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams struct {
	// This parameter has additional properties
	P1 ParamsWithAddPropsParams_P1 `json:"p1" validate:"required"`

	// This parameter has an anonymous inner property which needs to be
	// turned into a proper type for additionalProperties to work
	P2 struct {
		Inner map[string]string `json:"inner" validate:"required"`
	} `form:"p2" json:"p2" validate:"required"`
}

// BodyWithAddPropsJSONBody defines parameters for BodyWithAddProps.
type BodyWithAddPropsJSONBody struct {
	Inner                map[string]int         `json:"inner" validate:"required"`
	Name                 string                 `json:"name" validate:"required"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// EnsureEverythingIsReferencedJSONRequestBody defines body for EnsureEverythingIsReferenced for application/json ContentType.
type EnsureEverythingIsReferencedJSONRequestBody RequestBody

// BodyWithAddPropsJSONRequestBody defines body for BodyWithAddProps for application/json ContentType.
type BodyWithAddPropsJSONRequestBody BodyWithAddPropsJSONBody

// Getter for additional properties for BodyWithAddPropsJSONBody. Returns the specified
// element and whether it was found
func (a BodyWithAddPropsJSONBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BodyWithAddPropsJSONBody
func (a *BodyWithAddPropsJSONBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a *BodyWithAddPropsJSONBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return fmt.Errorf("error reading 'inner': %w", err)
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a BodyWithAddPropsJSONBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'inner': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject1. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject1) Get(fieldName string) (value int, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject1
func (a *AdditionalPropertiesObject1) Set(fieldName string, value int) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]int)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle AdditionalProperties
func (a *AdditionalPropertiesObject1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.Id)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if raw, found := object["optional"]; found {
		err = json.Unmarshal(raw, &a.Optional)
		if err != nil {
			return fmt.Errorf("error reading 'optional': %w", err)
		}
		delete(object, "optional")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]int)
		for fieldName, fieldBuf := range object {
			var fieldVal int
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle AdditionalProperties
func (a AdditionalPropertiesObject1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["id"], err = json.Marshal(a.Id)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	if a.Optional != nil {
		object["optional"], err = json.Marshal(a.Optional)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'optional': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject3. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject3) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject3
func (a *AdditionalPropertiesObject3) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a *AdditionalPropertiesObject3) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a AdditionalPropertiesObject3) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4
func (a *AdditionalPropertiesObject4) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a *AdditionalPropertiesObject4) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return fmt.Errorf("error reading 'inner': %w", err)
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a AdditionalPropertiesObject4) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'inner': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4_Inner. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4_Inner) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4_Inner
func (a *AdditionalPropertiesObject4_Inner) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a *AdditionalPropertiesObject4_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a AdditionalPropertiesObject4_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /ensure-everything-is-referenced)
	EnsureEverythingIsReferenced(ctx echo.Context, requestBody RequestBody) (resp *EnsureEverythingIsReferencedResponse, err error)

	// (GET /params_with_add_props)
	ParamsWithAddProps(ctx echo.Context, params ParamsWithAddPropsParams) (code int, err error)

	// (POST /params_with_add_props)
	BodyWithAddProps(ctx echo.Context, requestBody BodyWithAddPropsJSONBody) (code int, err error)
}

type EnsureEverythingIsReferencedResponse struct {
	Code        int
	JSON200     *Ensureeverythingisreferenced200JSONResponseBodySchema
	JSONDefault *struct {
		Field SchemaObject `json:"Field" validate:"required"`
	}
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// EnsureEverythingIsReferenced converts echo context to params.
func (w *ServerInterfaceWrapper) EnsureEverythingIsReferenced(ctx echo.Context) error {
	var err error

	var requestBody RequestBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.EnsureEverythingIsReferenced(ctx, requestBody)

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

// ParamsWithAddProps converts echo context to params.
func (w *ServerInterfaceWrapper) ParamsWithAddProps(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ParamsWithAddPropsParams
	// ------------- Required query parameter "p1" -------------

	err = runtime.BindQueryParameter("simple", true, true, "p1", ctx.QueryParams(), &params.P1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p1: %s", err))
	}

	// ------------- Required query parameter "p2" -------------

	err = runtime.BindQueryParameter("form", true, true, "p2", ctx.QueryParams(), &params.P2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p2: %s", err))
	}

	if err = runtime.ValidateInput(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.ParamsWithAddProps(ctx, params)

	if err != nil {
		return err
	}

	return ctx.NoContent(response)
}

// BodyWithAddProps converts echo context to params.
func (w *ServerInterfaceWrapper) BodyWithAddProps(ctx echo.Context) error {
	var err error

	var requestBody BodyWithAddPropsJSONBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.BodyWithAddProps(ctx, requestBody)

	if err != nil {
		return err
	}

	return ctx.NoContent(response)
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

	router.GET(baseURL+"/ensure-everything-is-referenced", wrapper.EnsureEverythingIsReferenced, m...)
	router.GET(baseURL+"/params_with_add_props", wrapper.ParamsWithAddProps, m...)
	router.POST(baseURL+"/params_with_add_props", wrapper.BodyWithAddProps, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xYX2/bNhD/Kgduj2zsOO0G+C1tU6wF2hRNgD3UQUGLp4idRKokZdco/N2HoyRLtinX",
	"cYIC20tiy+Td7/797k4/WGKK0mjU3rHpD1YKKwr0aMO3K10VH+nJOX2T6BKrSq+MZlMmYHOWcabo0bcK",
	"7YpxpkWBbMrCTRJxzjhzSYaFIDGoq4JNP3tbIU9F7pCzufEZu+PMr0q66LxV+p6t17xDMHkUgskTILh4",
	"FIKLGAImcpUg4wwX9Hdu5nEMH1st1/OvmHiSkBjtUYePoixzlQjCNPrqCNiPnqrSmhKtVxgi+kZhLunD",
	"7xZTNmW/jbr4j+pLbnQT/je6SL/Fb5WyKAlyLeGOHnv87kdlLtSOyl0D1vw0v/UsblGg8y+NbIz5tHmw",
	"+s+5pJbhSqNda0z95X8S4ktwqihzhNZIMJ2yBgUJupRS0RWRf9xYUcMKnCMiP/f0K+3xHi3bU/+XcNDd",
	"hc5DYFKgy6C0Z3zHdUrGZddJuWc1Z6asFcRcsu3TIIKThq7CW4/wA16YDHuhoa5tw18bdKCNB5HnZhn3",
	"wWPtfiLTLoZNC9S8m1FkkAOhVxGrVns2PQD7w2A/fxjskIna6FVhKgcplRYsM5VkkA3l6H58tEb7M7VP",
	"av5x12tc/BQvvjhU3ccz1/F1v1Q+g1oIpMaCVEk4ZGuH70G/0q6yiAu0K58pfa+cxRQt6gTlZDx+d3P9",
	"oaVsakA3A2ycqgX+zKRDXlpzlprKni7iOYmgnnFUX6jv/K189m5zhXhOP8KIc5LgM4uPkHERZCzN6RIm",
	"IXsjYa5i0+0ro9NcJb7JmzBDcqCxDRYir4hjESV4A3OE0mKqvqOsz5KKmaaSOJtpmu2aWe9aE0veLg39",
	"Dd7Yn/VqPJOj8Jz/MjyRyfeqpzlpsHXQHIi5WSCPgZrpDSrYBxUMq5G9MYTspbCDuJ7v4/pgNnAcZGiR",
	"g8tMlUvIxAKh0oWRKlUog+/clu5XgiaC1+aecfbeVG7YIS8iiqsCrUpEHgT35b7gf/A/O0m99hortj0C",
	"GZw9gvdD6qbGFsKzaShzxgeOTo44Gu/sjaYYwW+xcYT8rPMfhgywKOS1zlefGo1UsfuOvc2U23R5UA5a",
	"gCC0hFYGB2fAZwj3BgojMd8Ku/JAfQFKQ963HGZaOHq6VHkepqU5DYRJXkmqmgw1/agcONQeUmsKSHJF",
	"n70Bh3aBtk7bVn3bg/dtNHnc+KVVHmPW76RJ72StZDdGQQPvuXrAsUMq78LwrnQauJX2Ue2wSzz2/u1t",
	"4F7lyRJ2i87DTXABZQZaV4fp/Gx8Nq5HYtSiVGzKLs7GZ7T6l8JnIR1GGHrqs66pPlPuWddW6cw9+oEc",
	"QC1D/AC/K+ddHXDhoesCkAhNkUwsCo8SlAafKTfTrsQkpEsT6tJWmpiIBbg27FVvZaA1Ani1wffWferQ",
	"9TfQ1VAj2lpSR/0NdXfhm4zHD9ryDrW904aVyNpW5iLBzOSyXapSUeV+2NjGntHO9lrvt6Ow5bsvRPhf",
	"hJRfqIzdYJQvgTKl7iLde6gQNwFzI1fN3Nxkf3zOiwQ1vEtwRLOXMuR8GLB7L7o+xzmnPTE8qAdlsdcX",
	"5TnrV2lNDl0oD43xexTr/CpUXr1PszU/Bq3u7RxhSu8otHYidWZXt+aZ9pXVoV68IZYMJ+sOTZNyDC3d",
	"XBr7z7AHJgc98KD9JkKfO16K7iV36/XdTs318rlXdye/5NiqFvq5NC6S2GEpgYYZ+plMGITS9KtUFhMf",
	"9TWnEpjpgzGlmondjZQDEcBOMdgTX6Qdv44eG+HedHTiTtq+jWhTYL2bhutfmhPr9frfAAAA//99ou74",
	"ZBcAAA==",
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
