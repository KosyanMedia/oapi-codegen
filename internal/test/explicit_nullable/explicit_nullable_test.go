package explicit_nullable

import "testing"

func TestExplicitNullable(t *testing.T) {
	t.Parallel()

	// should be successfully compiled
	_ = Entity{
		IntField:            1,
		IntFieldNullable:    nil,
		StringField:         "value",
		StringFieldNullable: nil,
	}
}
