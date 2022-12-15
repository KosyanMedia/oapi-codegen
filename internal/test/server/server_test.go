package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestDefaults(t *testing.T) {
	t.Parallel()

	// given
	argument := 43
	body := Resource{
		Name:              "test",
		Value:             1.1,
		IntFieldDefault:   nil,
		FloatFieldDefault: nil,
	}

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/resource2/%d", argument), toJsonReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, response)
	ctx.SetPath("/resource2/:inline_argument")
	ctx.SetParamNames("inline_argument")
	ctx.SetParamValues(strconv.Itoa(argument))

	m := ServerInterfaceMock{}
	wrapper := ServerInterfaceWrapper{
		Handler: &m,
	}
	m.CreateResource2Func = func(ctx echo.Context, inlineArgument int, params CreateResource2Params, requestBody CreateResource2JSONBody) (*CreateResource2Response, error) {
		assert.Equal(t, inlineArgument, argument)
		assert.Nil(t, params.InlineQueryArgument)
		assert.Equal(t, float32(5.5), *requestBody.FloatFieldDefault)
		assert.Equal(t, 5, *requestBody.IntFieldDefault)
		return &CreateResource2Response{
			Code: 200,
			JSON200: &(struct {
				Name string `json:"name" validate:"required"`
			}{
				Name: "value",
			}),
		}, nil
	}

	// when
	err := wrapper.CreateResource2(ctx)

	// then
	assert.Nil(t, err)
	assert.Equal(t, echo.MIMEApplicationJSON, req.Header.Get("Content-Type"))
	resp, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"value"}`, unrettyfy(string(resp)))
}

func TestValidationsFail(t *testing.T) {
	t.Parallel()

	intField := -10
	floatField := float32(1)
	stringField := "123456"
	patternField := "ru_wrong"
	countryField := "fake"
	enumField := CustomEnumType("non-existing")
	var body EveryTypeOptional

	t.Run("int", func(t *testing.T) {
		body = EveryTypeOptional{
			IntField: &intField,
		}
		testValidationsFail(t, body, "Key: 'EveryTypeOptional.int_field' Error:Field validation for 'int_field' failed on the 'min' tag")
	})

	t.Run("float", func(t *testing.T) {
		body = EveryTypeOptional{
			FloatField: &floatField,
		}
		testValidationsFail(t, body, "Key: 'EveryTypeOptional.float_field' Error:Field validation for 'float_field' failed on the 'min' tag")
	})

	t.Run("string", func(t *testing.T) {
		body = EveryTypeOptional{
			StringField: &stringField,
		}
		testValidationsFail(t, body, "Key: 'EveryTypeOptional.string_field' Error:Field validation for 'string_field' failed on the 'max' tag")
	})

	t.Run("pattern", func(t *testing.T) {
		body = EveryTypeOptional{
			PatternField: &patternField,
		}
		testValidationsFail(t, body, "Key: 'EveryTypeOptional.pattern_field' Error:Field validation for 'pattern_field' failed on the 'pattern' tag")
	})

	t.Run("x-validate", func(t *testing.T) {
		body = EveryTypeOptional{
			CountryField: &countryField,
		}
		testValidationsFail(t, body, "Key: 'EveryTypeOptional.country_field' Error:Field validation for 'country_field' failed on the 'iso3166_1_alpha2' tag")
	})

	t.Run("enum", func(t *testing.T) {
		body = EveryTypeOptional{
			EnumField: &enumField,
		}
		testValidationsFail(t, body, "Key: 'EveryTypeOptional.enum_field' Error:Field validation for 'enum_field' failed on the 'oneof' tag")
	})
}

func TestValidationsSuccess(t *testing.T) {
	t.Parallel()

	intField := 4
	floatField := float32(1.7)
	stringField := "12345"
	patternField := "ru_RU"
	countryField := "RU"
	enumField := CustomEnumType("first")
	var body EveryTypeOptional

	t.Run("int", func(t *testing.T) {
		body = EveryTypeOptional{
			IntField: &intField,
		}
		testValidationsSuccess(t, body)
	})

	t.Run("float", func(t *testing.T) {
		body = EveryTypeOptional{
			FloatField: &floatField,
		}
		testValidationsSuccess(t, body)
	})

	t.Run("string", func(t *testing.T) {
		body = EveryTypeOptional{
			StringField: &stringField,
		}
		testValidationsSuccess(t, body)
	})

	t.Run("pattern", func(t *testing.T) {
		body = EveryTypeOptional{
			PatternField: &patternField,
		}
		testValidationsSuccess(t, body)
	})

	t.Run("x-validate", func(t *testing.T) {
		body = EveryTypeOptional{
			CountryField: &countryField,
		}
		testValidationsSuccess(t, body)
	})

	t.Run("enum", func(t *testing.T) {
		body = EveryTypeOptional{
			EnumField: &enumField,
		}
		testValidationsSuccess(t, body)
	})
}

func TestMiddlewares(t *testing.T) {
	t.Parallel()

	// given
	routes := make([]Route, 0)
	e := EchoRouterMock{}
	e.GETFunc = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
		routes = append(routes, Route{
			Path:        path,
			Method:      "GET",
			Handler:     h,
			Middlewares: m,
		})
		return nil
	}
	e.PUTFunc = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
		routes = append(routes, Route{
			Path:        path,
			Method:      "PUT",
			Handler:     h,
			Middlewares: m,
		})
		return nil
	}
	e.POSTFunc = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
		routes = append(routes, Route{
			Path:        path,
			Method:      "POST",
			Handler:     h,
			Middlewares: m,
		})
		return nil
	}
	server := ServerInterfaceMock{}
	baseUrl := "/api/v1/test"

	// when
	RegisterHandlersWithBaseURL(&e, &server, baseUrl, adminMiddleware, premiumMiddleware, coverAllMiddleware)

	// then
	sort.Sort(ByPath(routes))
	assert.Equal(t, 11, len(routes))
	assert.Equal(t, "GET", routes[0].Method)
	assert.Equal(t, baseUrl+"/every-type-optional", routes[0].Path)
	assert.Equal(t, ptr(coverAllMiddleware), ptr(routes[0].Middlewares[0]))
	assert.Equal(t, "POST", routes[1].Method)
	assert.Equal(t, baseUrl+"/every-type-optional", routes[1].Path)
	assert.Equal(t, ptr(coverAllMiddleware), ptr(routes[1].Middlewares[0]))
	assert.Equal(t, "GET", routes[2].Method)
	assert.Equal(t, baseUrl+"/get-simple", routes[2].Path)
	assert.Equal(t, ptr(coverAllMiddleware), ptr(routes[2].Middlewares[0]))
	assert.Equal(t, "GET", routes[3].Method)
	assert.Equal(t, baseUrl+"/get-with-args", routes[3].Path)
	assert.Equal(t, ptr(premiumMiddleware), ptr(routes[3].Middlewares[0]))
	assert.Equal(t, ptr(coverAllMiddleware), ptr(routes[3].Middlewares[1]))
	assert.Equal(t, "GET", routes[4].Method)
	assert.Equal(t, baseUrl+"/get-with-references/:global_argument/:argument", routes[4].Path)
	assert.Equal(t, "GET", routes[5].Method)
	assert.Equal(t, baseUrl+"/get-with-type/:content_type", routes[5].Path)
	assert.Equal(t, "GET", routes[6].Method)
	assert.Equal(t, baseUrl+"/reserved-keyword", routes[6].Path)
	assert.Equal(t, "POST", routes[7].Method)
	assert.Equal(t, baseUrl+"/resource/:argument", routes[7].Path)
	assert.Equal(t, "POST", routes[8].Method)
	assert.Equal(t, baseUrl+"/resource2/:inline_argument", routes[8].Path)
	assert.Equal(t, "PUT", routes[9].Method)
	assert.Equal(t, baseUrl+"/resource3/:fallthrough", routes[9].Path)
	assert.Equal(t, ptr(adminMiddleware), ptr(routes[9].Middlewares[0]))
	assert.Equal(t, ptr(premiumMiddleware), ptr(routes[9].Middlewares[1]))
	assert.Equal(t, ptr(coverAllMiddleware), ptr(routes[9].Middlewares[2]))
	assert.Equal(t, "GET", routes[10].Method)
	assert.Equal(t, baseUrl+"/response-with-reference", routes[10].Path)
}

func premiumMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return f
}

func adminMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return f
}

func coverAllMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return f
}

func testValidationsSuccess(t *testing.T, body EveryTypeOptional) {
	// given
	req := httptest.NewRequest(http.MethodPost, "/every-type-optional", toJsonReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, response)
	ctx.SetPath("/every-type-optional")

	m := ServerInterfaceMock{}
	wrapper := ServerInterfaceWrapper{
		Handler: &m,
	}
	m.CreateEveryTypeOptionalFunc = func(ctx echo.Context, params CreateEveryTypeOptionalParams, requestBody CreateEveryTypeOptionalJSONBody) (int, error) {
		return 200, nil
	} // must not be called

	// when
	err := wrapper.CreateEveryTypeOptional(ctx)

	// then
	assert.Nil(t, err)
}

func testValidationsFail(t *testing.T, body EveryTypeOptional, expectedError string) {
	// given
	req := httptest.NewRequest(http.MethodPost, "/every-type-optional", toJsonReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, response)
	ctx.SetPath("/every-type-optional")

	m := ServerInterfaceMock{}
	wrapper := ServerInterfaceWrapper{
		Handler: &m,
	}
	m.CreateEveryTypeOptionalFunc = nil // must not be called

	// when
	err := wrapper.CreateEveryTypeOptional(ctx)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, 400, httpErr.Code)
	assert.Equal(t, expectedError, httpErr.Message)
}

func toJsonReader(body interface{}) io.Reader {
	if body == nil {
		return nil
	}
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(body)
	if err != nil {
		panic(err)
	}
	return &b
}

func unrettyfy(val string) string {
	return strings.ReplaceAll(val, "\n", "")
}

func ptr(i interface{}) uintptr {
	return reflect.ValueOf(i).Pointer()
}

type Route struct {
	Path        string
	Method      string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

type ByPath []Route

func (a ByPath) Len() int           { return len(a) }
func (a ByPath) Less(i, j int) bool { return a[i].Path < a[j].Path }
func (a ByPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
