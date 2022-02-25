package codegen

import (
	"net/http"
	"testing"
)

func TestGenerateDefaultOperationID(t *testing.T) {
	type test struct {
		op      string
		path    string
		want    string
		wantErr bool
	}

	suite := []test{
		{
			op:      http.MethodGet,
			path:    "/v1/foo/bar",
			want:    "GetV1FooBar",
			wantErr: false,
		},
		{
			op:      http.MethodGet,
			path:    "/v1/foo/bar/",
			want:    "GetV1FooBar",
			wantErr: false,
		},
		{
			op:      http.MethodPost,
			path:    "/v1",
			want:    "PostV1",
			wantErr: false,
		},
		{
			op:      http.MethodPost,
			path:    "v1",
			want:    "PostV1",
			wantErr: false,
		},
		{
			path:    "v1",
			want:    "",
			wantErr: true,
		},
		{
			path:    "",
			want:    "PostV1",
			wantErr: true,
		},
	}

	for _, test := range suite {
		got, err := generateDefaultOperationID(test.op, test.path)
		if err != nil {
			if !test.wantErr {
				t.Fatalf("did not expected error but got %v", err)
			}
		}

		if test.wantErr {
			return
		}
		if got != test.want {
			t.Fatalf("Operation ID generation error. Want [%v] Got [%v]", test.want, got)
		}
	}
}
