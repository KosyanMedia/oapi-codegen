package codegen

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/KosyanMedia/oapi-codegen/v2/pkg/util"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// This describes a Schema, a type definition.
type Schema struct {
	GoType  string // The Go type needed to represent the schema
	RefType string // If the type has a type name, this is set

	ArrayType *Schema // The schema of array element

	EnumValues map[string]string // Enum values

	Properties               []Property       // For an object, the fields with names
	HasAdditionalProperties  bool             // Whether we support additional properties
	AdditionalPropertiesType *Schema          // And if we do, their type
	AdditionalTypes          []TypeDefinition // We may need to generate auxiliary helper types, stored here

	SkipOptionalPointer bool // Some types don't need a * in front when they're optional

	Description string // The description of the element

	// If this is set, the schema will declare a type via alias, eg,
	// `type Foo = bool`. If this is not set, we will define this type via
	// type definition `type Foo bool`
	DefineViaAlias bool

	// The original OpenAPIv3 Schema.
	OAPISchema *openapi3.Schema
}

func (s Schema) IsRef() bool {
	return s.RefType != ""
}

func (s Schema) TypeDecl() string {
	if s.IsRef() {
		return s.RefType
	}
	return s.GoType
}

func (s Schema) Comments() string {
	if s.OAPISchema == nil {
		return ""
	}
	commentsRaw, found := s.OAPISchema.Extensions[extComments]
	if !found {
		return ""
	}
	comments, err := extStringSlice(commentsRaw)
	if err != nil {
		comment, err := extString(commentsRaw)
		if err != nil {
			return ""
		}
		comments = []string{comment}
	}
	for i, comment := range comments {
		comments[i] = "//" + comment
	}
	return "\n" + strings.Join(comments, "\n")
}

// AddProperty adds a new property to the current Schema, and returns an error
// if it collides. Two identical fields will not collide, but two properties by
// the same name, but different definition, will collide. It's safe to merge the
// fields of two schemas with overalapping properties if those properties are
// identical.
func (s *Schema) AddProperty(p Property) error {
	// Scan all existing properties for a conflict
	for _, e := range s.Properties {
		if e.JsonFieldName == p.JsonFieldName && !PropertiesEqual(e, p) {
			return errors.New(fmt.Sprintf("property '%s' already exists with a different type", e.JsonFieldName))
		}
	}
	s.Properties = append(s.Properties, p)
	return nil
}

func (s Schema) GetAdditionalTypeDefs() []TypeDefinition {
	var result []TypeDefinition
	for _, p := range s.Properties {
		result = append(result, p.Schema.GetAdditionalTypeDefs()...)
	}
	result = append(result, s.AdditionalTypes...)
	return result
}

func (s Schema) IsAdditionalPropertiesOnly() bool {
	return s.HasAdditionalProperties && len(s.Properties) == 0
}

func (s Schema) HasCustomMarshalling() bool {
	return s.HasAdditionalProperties && len(s.Properties) != 0
}

func (s Schema) HasDefaults() bool {
	if s.OAPISchema == nil {
		return false
	}
	if s.OAPISchema.Default != nil {
		return true
	}
	for _, prop := range s.OAPISchema.Properties {
		if prop.Value != nil && prop.Value.Default != nil {
			return true
		}
	}
	for _, property := range s.Properties {
		if property.Schema.HasDefaults() {
			return true
		}
	}
	return false
}

type Property struct {
	Description    string
	JsonFieldName  string
	Schema         Schema
	Required       bool
	Nullable       bool
	ReadOnly       bool
	WriteOnly      bool
	NeedsFormTag   bool
	ExtensionProps *openapi3.ExtensionProps
}

func (p Property) GoFieldName() string {
	return SchemaNameToTypeName(p.JsonFieldName)
}

func (p Property) GoTypeDef() string {
	typeDef := p.Schema.TypeDecl()
	if !p.Schema.SkipOptionalPointer &&
		(!p.Required || p.Nullable || p.ReadOnly || p.WriteOnly) {

		typeDef = "*" + typeDef
	}
	return typeDef
}

// EnumDefinition holds type information for enum
type EnumDefinition struct {
	// Schema is the scheme of a type which has a list of enum values, eg, the
	// "container" of the enum.
	Schema Schema
	// TypeName is the name of the enum's type, usually aliased from something.
	TypeName string
	// ValueWrapper wraps the value. It's used to conditionally apply quotes
	// around strings.
	ValueWrapper string
	// Conflicts is set to true when this enum conflicts with another in
	// terms of TypeNames
	Conflicts bool
}

