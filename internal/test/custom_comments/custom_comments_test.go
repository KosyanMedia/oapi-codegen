package custom_comments

import (
	"github.com/stretchr/testify/require"
	"go/doc"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestOneCommentGenerated(t *testing.T) {
	t.Parallel()

	doc := getStructDoc(t, reflect.TypeOf(OneComment{}).Name())

	require.Contains(t, doc, "easyjson:json")
}

func TestMultilineCommentGenerated(t *testing.T) {
	t.Parallel()

	doc := getStructDoc(t, reflect.TypeOf(MultilineComment{}).Name())
	require.Contains(t, doc, "first line\nsecond line")
}

func getStructDoc(t *testing.T, name string) string {
	fset := token.NewFileSet()

	d, err := parser.ParseDir(fset, "./", nil, parser.ParseComments)
	require.Nil(t, err)

	for _, f := range d {
		p := doc.New(f, "./types.gen.go", 0)

		for _, t := range p.Types {
			if t.Name == name {
				return t.Doc
			}
		}
	}
	require.Fail(t, "Struct {} not found", name)
	return ""
}
