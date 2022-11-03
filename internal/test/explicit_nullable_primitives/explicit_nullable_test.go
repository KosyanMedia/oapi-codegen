package explicit_nullable_primitives

import "testing"

func TestExplicitNullable(t *testing.T) {
	t.Parallel()

	// should be successfully compiled
	_ = Entity{
		IntField:            1,
		IntFieldNullable:    nil,
		StringField:         "value",
		StringFieldNullable: nil,
		ObjectField: &NestedEntity{
			BoolField:           false,
			BoolFieldNullable:   nil,
			NumberField:         1.0,
			NumberFieldNullable: nil,
		},
	}
}
