// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/KosyanMedia/oapi-codegen/v2 version (devel) DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KosyanMedia/oapi-codegen/v2/pkg/runtime"
	openapi_types "github.com/KosyanMedia/oapi-codegen/v2/pkg/types"
	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
)

// EveryTypeOptional defines model for EveryTypeOptional.
type EveryTypeOptional struct {
	ArrayInlineField     []int               `json:"array_inline_field,omitempty"`
	ArrayReferencedField []SomeObject        `json:"array_referenced_field,omitempty"`
	BoolField            *bool               `json:"bool_field,omitempty"`
	ByteField            []byte              `json:"byte_field,omitempty"`
	CountryField         *string             `json:"country_field,omitempty" validate:"omitempty,iso3166_1_alpha2"`
	DateField            *openapi_types.Date `json:"date_field,omitempty"`
	DateTimeField        *time.Time          `json:"date_time_field,omitempty"`
	DoubleField          *float64            `json:"double_field,omitempty"`
	EnumField            *CustomEnumType     `json:"enum_field,omitempty" validate:"omitempty,oneof=first second"`
	FloatField           *float32            `json:"float_field,omitempty" validate:"omitempty,min=1.5,max=5.5"`
	InlineObjectField    *struct {
		Name   string `json:"name" validate:"required"`
		Number int    `json:"number" validate:"required"`
	} `json:"inline_object_field,omitempty"`
	Int32Field      *int32      `json:"int32_field,omitempty"`
	Int64Field      *int64      `json:"int64_field,omitempty"`
	IntField        *int        `json:"int_field,omitempty" validate:"omitempty,min=1,max=5"`
	NumberField     *float32    `json:"number_field,omitempty"`
	PatternField    *string     `json:"pattern_field,omitempty" validate:"omitempty,pattern=KFtcd117Mn0pXyhbXHddezJ9KQ=="`
	ReferencedField *SomeObject `json:"referenced_field,omitempty"`
	StringField     *string     `json:"string_field,omitempty" validate:"omitempty,min=1,max=5"`
}

// EveryTypeRequired defines model for EveryTypeRequired.
type EveryTypeRequired struct {
	ArrayInlineField     []int                `json:"array_inline_field" validate:"required"`
	ArrayReferencedField []SomeObject         `json:"array_referenced_field" validate:"required"`
	BoolField            bool                 `json:"bool_field" validate:"required"`
	ByteField            []byte               `json:"byte_field" validate:"required"`
	DateField            openapi_types.Date   `json:"date_field" validate:"required"`
	DateTimeField        time.Time            `json:"date_time_field" validate:"required"`
	DoubleField          float64              `json:"double_field" validate:"required"`
	EmailField           *openapi_types.Email `json:"email_field,omitempty"`
	FloatField           float32              `json:"float_field" validate:"required"`
	InlineObjectField    struct {
		Name   string `json:"name" validate:"required"`
		Number int    `json:"number" validate:"required"`
	} `json:"inline_object_field" validate:"required"`
	Int32Field      int32      `json:"int32_field" validate:"required"`
	Int64Field      int64      `json:"int64_field" validate:"required"`
	IntField        int        `json:"int_field" validate:"required"`
	NumberField     float32    `json:"number_field" validate:"required"`
	ReferencedField SomeObject `json:"referenced_field" validate:"required"`
	StringField     string     `json:"string_field" validate:"required"`
}

// ReservedKeyword defines model for ReservedKeyword.
type ReservedKeyword struct {
	Channel *string `json:"channel,omitempty"`
}

// Resource defines model for Resource.
type Resource struct {
	FloatFieldDefault *float32 `default:"5.5" json:"float_field_default,omitempty"`
	IntFieldDefault   *int     `default:"5" json:"int_field_default,omitempty"`
	Name              string   `json:"name" validate:"required"`
	Value             float32  `json:"value" validate:"required"`
}

// SomeObject defines model for some_object.
type SomeObject struct {
	Name string `json:"name" validate:"required"`
}

// Argument defines model for argument.
type Argument = string

// Error defines model for Error.
type Error struct {
	Message string `json:"message" validate:"required"`
}

// ResponseWithReference defines model for ResponseWithReference.
type ResponseWithReference = SomeObject

