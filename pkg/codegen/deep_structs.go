package codegen

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-test/deep"
)

func flatSchemas(schemas openapi3.Schemas, aliases aliases) openapi3.Schemas {
	target := make(openapi3.Schemas)
	for _, name := range SortedSchemaKeys(schemas) {
		schema := schemas[name]
		target[name] = schema
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
		if !isEmbeddedStruct(propSchema) {
			continue
		}

		switch propSchema.Value.Type {
		case "object":
			newBaseNamePart, newPropName := aliases.pathToTypeName(append(baseNameParts, propName)...)

			// go deeper
			flatSchemasNext(propSchema.Value.Properties, aliases, target, newBaseNamePart...)

			// move to global context
			putCarefully(target, newPropName, propSchema)

			// make this property ref
			propSchemaCopy := *propSchema
			propSchemaCopy.Ref = globalCtxRef(newPropName)
			properties[propName] = &propSchemaCopy
		case "array":
			flatArrayNext(propSchema, aliases, target, append(baseNameParts, propName)...)
		}
	}
}

func flatArrayNext(schema *openapi3.SchemaRef, aliases aliases, target openapi3.Schemas, baseNameParts ...string) {
	items := schema.Value.Items
	if !isEmbeddedStruct(items) {
		return
	}

	if items.Value.Type == "array" {
		// no need to generate redundant slice structs
		flatArrayNext(items, aliases, target, baseNameParts...)
		return
	}

	newBaseNamePart, newPropName := aliases.pathToTypeName(append(baseNameParts, "item")...)

	// go deeper
	flatSchemasNext(items.Value.Properties, aliases, target, newBaseNamePart...)

	// move to global context
	putCarefully(target, newPropName, items)

	// make this property ref
	itemsCopy := *items
	itemsCopy.Ref = globalCtxRef(newPropName)
	schema.Value.Items = &itemsCopy
}

func isEmbeddedStruct(schema *openapi3.SchemaRef) bool {
	return schema.Ref == "" && schema.Value != nil &&
		(schema.Value.Type == "object" && len(schema.Value.Properties) > 0 ||
			schema.Value.Type == "array")
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