// GetValues generates enum names in a way to minimize global conflicts
func (e *EnumDefinition) GetValues() map[string]string {
	// in case there are no conflicts, it's safe to use the values as-is
	if !e.Conflicts {
		return e.Schema.EnumValues
	}
	// If we do have conflicts, we will prefix the enum's typename to the values.
	newValues := make(map[string]string, len(e.Schema.EnumValues))
	for k, v := range e.Schema.EnumValues {
		newName := e.TypeName + UppercaseFirstCharacter(k)
		newValues[newName] = v
	}
	return newValues
}

type Constants struct {
	// SecuritySchemeProviderNames holds all provider names for security schemes.
	SecuritySchemeProviderNames []string
	// EnumDefinitions holds type and value information for all enums
	EnumDefinitions []EnumDefinition
}

// TypeDefinition describes a Go type definition in generated code.
//
// Let's use this example schema:
// components:
//  schemas:
//    Person:
//      type: object
//      properties:
//      name:
//        type: string
type TypeDefinition struct {
	// The name of the type, eg, type <...> Person
	TypeName string

	// The name of the corresponding JSON description, as it will sometimes
	// differ due to invalid characters.
	JsonName string

	// This is the Schema wrapper is used to populate the type description
	Schema Schema
}

func (t *TypeDefinition) IsCustomType() bool {
	if t.Schema.OAPISchema == nil {
		return false
	}
	_, xGoTypeExists := t.Schema.OAPISchema.Extensions[extPropGoType]
	return xGoTypeExists
}

// ResponseTypeDefinition is an extension of TypeDefinition, specifically for
// response unmarshaling in ClientWithResponses.
type ResponseTypeDefinition struct {
	TypeDefinition
	// The content type name where this is used, eg, application/json
	ContentTypeName string

	// The type name of a response model.
	ResponseName string
}

func (t *ResponseTypeDefinition) IsGenericError() bool {
	if t.Schema.OAPISchema == nil {
		return false
	}
	if extGenericErr, err := extBool(t.Schema.OAPISchema.Extensions[extPropGenericErrResponse]); err == nil && extGenericErr {
		return true
	}
	return false
}

func (t *TypeDefinition) IsAlias() bool {
	return !options.Compatibility.OldAliasing && t.Schema.DefineViaAlias
}

func PropertiesEqual(a, b Property) bool {
	return a.JsonFieldName == b.JsonFieldName && a.Schema.TypeDecl() == b.Schema.TypeDecl() && a.Required == b.Required
}