// SimpleResponse defines model for SimpleResponse.
type SimpleResponse struct {
	Name string `json:"name" validate:"required"`
}

// CreateEveryTypeOptionalJSONBody defines parameters for CreateEveryTypeOptional.
type CreateEveryTypeOptionalJSONBody = EveryTypeOptional

// CreateEveryTypeOptionalParams defines parameters for CreateEveryTypeOptional.
type CreateEveryTypeOptionalParams struct {
	EnumType *CustomEnumType `form:"enum_type,omitempty" json:"enum_type,omitempty" validate:"omitempty,oneof=first second"`
}

// GetWithArgsParams defines parameters for GetWithArgs.
type GetWithArgsParams struct {
	// An optional query argument
	OptionalArgument *int64 `form:"optional_argument,omitempty" json:"optional_argument,omitempty"`

	// An optional query argument
	RequiredArgument int64 `form:"required_argument" json:"required_argument" validate:"required"`

	// An optional query argument
	HeaderArgument *int32 `json:"header_argument,omitempty"`
}

// GetWithContentTypeParamsContentType defines parameters for GetWithContentType.
type GetWithContentTypeParamsContentType string

// CreateResourceJSONBody defines parameters for CreateResource.
type CreateResourceJSONBody = EveryTypeRequired

// CreateResource2JSONBody defines parameters for CreateResource2.
type CreateResource2JSONBody = Resource

// CreateResource2Params defines parameters for CreateResource2.
type CreateResource2Params struct {
	// Some query argument
	InlineQueryArgument *int `form:"inline_query_argument,omitempty" json:"inline_query_argument,omitempty"`
}

// UpdateResource3JSONBody defines parameters for UpdateResource3.
type UpdateResource3JSONBody struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateEveryTypeOptionalJSONRequestBody defines body for CreateEveryTypeOptional for application/json ContentType.
type CreateEveryTypeOptionalJSONRequestBody = CreateEveryTypeOptionalJSONBody

// CreateResourceJSONRequestBody defines body for CreateResource for application/json ContentType.
type CreateResourceJSONRequestBody = CreateResourceJSONBody

// CreateResource2JSONRequestBody defines body for CreateResource2 for application/json ContentType.
type CreateResource2JSONRequestBody = CreateResource2JSONBody

// UpdateResource3JSONRequestBody defines body for UpdateResource3 for application/json ContentType.
type UpdateResource3JSONRequestBody UpdateResource3JSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get every type optional
	// (GET /every-type-optional)
	GetEveryTypeOptional(ctx echo.Context) (resp *GetEveryTypeOptionalResponse, err error)
	// create every type optional
	// (POST /every-type-optional)
	CreateEveryTypeOptional(ctx echo.Context, params CreateEveryTypeOptionalParams, requestBody CreateEveryTypeOptionalJSONBody) (code int, err error)
	// Get resource via simple path
	// (GET /get-simple)
	GetSimple(ctx echo.Context) (resp *GetSimpleResponse, err error)
	// Getter with referenced parameter and referenced response
	// (GET /get-with-args)
	GetWithArgs(ctx echo.Context, params GetWithArgsParams) (resp *GetWithArgsResponse, err error)
	// Getter with referenced parameter and referenced response
	// (GET /get-with-references/{global_argument}/{argument})
	GetWithReferences(ctx echo.Context, globalArgument int64, argument Argument) (resp *GetWithReferencesResponse, err error)
	// Get an object by ID
	// (GET /get-with-type/{content_type})
	GetWithContentType(ctx echo.Context, contentType GetWithContentTypeParamsContentType) (resp *GetWithContentTypeResponse, err error)
	// get with reserved keyword
	// (GET /reserved-keyword)
	GetReservedKeyword(ctx echo.Context) (resp *GetReservedKeywordResponse, err error)
	// Create a resource
	// (POST /resource/{argument})
	CreateResource(ctx echo.Context, argument Argument, requestBody CreateResourceJSONBody) (resp *CreateResourceResponse, err error)
	// Create a resource with inline parameter
	// (POST /resource2/{inline_argument})
	CreateResource2(ctx echo.Context, inlineArgument int, params CreateResource2Params, requestBody CreateResource2JSONBody) (resp *CreateResource2Response, err error)
	// Update a resource with inline body. The parameter name is a reserved
	// keyword, so make sure that gets prefixed to avoid syntax errors
	// (PUT /resource3/{fallthrough})
	UpdateResource3(ctx echo.Context, pFallthrough int, requestBody UpdateResource3JSONBody) (code int, err error)
	// get response with reference
	// (GET /response-with-reference)
	GetResponseWithReference(ctx echo.Context) (resp *GetResponseWithReferenceResponse, err error)

	Error(err error) (status int, resp Error)
}

