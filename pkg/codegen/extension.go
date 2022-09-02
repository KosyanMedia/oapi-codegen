package codegen

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
)

const (
	extPropGoType             = "x-go-type"
	extPropMiddlewares        = "x-middlewares"
	extPropValidate           = "x-validate"
	extPropGenericErrResponse = "x-generic-err-response"
	extPropOmitEmpty          = "x-omitempty"
	extPropExtraTags          = "x-oapi-codegen-extra-tags"
	extGoFieldName            = "x-go-name"
	extComments               = "x-go-comments"
	extPatternProperties      = "x-patternProperties"
)

func extString(extPropValue interface{}) (string, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return "", fmt.Errorf("failed to convert type: %T", extPropValue)
	}
	var str string
	if err := json.Unmarshal(raw, &str); err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return str, nil
}

func extStringSlice(extPropValue interface{}) ([]string, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return nil, fmt.Errorf("failed to convert type: %T", extPropValue)
	}
	var slice []string
	if err := json.Unmarshal(raw, &slice); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return slice, nil
}

func extBool(extPropValue interface{}) (bool, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return false, fmt.Errorf("failed to convert type: %T", extPropValue)
	}

	var boolValue bool
	if err := json.Unmarshal(raw, &boolValue); err != nil {
		return false, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return boolValue, nil
}

func extExtraTags(extPropValue interface{}) (map[string]string, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return nil, fmt.Errorf("failed to convert type: %T", extPropValue)
	}
	var tags map[string]string
	if err := json.Unmarshal(raw, &tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	return tags, nil
}

func extPatternPropertiesValue(extPropValue interface{}) (*openapi3.Schema, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return nil, fmt.Errorf("failed to convert type: %T", extPropValue)
	}

	var props map[string]*openapi3.Schema
	if err := json.Unmarshal(raw, &props); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	if len(props) != 1 {
		return nil, fmt.Errorf("x-patternProperties supports only one key:value pair")
	}

	for _, schema := range props {
		return schema, nil
	}
	return nil, nil
}
