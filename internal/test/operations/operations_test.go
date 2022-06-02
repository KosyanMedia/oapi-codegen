package operations

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCompiled(t *testing.T) {
	client, err := NewClient("https://my-api.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	// first
	assert.NotNil(t, client.ShouldHaveBoth)
	shouldHaveBothParam := ShouldHaveBothParams{
		First:  nil,
		Second: nil,
	}
	assert.Equal(t, 2, reflect.TypeOf(shouldHaveBothParam).NumField())

	// second
	assert.NotNil(t, client.ShouldHaveSecond)
	shouldHaveSecondParam := ShouldHaveSecondParams{
		Second: nil,
	}
	assert.Equal(t, 1, reflect.TypeOf(shouldHaveSecondParam).NumField())

	// third
	assert.NotNil(t, client.LeavePostOnlyPost)
	leavePostOnlyPostParam := LeavePostOnlyPostParams{
		First:  nil,
		Second: nil,
	}
	assert.Equal(t, 2, reflect.TypeOf(leavePostOnlyPostParam).NumField())

	// assert no more
	assertExportedMethods(t, 3, client)
}

func assertExportedMethods(t *testing.T, expected int, service interface{}) {
	typ := reflect.TypeOf(service)
	actual := 0
	for i := 0; i < typ.NumMethod(); i++ {
		if typ.Method(i).IsExported() {
			actual++
		}
	}
	assert.Equal(t, expected, actual)
}