type GetEveryTypeOptionalResponse struct {
	Code    int
	JSON200 *EveryTypeOptional
}

type GetSimpleResponse struct {
	Code    int
	JSON200 *SomeObject
}

type GetWithArgsResponse struct {
	Code    int
	JSON200 *struct {
		Name string `json:"name" validate:"required"`
	}
}

type GetWithReferencesResponse struct {
	Code    int
	JSON200 *struct {
		Name string `json:"name" validate:"required"`
	}
}

type GetWithContentTypeResponse struct {
	Code    int
	JSON200 *SomeObject
}

type GetReservedKeywordResponse struct {
	Code    int
	JSON200 *ReservedKeyword
}

type CreateResourceResponse struct {
	Code    int
	JSON200 *struct {
		Name string `json:"name" validate:"required"`
	}
}

type CreateResource2Response struct {
	Code    int
	JSON200 *struct {
		Name string `json:"name" validate:"required"`
	}
}

type GetResponseWithReferenceResponse struct {
	Code    int
	JSON200 *SomeObject
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetEveryTypeOptional converts echo context to params.
func (w *ServerInterfaceWrapper) GetEveryTypeOptional(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetEveryTypeOptional(ctx)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// CreateEveryTypeOptional converts echo context to params.
func (w *ServerInterfaceWrapper) CreateEveryTypeOptional(ctx echo.Context) error {
	var err error

	var requestBody CreateEveryTypeOptionalJSONBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateEveryTypeOptionalParams
	// ------------- Optional query parameter "enum_type" -------------

	err = runtime.BindQueryParameter("form", false, false, "enum_type", ctx.QueryParams(), &params.EnumType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter enum_type: %s", err))
	}

	if err = runtime.ValidateInput(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.CreateEveryTypeOptional(ctx, params, requestBody)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	return ctx.NoContent(response)
}

// GetSimple converts echo context to params.
func (w *ServerInterfaceWrapper) GetSimple(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetSimple(ctx)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// GetWithArgs converts echo context to params.
func (w *ServerInterfaceWrapper) GetWithArgs(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetWithArgsParams
	// ------------- Optional query parameter "optional_argument" -------------

	err = runtime.BindQueryParameter("form", false, false, "optional_argument", ctx.QueryParams(), &params.OptionalArgument)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter optional_argument: %s", err))
	}

	// ------------- Required query parameter "required_argument" -------------

	err = runtime.BindQueryParameter("form", false, true, "required_argument", ctx.QueryParams(), &params.RequiredArgument)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter required_argument: %s", err))
	}

	headers := ctx.Request().Header
	// ------------- Optional header parameter "header_argument" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header_argument")]; found {
		var HeaderArgument int32
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for header_argument, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "header_argument", runtime.ParamLocationHeader, valueList[0], &HeaderArgument)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter header_argument: %s", err))
		}

		params.HeaderArgument = &HeaderArgument
	}

	if err = runtime.ValidateInput(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetWithArgs(ctx, params)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// GetWithReferences converts echo context to params.
func (w *ServerInterfaceWrapper) GetWithReferences(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "global_argument" -------------
	var globalArgument int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "global_argument", runtime.ParamLocationPath, ctx.Param("global_argument"), &globalArgument)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter global_argument: %s", err))
	}

	// ------------- Path parameter "argument" -------------
	var argument Argument

	err = runtime.BindStyledParameterWithLocation("simple", false, "argument", runtime.ParamLocationPath, ctx.Param("argument"), &argument)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter argument: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetWithReferences(ctx, globalArgument, argument)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// GetWithContentType converts echo context to params.
