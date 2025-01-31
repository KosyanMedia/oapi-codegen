package deep_structs

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDeepStructs(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 2, reflect.TypeOf(Coordinates{}).NumField())

	// should be successfully compiled
	_ = Entity{
		Embedded: &EntityEmbedded{
			Coordinate: &Coordinates{
				Latitude:  nil,
				Longitude: nil,
			},
			Id: 1,
		},
	}

	assert.Equal(t, 1, reflect.TypeOf(Entity{}).NumField())
	assert.Equal(t, 2, reflect.TypeOf(EntityEmbedded{}).NumField())

	// should be successfully compiled
	_ = Entities{
		Embedded: []EntitiesEmbeddedItem{
			{
				Coordinates: [][]Coordinates{
					{
						Coordinates{
							Latitude:  nil,
							Longitude: nil,
						},
					},
				},
				Id: 1,
			},
		},
	}

	assert.Equal(t, 1, reflect.TypeOf(Entities{}).NumField())
	assert.Equal(t, 2, reflect.TypeOf(EntitiesEmbeddedItem{}).NumField())

	mySlice := make(MySlice, 0)
	mySlice = append(mySlice, MySliceItem{
		Id: 1,
	})
	assert.Equal(t, 1, reflect.TypeOf(MySliceItem{}).NumField())

	// should be successfully compiled
	var val int = 1
	_ = KekResponse{
		Code:    200,
		JSON200: interface{}(nil),
		JSON404: &Kek404JSONResponseBodySchema{
			Numfield: &val,
		},
	}

	assert.Equal(t, 1, reflect.TypeOf(Kek404JSONResponseBodySchema{}).NumField())
	assert.Equal(t, 3, reflect.TypeOf(KekResponse{}).NumField())
}
