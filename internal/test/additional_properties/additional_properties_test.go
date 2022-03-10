package additional_properties

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const stringMapJson = `
{
	"string_field": "string_value",
	"key": "value"
}
`

const genericMapJson = `
{
	"string_field": "string_value",
	"key2": 2
}
`

const complexMapJson = `
{
	"simple_key": {
		"string_field": "string_value"
	}
}
`

func TestSerialization(t *testing.T) {
	t.Parallel()

	t.Run("MustBeMap", func(t *testing.T) {
		var val MustBeMap = make(map[string]string)
		val["key"] = "value"

		bytes, err := json.Marshal(val)
		assert.Nil(t, err)
		assert.Equal(t, `{"key":"value"}`, string(bytes))
	})

	t.Run("MustBeMapToo", func(t *testing.T) {
		var val MustBeMapToo = make(map[string]interface{})
		val["key"] = "value"
		val["key2"] = 2

		bytes, err := json.Marshal(val)
		assert.Nil(t, err)
		assert.Equal(t, `{"key":"value","key2":2}`, string(bytes))
	})

	t.Run("MustBeStruct", func(t *testing.T) {
		var val MustBeStruct
		val.StringField = "string_value"
		val.AdditionalProperties = make(map[string]string)
		val.AdditionalProperties["key"] = "value"

		bytes, err := json.Marshal(val)
		assert.Nil(t, err)
		assert.Equal(t, `{"key":"value","string_field":"string_value"}`, string(bytes))
	})

	t.Run("MustBeMapWithStructs", func(t *testing.T) {
		var val MustBeMapWithStructs = make(map[string]SimpleObject)
		val["key"] = SimpleObject{StringField: "string_value"}

		bytes, err := json.Marshal(val)
		assert.Nil(t, err)
		assert.Equal(t, `{"key":{"string_field":"string_value"}}`, string(bytes))
	})
}

func TestDeSerialization(t *testing.T) {
	t.Parallel()

	t.Run("MustBeMap", func(t *testing.T) {
		var val MustBeMap
		assert.Nil(t, json.Unmarshal([]byte(stringMapJson), &val))
		assert.Equal(t, 2, len(val))
		assert.Equal(t, "string_value", val["string_field"])
		assert.Equal(t, "value", val["key"])
	})

	t.Run("MustBeMapToo", func(t *testing.T) {
		var val MustBeMapToo
		assert.Nil(t, json.Unmarshal([]byte(genericMapJson), &val))
		assert.Equal(t, 2, len(val))
		assert.Equal(t, "string_value", val["string_field"])
		assert.Equal(t, float64(2), val["key2"])
	})

	t.Run("MustBeStruct", func(t *testing.T) {
		var val MustBeStruct
		assert.Nil(t, json.Unmarshal([]byte(stringMapJson), &val))
		assert.Equal(t, 1, len(val.AdditionalProperties))
		assert.Equal(t, "value", val.AdditionalProperties["key"])
		assert.Equal(t, "string_value", val.StringField)
	})

	t.Run("MustBeMapWithStructs", func(t *testing.T) {
		var val MustBeMapWithStructs
		assert.Nil(t, json.Unmarshal([]byte(complexMapJson), &val))
		assert.Equal(t, 1, len(val))
		obj := val["simple_key"]
		assert.Equal(t, obj, SimpleObject{StringField: "string_value"})
	})
}

func TestGetSet(t *testing.T) {
	t.Parallel()

	var val MustBeStruct
	val.Set("key", "value")

	value, found := val.Get("key")
	assert.True(t, found)
	assert.Equal(t, "value", value)
	assert.Equal(t, "value", val.AdditionalProperties["key"])
}
