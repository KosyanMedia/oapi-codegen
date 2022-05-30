package no_client_editors

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompiled(t *testing.T) {
	var (
		withTrailingSlash    string = "https://my-api.com/some-base-url/v1/"
		withoutTrailingSlash string = "https://my-api.com/some-base-url/v1"
	)

	client1, err := NewClient(
		withTrailingSlash,
	)
	assert.NoError(t, err)

	client2, err := NewClient(
		withoutTrailingSlash,
	)
	assert.NoError(t, err)

	client3, err := NewClient(
		"",
		WithBaseURL(withTrailingSlash),
	)
	assert.NoError(t, err)

	client4, err := NewClient(
		"",
		WithBaseURL(withoutTrailingSlash),
	)
	assert.NoError(t, err)

	expectedURL := withTrailingSlash

	assert.Equal(t, expectedURL, client1.Server)
	assert.Equal(t, expectedURL, client2.Server)
	assert.Equal(t, expectedURL, client3.Server)
	assert.Equal(t, expectedURL, client4.Server)
	assert.Nil(t, client1.RequestEditors)
	assert.Nil(t, client2.RequestEditors)
	assert.Nil(t, client3.RequestEditors)
	assert.Nil(t, client4.RequestEditors)
	assert.Equal(t, 1, reflect.TypeOf(client1.GetJson).NumIn())
}
