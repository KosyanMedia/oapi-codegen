package codegen

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-test/deep"
)

func flatSchemas(schemas openapi3.Schemas, aliases aliases) openapi3.Schemas {
	target := make(openapi3.Schemas)
	for name, schema := range schemas {
		target[name] = schema
	}
	for _, name := range SortedSchemaKeys(schemas) {
		schema := schemas[name]
		if !isEmbeddedStruct(schema) {
			continue
		}

		switch schema.Value.Type {
		case "object":
			properties := schema.Value.Properties
			flatSchemasNext(properties, aliases, target, name)
		case "array":
			flatArrayNext(schema, aliases, target, name)
		}
	}
	return target
}

func flatSchemasNext(properties openapi3.Schemas, aliases aliases, target openapi3.Schemas, baseNameParts ...string) {
	for _, propName := range SortedSchemaKeys(properties) {
		propSchema := properties[propName]
		if !isEmbeddedStruct(propSchema) && !isEmbeddedEnum(propSchema) {
			continue
		}

		switch propSchema.Value.Type {
		default:
			if len(propSchema.Value.Properties) > 0 || isEmbeddedEnum(propSchema) {
				newBaseNamePart, newPropName := aliases.pathToTypeName(append(baseNameParts, propName)...)

				// move to global context
				putCarefully(target, newPropName, propSchema)

				// make this property ref
				propSchemaCopy := *propSchema
				propSchemaCopy.Ref = globalCtxRef(newPropName)
				properties[propName] = &propSchemaCopy

				// go deeper
				flatSchemasNext(propSchema.Value.Properties, aliases, target, newBaseNamePart...)
			}

			additionalProps := propSchema.Value.AdditionalProperties
			if additionalProps != nil && additionalProps.Value != nil && additionalProps.Ref == "" &&
				(isEmbeddedStruct(additionalProps) || isEmbeddedEnum(additionalProps)) {
				newBaseNamePart, newPropName := aliases.pathToTypeName(append(baseNameParts, propName, "props")...)

				// move to global context
				putCarefully(target, newPropName, additionalProps)

				// make this property ref
				additionalPropsCopy := *additionalProps
				additionalPropsCopy.Ref = globalCtxRef(newPropName)
				propSchema.Value.AdditionalProperties = &additionalPropsCopy

				// go deeper
				flatSchemasNext(additionalProps.Value.Properties, aliases, target, newBaseNamePart...)
			}
		case "array":
			flatArrayNext(propSchema, aliases, target, append(baseNameParts, propName)...)
		}
	}
}

func flatArrayNext(schema *openapi3.SchemaRef, aliases aliases, target openapi3.Schemas, baseNameParts ...string) {
	items := schema.Value.Items
	if !isEmbeddedStruct(items) && !isEmbeddedEnum(items) {
		return
	}

	if items.Value.Type == "array" {
		// no need to generate redundant slice structs
		flatArrayNext(items, aliases, target, baseNameParts...)
		return
	}

	newBaseNamePart, newPropName := aliases.pathToTypeName(append(baseNameParts, "item")...)

	// move to global context
	putCarefully(target, newPropName, items)

	// make this property ref
	itemsCopy := *items
	itemsCopy.Ref = globalCtxRef(newPropName)
	schema.Value.Items = &itemsCopy

	// go deeper
	flatSchemasNext(items.Value.Properties, aliases, target, newBaseNamePart...)
}

func isEmbeddedStruct(schema *openapi3.SchemaRef) bool {
	return schema.Ref == "" && schema.Value != nil &&
		(schema.Value.Type == "object" ||
			schema.Value.Type == "array")
}

func isEmbeddedEnum(schema *openapi3.SchemaRef) bool {
	return schema.Ref == "" && schema.Value != nil && len(schema.Value.Enum) > 0
}

func globalCtxRef(modelName string) string {
	return fmt.Sprintf("#/components/schemas/%s", modelName)
}

func putCarefully(target openapi3.Schemas, key string, value *openapi3.SchemaRef) {
	if existing, found := target[key]; found {
		if diff := deep.Equal(existing, value); diff != nil {
			fmt.Printf(
				"WARN: Conflict found in '%s' deep struct, found 2 different structs with the same name. Diff: %v\n",
				key, diff)
		}
		*value = *existing
		return
	}
	target[key] = value
}

type aliases map[string]string

func (a aliases) pathToTypeName(path ...string) (newPath []string, name string) {
	name = SchemaNameToTypeName(PathToTypeName(path...))
	if newName := a[name]; newName != "" {
		return []string{newName}, newName
	}
	return path, name
}
