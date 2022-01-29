package file_test

import (
	"errors"
	file_api "file_upload/api/v1/file"
	"file_upload/internal/file"
	"file_upload/internal/test"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"
)

func Test_handler_HandleGetFile(t *testing.T) {
	tests := []struct {
		name               string
		fileService        file.Service
		r                  *http.Request
		expectedStatusCode int
		expectedResponse   []byte
	}{
		{
			name: "valid",
			r: &http.Request{
				URL: &url.URL{
					Path: "/api/v1/file/some-file.png",
				},
			},
			fileService:        test.NewMockFileService().WithExpectedGetFileResponse([]byte(`img`), true),
			expectedStatusCode: http.StatusOK,
			expectedResponse:   []byte(`img`),
		},
		{
			name: "file not found",
			r: &http.Request{
				URL: &url.URL{
					Path: "/api/v1/file/some-file.png",
				},
			},
			fileService:        test.NewMockFileService().WithExpectedGetFileResponse(nil, false),
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   []byte(`{}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := file_api.NewFileHandlerWithInterfaces(tt.fileService, "")
			response := h.HandleGetFile(tt.r)
			test.AssertInt(t, tt.expectedStatusCode, response.StatusCode)
			test.AssertByteArray(t, tt.expectedResponse, response.Response)
		})
	}
}

func Test_handler_HandleUpload(t *testing.T) {
	tests := []struct {
		name               string
		fileService        file.Service
		r                  *http.Request
		expectedStatusCode int
		expectedResponse   []byte
	}{
		{
			name: "invalid file",
			r: &http.Request{
				MultipartForm: &multipart.Form{
					File: map[string][]*multipart.FileHeader{
						"file": {
							{
								Filename: "test.png",
								Header:   nil,
								Size:     0,
							},
						},
					},
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   []byte("{\"message\":\"`file` is a required field\"}"),
		},
		{
			name:               "store file fails",
			r:                  test.CreateMultipartRequest("file", "test.png"),
			fileService:        test.NewMockFileService().WithExpectedStoreFileResponse(errors.New("error")),
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   []byte(`{"message":"failed to store file"}`),
		},
		{
			name:               "successful file upload",
			r:                  test.CreateMultipartRequest("file", "test.png"),
			fileService:        test.NewMockFileService().WithExpectedStoreFileResponse(nil),
			expectedStatusCode: http.StatusOK,
			expectedResponse:   []byte(`{"file_name":"test.png","asset_url":"https://domain.com/api/v1/file/test.png"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := file_api.NewFileHandlerWithInterfaces(tt.fileService, "https://domain.com")

			response := h.HandleUpload(tt.r)
			test.AssertInt(t, tt.expectedStatusCode, response.StatusCode)
			test.AssertByteArray(t, tt.expectedResponse, response.Response)
		})
	}
}