func GenerateGoSchema(sref *openapi3.SchemaRef, path []string) (Schema, error) {
	// Add a fallback value in case the sref is nil.
	// i.e. the parent schema defines a type:array, but the array has
	// no items defined. Therefore we have at least valid Go-Code.
	if sref == nil {
		return defineIsSkipOptionalPointer(Schema{GoType: "interface{}"}), nil
	}

	schema := sref.Value

	// If Ref is set on the SchemaRef, it means that this type is actually a reference to
	// another type. We're not de-referencing, so simply use the referenced type.
	if IsGoTypeReference(sref.Ref) {
		var refType string
		var err error

		// try to get custom type
		if extension, ok := schema.Extensions[extPropGoType]; ok {
			refType, err = extString(extension)
			if err != nil {
				return Schema{}, fmt.Errorf("invalid value for %q: %w", extPropGoType, err)
			}
		} else {
			// Or convert the reference path to Go type
			refType, err = RefPathToGoType(sref.Ref)
		}

		if err != nil {
			return Schema{}, fmt.Errorf("error turning reference (%s) into a Go type: %s",
				sref.Ref, err)
		}
		return defineIsSkipOptionalPointer(Schema{
			GoType:         refType,
			Description:    StringToGoComment(schema.Description),
			OAPISchema:     schema,
			DefineViaAlias: true,
		}), nil
	}

	outSchema := Schema{
		Description: StringToGoComment(schema.Description),
		OAPISchema:  schema,
	}

	// We can't support this in any meaningful way
	if schema.AnyOf != nil {
		outSchema.GoType = "interface{}"
		return defineIsSkipOptionalPointer(outSchema), nil
	}
	// We can't support this in any meaningful way
	if schema.OneOf != nil {
		outSchema.GoType = "interface{}"
		return defineIsSkipOptionalPointer(outSchema), nil
	}

	// AllOf is interesting, and useful. It's the union of a number of other
	// schemas. A common usage is to create a union of an object with an ID,
	// so that in a RESTful paradigm, the Create operation can return
	// (object, id), so that other operations can refer to (id)
	if schema.AllOf != nil {
		mergedSchema, err := MergeSchemas(schema.AllOf, path)
		if err != nil {
			return Schema{}, fmt.Errorf("error merging schemas: %w", err)
		}
		mergedSchema.OAPISchema = schema
		return defineIsSkipOptionalPointer(mergedSchema), nil
	}

	// Check for custom Go type extension
	if extension, ok := schema.Extensions[extPropGoType]; ok {
		typeName, err := extString(extension)
		if err != nil {
			return outSchema, fmt.Errorf("invalid value for %q: %w", extPropGoType, err)
		}
		outSchema.GoType = typeName
		return defineIsSkipOptionalPointer(outSchema), nil
	}

	// Schema type and format, eg. string / binary
	t := schema.Type
	// Handle objects and empty schemas first as a special case
	if t == "" || t == "object" {
		var outType string

		if len(schema.Properties) == 0 && !SchemaHasAdditionalProperties(schema) {
			// If the object has no properties or additional properties, we
			// have some special cases for its type.
			if t == "object" {
				// We have an object with no properties. This is a generic object
				// expressed as a map.
				outType = "map[string]interface{}"
			} else { // t == ""
				// If we don't even have the object designator, we're a completely
				// generic type.
				outType = "interface{}"
			}
			outSchema.GoType = outType
			outSchema.DefineViaAlias = true
		} else {
			// When we define an object, we want it to be a type definition,
			// not a type alias, eg, "type Foo struct {...}"
			outSchema.DefineViaAlias = false
			// We've got an object with some properties.
			for _, pName := range SortedSchemaKeys(schema.Properties) {
				p := schema.Properties[pName]
				propertyPath := append(path, pName)
				pSchema, err := GenerateGoSchema(p, propertyPath)
				if err != nil {
					return Schema{}, fmt.Errorf("error generating Go schema for property '%s': %w", pName, err)
				}

				required := StringInArray(pName, schema.Required)

				if pSchema.HasAdditionalProperties && pSchema.RefType == "" && len(pSchema.Properties) != 0 {
					// If we have fields present which have additional properties,
					// but are not a pre-defined type, we need to define a type
					// for them, which will be based on the field names we followed
					// to get to the type.
					typeName := PathToTypeName(propertyPath...)

					typeDef := TypeDefinition{
						TypeName: typeName,
						JsonName: strings.Join(propertyPath, "."),
						Schema:   pSchema,
					}
					pSchema.AdditionalTypes = append(pSchema.AdditionalTypes, typeDef)

					pSchema.RefType = typeName
				}
				description := ""
				if p.Value != nil {
					description = p.Value.Description
				}
				prop := Property{
					JsonFieldName: pName,
					Schema:        pSchema,
					Required: required ||
						options.OutputOptions.ExplicitNullable && !p.Value.Nullable,
					Description:    description,
					Nullable:       p.Value.Nullable,
					ReadOnly:       p.Value.ReadOnly,
					WriteOnly:      p.Value.WriteOnly,
					ExtensionProps: &p.Value.ExtensionProps,
				}
				outSchema.Properties = append(outSchema.Properties, prop)
			}

			outSchema.HasAdditionalProperties = SchemaHasAdditionalProperties(schema)
			outSchema.AdditionalPropertiesType = &Schema{
				GoType: "interface{}",
			}
			if schema.AdditionalProperties != nil {
				additionalSchema, err := GenerateGoSchema(schema.AdditionalProperties, path)
				if err != nil {
					return Schema{}, fmt.Errorf("error generating type for additional properties: %w", err)
				}
				outSchema.AdditionalPropertiesType = &additionalSchema
			}

			outSchema.GoType = GenStructFromSchema(outSchema)
		}
		return defineIsSkipOptionalPointer(outSchema), nil
	} else if len(schema.Enum) > 0 {
		err := oapiSchemaToGoType(schema, path, &outSchema)
		// Enums need to be typed, so that the values aren't interchangeable,
		// so no matter what schema conversion thinks, we need to define a
		// new type.
		outSchema.DefineViaAlias = false

		if err != nil {
			return Schema{}, fmt.Errorf("error resolving primitive type: %w", err)
		}
		enumValues := make([]string, len(schema.Enum))
		for i, enumValue := range schema.Enum {
			enumValues[i] = fmt.Sprintf("%v", enumValue)
		}

		sanitizedValues := SanitizeEnumNames(enumValues)
		outSchema.EnumValues = make(map[string]string, len(sanitizedValues))

		for k, v := range sanitizedValues {
			var enumName string
			if v == "" {
				enumName = "Empty"
			} else {
				enumName = k
			}
			if options.Compatibility.OldEnumConflicts {
				outSchema.EnumValues[SchemaNameToTypeName(PathToTypeName(append(path, enumName)...))] = v
			} else {
				outSchema.EnumValues[SchemaNameToTypeName(k)] = v
			}
		}
		if len(path) > 1 { // handle additional type only on non-toplevel types
			typeName := SchemaNameToTypeName(PathToTypeName(path...))
			typeDef := TypeDefinition{
				TypeName: typeName,
				JsonName: strings.Join(path, "."),
				Schema:   outSchema,
			}
			outSchema.AdditionalTypes = append(outSchema.AdditionalTypes, typeDef)
			outSchema.RefType = typeName
		}
		//outSchema.RefType = typeName
	} else {
		err := oapiSchemaToGoType(schema, path, &outSchema)
		if err != nil {
			return Schema{}, fmt.Errorf("error resolving primitive type: %w", err)
		}
	}
	return defineIsSkipOptionalPointer(outSchema), nil
}

