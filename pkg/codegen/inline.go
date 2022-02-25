package codegen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

// This generates a gzipped, base64 encoded JSON representation of the
// swagger definition, which we embed inside the generated code.
func GenerateInlinedSpec(t *template.Template, importMapping importMap, swagger *openapi3.T) (string, error) {
	// Marshal to json
	encoded, err := swagger.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("error marshaling swagger: %s", err)
	}

	// gzip
	var buf bytes.Buffer
	zw, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return "", fmt.Errorf("error creating gzip compressor: %s", err)
	}
	_, err = zw.Write(encoded)
	if err != nil {
		return "", fmt.Errorf("error gzipping swagger file: %s", err)
	}
	err = zw.Close()
	if err != nil {
		return "", fmt.Errorf("error gzipping swagger file: %s", err)
	}
	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	var parts []string
	const width = 80

	// Chop up the string into an array of strings.
	for len(str) > width {
		part := str[0:width]
		parts = append(parts, part)
		str = str[width:]
	}
	if len(str) > 0 {
		parts = append(parts, str)
	}

	return GenerateTemplates(
		[]string{"inline.tmpl"},
		t,
		struct {
			SpecParts     []string
			ImportMapping importMap
		}{
			SpecParts:     parts,
			ImportMapping: importMapping,
		})
}
