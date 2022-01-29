package http_handler_test

import (
	"bytes"
	"file_upload/internal/test"
	http_handler "file_upload/pkg/http-handler"
	"net/http"
	"testing"
)

func Test_handler_Handle(t *testing.T) {
	tests := []struct {
		name               string
		requestMethod      string
		expectedStatusCode int
		expectedHeaders    map[string]string
		expectedResponse   []byte
		postRequestHandler http_handler.RequestHandler
		getRequestHandler  http_handler.RequestHandler
	}{
		{
			name:               "no handlers configured",
			requestMethod:      http.MethodPost,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "no handlers found",
			requestMethod:      http.MethodPost,
			expectedStatusCode: http.StatusNotFound,
			getRequestHandler: func(r *http.Request) *http_handler.Response {
				return nil
			},
		},
		{
			name:               "get request handler returns no response",
			requestMethod:      http.MethodGet,
			expectedStatusCode: http.StatusNoContent,
			getRequestHandler: func(r *http.Request) *http_handler.Response {
				return nil
			},
		},
		{
			name:               "post request handler returns no response",
			requestMethod:      http.MethodPost,
			expectedStatusCode: http.StatusNoContent,
			postRequestHandler: func(r *http.Request) *http_handler.Response {
				return nil
			},
		},
		{
			name:               "post request handler returns response",
			requestMethod:      http.MethodPost,
			expectedStatusCode: http.StatusOK,
			expectedHeaders:    map[string]string{http_handler.ContentTypeHeader: http_handler.JsonContentType},
			expectedResponse:   []byte("ok"),
			postRequestHandler: func(r *http.Request) *http_handler.Response {
				return &http_handler.Response{
					Headers:    map[string]string{http_handler.ContentTypeHeader: http_handler.JsonContentType},
					StatusCode: http.StatusOK,
					Response:   []byte("ok"),
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := test.NewMockResponseWritter()

			h := http_handler.NewHandler()
			if tt.postRequestHandler != nil {
				h.POST(tt.postRequestHandler)
			}
			if tt.getRequestHandler != nil {
				h.GET(tt.getRequestHandler)
			}

			h.Handle(responseWriter, &http.Request{Method: tt.requestMethod})

			if responseWriter.StatusCode != tt.expectedStatusCode {
				t.Errorf("expected %d status code and received %d", tt.expectedStatusCode, responseWriter.StatusCode)
			}
			if res := bytes.Compare(responseWriter.Output, tt.expectedResponse); res != 0 {
				t.Errorf("expected %s output and received %s", tt.expectedResponse, responseWriter.Output)
			}
			for expectedHeaderKey, expectedHeaderVal := range tt.expectedHeaders {
				if responseWriter.Header().Get(expectedHeaderKey) != expectedHeaderVal {
					t.Errorf("expected %s header with value %s but received received %s", expectedHeaderKey, expectedHeaderVal, responseWriter.Output)
				}
			}
		})
	}
}