// oapiSchemaToGoType converts an OpenApi schema into a Go type definition for
// all non-object types.
func oapiSchemaToGoType(schema *openapi3.Schema, path []string, outSchema *Schema) error {
	f := schema.Format
	t := schema.Type

	switch t {
	case "array":
		// For arrays, we'll get the type of the Items and throw a
		// [] in front of it.
		arrayType, err := GenerateGoSchema(schema.Items, path)
		if err != nil {
			return fmt.Errorf("error generating type for array: %w", err)
		}
		outSchema.ArrayType = &arrayType
		outSchema.GoType = "[]" + arrayType.TypeDecl()
		outSchema.AdditionalTypes = arrayType.AdditionalTypes
		outSchema.Properties = arrayType.Properties
		outSchema.DefineViaAlias = true
	case "integer":
		// We default to int if format doesn't ask for something else.
		if f == "int64" {
			outSchema.GoType = "int64"
		} else if f == "int32" {
			outSchema.GoType = "int32"
		} else if f == "int16" {
			outSchema.GoType = "int16"
		} else if f == "int8" {
			outSchema.GoType = "int8"
		} else if f == "int" {
			outSchema.GoType = "int"
		} else if f == "uint64" {
			outSchema.GoType = "uint64"
		} else if f == "uint32" {
			outSchema.GoType = "uint32"
		} else if f == "uint16" {
			outSchema.GoType = "uint16"
		} else if f == "uint8" {
			outSchema.GoType = "uint8"
		} else if f == "uint" {
			outSchema.GoType = "uint"
		} else if f == "" {
			outSchema.GoType = "int"
		} else {
			return fmt.Errorf("invalid integer format: %s", f)
		}
		outSchema.DefineViaAlias = true
	case "number":
		// We default to float for "number"
		if f == "double" {
			outSchema.GoType = "float64"
		} else if f == "float" || f == "" {
			outSchema.GoType = "float32"
		} else {
			return fmt.Errorf("invalid number format: %s", f)
		}
		outSchema.DefineViaAlias = true
	case "boolean":
		if f != "" {
			return fmt.Errorf("invalid format (%s) for boolean", f)
		}
		outSchema.GoType = "bool"
		outSchema.DefineViaAlias = true
	case "string":
		// Special case string formats here.
		switch f {
		case "byte":
			outSchema.GoType = "[]byte"
		case "email":
			outSchema.GoType = "openapi_types.Email"
		case "date":
			outSchema.GoType = "openapi_types.Date"
		case "date-time":
			outSchema.GoType = "time.Time"
		case "json":
			outSchema.GoType = "json.RawMessage"
			outSchema.SkipOptionalPointer = true
		case "uuid":
			outSchema.GoType = "openapi_types.UUID"
		default:
			// All unrecognized formats are simply a regular string.
			outSchema.GoType = "string"
		}
		outSchema.DefineViaAlias = true
	default:
		return fmt.Errorf("unhandled Schema type: %s", t)
	}

	outSchema.SkipOptionalPointer = isSkipOptionalPointer(t, f)
	return nil
}

