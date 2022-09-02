package codegen

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
)

// analyses x-pattern-properties and propagates their value to additionalProperties
func processPatternProperties(swagger *openapi3.T) {
	for _, pathKey := range SortedPathsKeys(swagger.Paths) {
		path := swagger.Paths[pathKey]
		ops := path.Operations()

		for _, opKey := range SortedOperationsKeys(ops) {
			op := ops[opKey]
			processRequestBody(op.RequestBody)
			processParameters(op.Parameters)
			processResponses(op.Responses)
		}
	}

	for _, schemaKey := range SortedSchemaKeys(swagger.Components.Schemas) {
		schema := swagger.Components.Schemas[schemaKey]
		processPatternPropertiesDeeper(schema)
	}

	processRequestBodies(swagger.Components.RequestBodies)
	processParametersMap(swagger.Components.Parameters)
	processResponses(swagger.Components.Responses)
}

func processParameters(params openapi3.Parameters) {
	for _, param := range params {
		processParameter(param)
	}
}

func processParametersMap(params openapi3.ParametersMap) {
	for _, schemaKey := range SortedParameterKeys(params) {
		schema := params[schemaKey]
		processParameter(schema)
	}
}

func processParameter(param *openapi3.ParameterRef) {
	if param == nil || param.Ref != "" || param.Value == nil {
		return
	}
	processPatternPropertiesDeeper(param.Value.Schema)
}

func processRequestBodies(requestBodies openapi3.RequestBodies) {
	for _, schemaKey := range SortedRequestBodyKeys(requestBodies) {
		schema := requestBodies[schemaKey]
		processRequestBody(schema)
	}
}

func processRequestBody(schema *openapi3.RequestBodyRef) {
	if schema == nil {
		return
	}
	content := schema.Value.Content
	for _, mediaTypeKey := range SortedMediaTypeKeys(content) {
		mediaType := content[mediaTypeKey]
		processPatternPropertiesDeeper(mediaType.Schema)
	}
}

func processResponses(responses openapi3.Responses) {
	for _, schemaKey := range SortedResponsesKeys(responses) {
		schema := responses[schemaKey]
		content := schema.Value.Content
		for _, mediaTypeKey := range SortedMediaTypeKeys(content) {
			mediaType := content[mediaTypeKey]
			processPatternPropertiesDeeper(mediaType.Schema)
		}
	}
}

func processPatternPropertiesDeeper(ref *openapi3.SchemaRef) {
	if ref == nil || ref.Ref != "" || ref.Value == nil {
		return
	}
	if ref.Value.Type == "array" {
		processPatternPropertiesDeeper(ref.Value.Items)
	}
	if ref.Value.Type == "object" {
		extProp, found := ref.Value.Extensions[extPatternProperties]
		if found {
			schema, err := extPatternPropertiesValue(extProp)
			if err != nil {
				fmt.Printf(
					"failed to parse %q property: %v", extPatternProperties, err)
			} else {
				ref.Value.AdditionalProperties = &openapi3.SchemaRef{Value: schema}
			}
		}

		for _, propName := range SortedSchemaKeys(ref.Value.Properties) {
			propValue := ref.Value.Properties[propName]
			processPatternPropertiesDeeper(propValue)
		}
	}
}
