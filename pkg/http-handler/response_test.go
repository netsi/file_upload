package http_handler_test

import (
	http_handler "file_upload/pkg/http-handler"
	"net/http"
	"reflect"
	"testing"
)

func TestImage(t *testing.T) {
	type args struct {
		statusCode int
		bytes      []byte
	}
	tests := []struct {
		name string
		args args
		want *http_handler.Response
	}{
		{
			name: "check",
			args: args{
				statusCode: http.StatusOK,
				bytes:      []byte{1, 3, 2},
			},
			want: &http_handler.Response{
				Headers: map[string]string{
					http_handler.ContentTypeHeader: http_handler.StreamContent,
				},
				StatusCode: http.StatusOK,
				Response:   []byte{1, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := http_handler.Image(tt.args.statusCode, tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Image() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON(t *testing.T) {
	type args struct {
		statusCode int
		resp       interface{}
	}
	tests := []struct {
		name string
		args args
		want *http_handler.Response
	}{
		{
			name: "check",
			args: args{
				statusCode: http.StatusOK,
				resp: struct {
					Name string `json:"name"`
				}{Name: "john"},
			},
			want: &http_handler.Response{
				Headers: map[string]string{
					http_handler.ContentTypeHeader: http_handler.JsonContentType,
				},
				StatusCode: http.StatusOK,
				Response:   []byte(`{"name":"john"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := http_handler.JSON(tt.args.statusCode, tt.args.resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