func isSkipOptionalPointer(tType string, format string) bool {
	switch tType {
	case "array":
		return true
	case "string":
		// Special case string formats here.
		switch format {
		case "byte", "json":
			return true
		}
	}
	return false
}

func defineIsSkipOptionalPointer(schema Schema) Schema {
	if schema.OAPISchema != nil {
		t := schema.OAPISchema.Type
		f := schema.OAPISchema.Format

		if schema.SkipOptionalPointer || isSkipOptionalPointer(t, f) {
			schema.SkipOptionalPointer = true
			return schema
		}
		if len(schema.OAPISchema.Properties) == 0 && len(schema.OAPISchema.AllOf) == 0 && (t == "" || t == "object") {
			schema.SkipOptionalPointer = true
			return schema
		}
		if schema.OAPISchema.Default != nil {
			schema.SkipOptionalPointer = true
			return schema
		}
	}

	switch schema.GoType {
	case "interface{}", "[]byte":
		schema.SkipOptionalPointer = true
		return schema
	}

	if schema.IsAdditionalPropertiesOnly() {
		schema.SkipOptionalPointer = true
		return schema
	}

	return schema
}

// SchemaDescriptor describes a Schema, a type definition.
type SchemaDescriptor struct {
	Fields                   []FieldDescriptor
	HasAdditionalProperties  bool
	AdditionalPropertiesType string
}

type FieldDescriptor struct {
	Required bool   // Is the schema required? If not, we'll pass by pointer
	GoType   string // The Go type needed to represent the json type.
	GoName   string // The Go compatible type name for the type
	JsonName string // The json type name for the type
	IsRef    bool   // Is this schema a reference to predefined object?
}

// Given a list of schema descriptors, produce corresponding field names with
// JSON annotations
func GenFieldsFromProperties(props []Property) []string {
	var fields []string
	for i, p := range props {
		field := ""
		// Add a comment to a field in case we have one, otherwise skip.
		if p.Description != "" {
			// Separate the comment from a previous-defined, unrelated field.
			// Make sure the actual field is separated by a newline.
			if i != 0 {
				field += "\n"
			}
			field += fmt.Sprintf("%s\n", StringToGoComment(p.Description))
		}

		goFieldName := p.GoFieldName()
		if _, ok := p.ExtensionProps.Extensions[extGoFieldName]; ok {
			if extGoFieldName, err := extString(p.ExtensionProps.Extensions[extGoFieldName]); err == nil {
				goFieldName = extGoFieldName
			}
		}

		field += fmt.Sprintf("    %s %s", goFieldName, p.GoTypeDef())

		// Support x-omitempty
		overrideOmitEmpty := true
		if _, ok := p.ExtensionProps.Extensions[extPropOmitEmpty]; ok {
			if extOmitEmpty, err := extBool(p.ExtensionProps.Extensions[extPropOmitEmpty]); err == nil {
				overrideOmitEmpty = extOmitEmpty
			}
		}

		fieldTags := make(map[string]string)

		if (p.Required && !p.ReadOnly && !p.WriteOnly) || !overrideOmitEmpty {
			fieldTags["json"] = p.JsonFieldName
			if p.NeedsFormTag {
				fieldTags["form"] = p.JsonFieldName
			}
		} else {
			fieldTags["json"] = p.JsonFieldName + ",omitempty"
			if p.NeedsFormTag {
				fieldTags["form"] = p.JsonFieldName + ",omitempty"
			}
		}

		if p.Schema.OAPISchema != nil {
			if p.Schema.OAPISchema.Default != nil {
				fieldTags["default"] = fmt.Sprint(p.Schema.OAPISchema.Default)
			}
			if validations := getValidationTagsForInputSchema(p); validations != "" {
				fieldTags["validate"] = validations
			}
		}

		if extension, ok := p.ExtensionProps.Extensions[extPropExtraTags]; ok {
			if tags, err := extExtraTags(extension); err == nil {
				keys := SortedStringKeys(tags)
				for _, k := range keys {
					fieldTags[k] = tags[k]
				}
			}
		}
		// Convert the fieldTags map into Go field annotations.
		keys := SortedStringKeys(fieldTags)
		tags := make([]string, len(keys))
		for i, k := range keys {
			tags[i] = fmt.Sprintf(`%s:"%s"`, k, fieldTags[k])
		}
		field += "`" + strings.Join(tags, " ") + "`"
		fields = append(fields, field)
	}
	return fields
}

