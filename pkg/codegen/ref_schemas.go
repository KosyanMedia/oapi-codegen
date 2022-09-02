package codegen

import (
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func refSchemas(swagger *openapi3.T) {
	var err error
	if swagger.Components.Schemas == nil {
		swagger.Components.Schemas = make(map[string]*openapi3.SchemaRef)
	}
	schemas := swagger.Components.Schemas

	for _, requestPath := range SortedPathsKeys(swagger.Paths) {
		path := swagger.Paths[requestPath]
		ops := path.Operations()
		for _, opName := range SortedOperationsKeys(ops) {
			operation := ops[opName]
			opId := operation.OperationID
			if opId == "" {
				opId, err = generateDefaultOperationID(opName, requestPath)
				if err != nil {
					continue
				}
			} else {
				opId = ToCamelCase(opId)
			}
			opId = cases.Title(language.Und).String(opId)

			if schemaName, schema := refRequestBody(opId, operation.RequestBody); schemaName != "" && schema != nil {
				schemas[schemaName] = schema
			}

			for _, respName := range SortedResponsesKeys(operation.Responses) {
				resp := operation.Responses[respName]
				if schemaName, schema := refResponseBody(opId, respName, resp); schemaName != "" && schema != nil {
					schemas[schemaName] = schema
				}
			}

		}
	}
}

func refRequestBody(operationId string, reqBody *openapi3.RequestBodyRef) (string, *openapi3.SchemaRef) {
	if reqBody == nil || reqBody.Ref != "" || reqBody.Value == nil {
		return "", nil
	}

	// currently only json is supported
	contentTypeJSON := reqBody.Value.Content["application/json"]
	if contentTypeJSON == nil || contentTypeJSON.Schema == nil || contentTypeJSON.Schema.Ref != "" {
		return "", nil
	}

	targetSchema := itemSchema(contentTypeJSON.Schema)
	if targetSchema == nil {
		return "", nil
	}

	schemaName := operationId + "JSONRequestBodySchema"
	schemaCopy := *targetSchema
	targetSchema.Ref = globalCtxRef(schemaName)

	return schemaName, &schemaCopy
}

func refResponseBody(operationId string, responseName string, response *openapi3.ResponseRef) (string, *openapi3.SchemaRef) {
	if response == nil || response.Ref != "" || response.Value == nil {
		return "", nil
	}

	// currently only json is supported
	contentTypeJSON := response.Value.Content["application/json"]
	if contentTypeJSON == nil || contentTypeJSON.Schema == nil || contentTypeJSON.Schema.Ref != "" {
		return "", nil
	}

	targetSchema := itemSchema(contentTypeJSON.Schema)
	if targetSchema == nil {
		return "", nil
	}

	schemaName := operationId + responseName + "JSONResponseBodySchema"
	schemaCopy := *targetSchema
	targetSchema.Ref = globalCtxRef(schemaName)

	return schemaName, &schemaCopy
}

func itemSchema(ref *openapi3.SchemaRef) *openapi3.SchemaRef {
	if ref == nil || ref.Ref != "" || ref.Value == nil {
		return nil
	}
	if ref.Value.Type == "array" {
		return itemSchema(ref.Value.Items)
	}
	if ref.Value.Type == "object" {
		return ref
	}
	return nil
}
