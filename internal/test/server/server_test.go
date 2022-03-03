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
	"strconv"
	"strings"
	"testing"
)

func TestDefaults(t *testing.T) {
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
	intField := -10
	floatField := float32(1)
	stringField := "123456"
	patternField := "ru_wrong"
	var body EveryTypeOptional

	t.Run("int", func(t *testing.T) {
		body = EveryTypeOptional{
			IntField: &intField,
		}
		testValidationsFail(t, body, "Key: 'CreateEveryTypeOptionalJSONBody.int_field' Error:Field validation for 'int_field' failed on the 'min' tag")
	})

	t.Run("float", func(t *testing.T) {
		body = EveryTypeOptional{
			FloatField: &floatField,
		}
		testValidationsFail(t, body, "Key: 'CreateEveryTypeOptionalJSONBody.float_field' Error:Field validation for 'float_field' failed on the 'min' tag")
	})

	t.Run("string", func(t *testing.T) {
		body = EveryTypeOptional{
			StringField: &stringField,
		}
		testValidationsFail(t, body, "Key: 'CreateEveryTypeOptionalJSONBody.string_field' Error:Field validation for 'string_field' failed on the 'max' tag")
	})

	t.Run("string", func(t *testing.T) {
		body = EveryTypeOptional{
			PatternField: &patternField,
		}
		testValidationsFail(t, body, "Key: 'CreateEveryTypeOptionalJSONBody.pattern_field' Error:Field validation for 'pattern_field' failed on the 'pattern' tag")
	})
}

func TestValidationsSuccess(t *testing.T) {
	intField := 4
	floatField := float32(1.7)
	stringField := "12345"
	patternField := "ru_RU"
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

	t.Run("string", func(t *testing.T) {
		body = EveryTypeOptional{
			PatternField: &patternField,
		}
		testValidationsSuccess(t, body)
	})
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
	m.CreateEveryTypeOptionalFunc = func(ctx echo.Context, requestBody CreateEveryTypeOptionalJSONBody) (int, error) {
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