func getValidationTagsForInputSchema(property Property) string {
	schema := property.Schema.OAPISchema

	validations := make([]string, 0, 1)
	if property.Required {
		validations = append(validations, "required")
	} else {
		validations = append(validations, "omitempty")
	}
	if len(schema.Enum) > 0 {
		validations = append(validations, "oneof="+util.JoinInterfaces(schema.Enum, " "))
	}
	if schema.Min != nil {
		validations = append(validations, fmt.Sprintf("min=%v", *schema.Min))
	}
	if schema.Max != nil {
		validations = append(validations, fmt.Sprintf("max=%v", *schema.Max))
	}
	if schema.MinLength != 0 {
		validations = append(validations, fmt.Sprintf("min=%d", schema.MinLength))
	}
	if schema.MaxLength != nil {
		validations = append(validations, fmt.Sprintf("max=%d", *schema.MaxLength))
	}
	if schema.Pattern != "" {
		validations = append(validations, fmt.Sprintf("pattern=%s", base64.StdEncoding.EncodeToString([]byte(schema.Pattern))))
	}
	if extValidate, ok := schema.Extensions[extPropValidate]; ok {
		validate, err := extString(extValidate)
		if err != nil {
			panic(errors.New(fmt.Sprintf("failed to parse %s:%s", extPropValidate, err.Error())))
		}
		validations = append(validations, validate)
	}

	// to skip "omitempty" on each field
	if len(validations) == 1 && !property.Required {
		return ""
	}

	return strings.Join(validations, ",")
}

func GenStructFromSchema(schema Schema) string {
	var objectParts []string
	if schema.IsAdditionalPropertiesOnly() {
		addPropsType := schema.AdditionalPropertiesType.GoType
		if schema.AdditionalPropertiesType.RefType != "" {
			addPropsType = schema.AdditionalPropertiesType.RefType
		}

		objectParts = append(objectParts,
			fmt.Sprintf("map[string]%s", addPropsType))
	} else {
		// Start out with struct {
		objectParts = []string{"struct {"}
		// Append all the field definitions
		objectParts = append(objectParts, GenFieldsFromProperties(schema.Properties)...)
		// Close the struct
		if schema.HasAdditionalProperties {
			addPropsType := schema.AdditionalPropertiesType.GoType
			if schema.AdditionalPropertiesType.RefType != "" {
				addPropsType = schema.AdditionalPropertiesType.RefType
			}

			objectParts = append(objectParts,
				fmt.Sprintf("AdditionalProperties map[string]%s `json:\"-\"`", addPropsType))
		}
		objectParts = append(objectParts, "}")
	}
	return strings.Join(objectParts, "\n")
}

// This constructs a Go type for a parameter, looking at either the schema or
// the content, whichever is available
func paramToGoType(param *openapi3.Parameter, path []string) (Schema, error) {
	if param.Content == nil && param.Schema == nil {
		return Schema{}, fmt.Errorf("parameter '%s' has no schema or content", param.Name)
	}

	// We can process the schema through the generic schema processor
	if param.Schema != nil {
		return GenerateGoSchema(param.Schema, path)
	}

	// At this point, we have a content type. We know how to deal with
	// application/json, but if multiple formats are present, we can't do anything,
	// so we'll return the parameter as a string, not bothering to decode it.
	if len(param.Content) > 1 {
		return Schema{
			GoType:      "string",
			Description: StringToGoComment(param.Description),
		}, nil
	}

	// Otherwise, look for application/json in there
	mt, found := param.Content["application/json"]
	if !found {
		// If we don't have json, it's a string
		return Schema{
			GoType:      "string",
			Description: StringToGoComment(param.Description),
		}, nil
	}

	// For json, we go through the standard schema mechanism
	return GenerateGoSchema(mt.Schema, path)
}
