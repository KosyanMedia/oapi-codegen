package schemas

import "testing"

func Test_defaults_non_nullable(t *testing.T) {
	t.Parallel()

	// must be compiled successfully
	_ = EnsuredefaultsnopointersJSONRequestBodySchema{
		BoolField:   false,
		FloatField:  0.0,
		IntField:    0,
		StringField: "",
	}
}
