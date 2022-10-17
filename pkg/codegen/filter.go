package codegen

import (
	"github.com/getkin/kin-openapi/openapi3"
	"regexp"
	"strings"
)

const (
	regexPrefix = "$regex:"
	notPrefix   = "$not:"
)

func filterOperationsByTag(swagger *openapi3.T, opts Configuration) {
	if len(opts.OutputOptions.ExcludeTags) > 0 {
		excludeOperationsWithTags(swagger.Paths, opts.OutputOptions.ExcludeTags)
	}
	if len(opts.OutputOptions.IncludeTags) > 0 {
		includeOperationsWithTags(swagger.Paths, opts.OutputOptions.IncludeTags, false)
	}
}

func excludeOperationsWithTags(paths openapi3.Paths, tags []string) {
	includeOperationsWithTags(paths, tags, true)
}

func includeOperationsWithTags(paths openapi3.Paths, tags []string, exclude bool) {
	for _, pathItem := range paths {
		ops := pathItem.Operations()
		names := make([]string, 0, len(ops))
		for name, op := range ops {
			if operationHasTag(op, tags) == exclude {
				names = append(names, name)
			}
		}
		for _, name := range names {
			pathItem.SetOperation(name, nil)
		}
	}
}

//operationHasTag returns true if the operation is tagged with any of tags
func operationHasTag(op *openapi3.Operation, tags []string) bool {
	if op == nil {
		return false
	}
	for _, hasTag := range op.Tags {
		for _, wantTag := range tags {
			if hasTag == wantTag {
				return true
			}
		}
	}
	return false
}

func configureOperationsByOptions(swagger *openapi3.T, opts Configuration) {
	for path, pathItem := range swagger.Paths {
		for _, operationOpt := range opts.OutputOptions.Operations {
			if applyOption(path, pathItem, operationOpt) {
				delete(swagger.Paths, path)
			}
		}
	}
}

func applyOption(path string, pathItem *openapi3.PathItem, operationOpt OperationOption) bool {
	if operationOpt.Path != "" && !strMatch(path, operationOpt.Path) {
		return false
	}
	if operationOpt.Method == "" {
		if operationOpt.Exclude {
			return true
		}
		pathItem.Parameters = filterParamsByOptions(pathItem.Parameters, operationOpt.Params)
	}
	for method, op := range pathItem.Operations() {
		if operationOpt.Method == "" || strMatch(method, operationOpt.Method) {
			if operationOpt.Exclude {
				pathItem.SetOperation(method, nil)
			} else {
				op.Parameters = filterParamsByOptions(op.Parameters, operationOpt.Params)
			}
		}
	}
	return false
}

func filterParamsByOptions(parameters openapi3.Parameters, params []ParamOption) openapi3.Parameters {
	result := make(openapi3.Parameters, 0, len(parameters))
	for _, param := range parameters {
		exclude := false
		for _, paramOpt := range params {
			if strMatch(param.Value.Name, paramOpt.Name) && paramOpt.Exclude {
				exclude = true
				break
			}
		}
		if !exclude {
			result = append(result, param)
		}
	}
	return result
}

func strMatch(original string, pattern string) bool {
	if strings.HasPrefix(pattern, regexPrefix) {
		match, err := regexp.MatchString(pattern[len(regexPrefix):], original)
		return err == nil && match
	}
	if strings.HasPrefix(pattern, notPrefix) {
		pattern := pattern[len(notPrefix):]
		return pattern != original
	}
	return strings.EqualFold(original, pattern)
}