func (w *ServerInterfaceWrapper) GetWithContentType(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "content_type" -------------
	var contentType GetWithContentTypeParamsContentType

	err = runtime.BindStyledParameterWithLocation("simple", false, "content_type", runtime.ParamLocationPath, ctx.Param("content_type"), &contentType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter content_type: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetWithContentType(ctx, contentType)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// GetReservedKeyword converts echo context to params.
func (w *ServerInterfaceWrapper) GetReservedKeyword(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetReservedKeyword(ctx)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// CreateResource converts echo context to params.
func (w *ServerInterfaceWrapper) CreateResource(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "argument" -------------
	var argument Argument

	err = runtime.BindStyledParameterWithLocation("simple", false, "argument", runtime.ParamLocationPath, ctx.Param("argument"), &argument)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter argument: %s", err))
	}

	var requestBody CreateResourceJSONBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.CreateResource(ctx, argument, requestBody)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// CreateResource2 converts echo context to params.
func (w *ServerInterfaceWrapper) CreateResource2(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "inline_argument" -------------
	var inlineArgument int

	err = runtime.BindStyledParameterWithLocation("simple", false, "inline_argument", runtime.ParamLocationPath, ctx.Param("inline_argument"), &inlineArgument)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inline_argument: %s", err))
	}

	var requestBody CreateResource2JSONBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}
	if err = defaults.Set(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to set defaults to request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateResource2Params
	// ------------- Optional query parameter "inline_query_argument" -------------

	err = runtime.BindQueryParameter("form", false, false, "inline_query_argument", ctx.QueryParams(), &params.InlineQueryArgument)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inline_query_argument: %s", err))
	}

	if err = runtime.ValidateInput(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.CreateResource2(ctx, inlineArgument, params, requestBody)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	if response.JSON200 != nil {
		if response.Code == 0 {
			response.Code = 200
		}
		return ctx.JSON(response.Code, response.JSON200)
	}
	return ctx.NoContent(response.Code)
}

// UpdateResource3 converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateResource3(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "fallthrough" -------------
	var pFallthrough int

	err = runtime.BindStyledParameterWithLocation("simple", false, "fallthrough", runtime.ParamLocationPath, ctx.Param("fallthrough"), &pFallthrough)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fallthrough: %s", err))
	}

	var requestBody UpdateResource3JSONBody
	err = ctx.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %s", err))
	}

	if err = runtime.ValidateInput(requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.UpdateResource3(ctx, pFallthrough, requestBody)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
	}

	return ctx.NoContent(response)
}

// GetResponseWithReference converts echo context to params.
func (w *ServerInterfaceWrapper) GetResponseWithReference(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	response, err := w.Handler.GetResponseWithReference(ctx)

	if err != nil {
		code, errResp := w.Handler.Error(err)
		return ctx.JSON(code, errResp)
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
func RegisterHandlers(router EchoRouter, si ServerInterface, admin echo.MiddlewareFunc, premium echo.MiddlewareFunc, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", admin, premium, m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, admin echo.MiddlewareFunc, premium echo.MiddlewareFunc, m ...echo.MiddlewareFunc) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/every-type-optional", wrapper.GetEveryTypeOptional, m...)
	router.POST(baseURL+"/every-type-optional", wrapper.CreateEveryTypeOptional, m...)
	router.GET(baseURL+"/get-simple", wrapper.GetSimple, m...)
	router.GET(baseURL+"/get-with-args", wrapper.GetWithArgs, append([]echo.MiddlewareFunc{premium}, m...)...)
	router.GET(baseURL+"/get-with-references/:global_argument/:argument", wrapper.GetWithReferences, m...)
	router.GET(baseURL+"/get-with-type/:content_type", wrapper.GetWithContentType, m...)
	router.GET(baseURL+"/reserved-keyword", wrapper.GetReservedKeyword, m...)
	router.POST(baseURL+"/resource/:argument", wrapper.CreateResource, m...)
	router.POST(baseURL+"/resource2/:inline_argument", wrapper.CreateResource2, m...)
	router.PUT(baseURL+"/resource3/:fallthrough", wrapper.UpdateResource3, append([]echo.MiddlewareFunc{admin, premium}, m...)...)
	router.GET(baseURL+"/response-with-reference", wrapper.GetResponseWithReference, append([]echo.MiddlewareFunc{admin}, m...)...)

}
