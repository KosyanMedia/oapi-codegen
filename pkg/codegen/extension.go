package codegen

import (
	"encoding/json"
	"fmt"
)

const (
	extPropGoType             = "x-go-type"
	extPropMiddlewares        = "x-middlewares"
	extPropValidate           = "x-validate"
	extPropGenericErrResponse = "x-generic-err-response"
	extPropOmitEmpty          = "x-omitempty"
	extPropExtraTags          = "x-oapi-codegen-extra-tags"
)

func extParseString(extPropValue interface{}) (string, error) {
	raw, ok := extPropValue.(json.RawMessage)
	if !ok {
		return "", fmt.Errorf("failed to convert type: %T", extPropValue)
	}
	var name string
	if err := json.Unmarshal(raw, &name); err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return name, nil
}

func extParseStringSlice(extPropValue interface{}) ([]string, error) {
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

func extParseBool(extPropValue interface{}) (bool, error) {
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
