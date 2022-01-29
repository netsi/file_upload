package file

import (
	internal_file "file_upload/internal/file"
	"file_upload/internal/image"
	"file_upload/internal/responses"
	http_handler "file_upload/pkg/http-handler"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	V1FileAPIPath = "/api/v1/file"
)

type handler struct {
	service   internal_file.Service
	publicURL string
}

// NewFileHandler is a factory of file.handler, initializes all the required dependencies and returns a
// file.handler instance.
func NewFileHandler(publicURL string) *handler {
	return NewFileHandlerWithInterfaces(internal_file.NewFileService(), publicURL)
}

// NewFileHandlerWithInterfaces returns a file.handler instance with the given dependencies.
func NewFileHandlerWithInterfaces(service internal_file.Service, publicURL string) *handler {
	return &handler{
		service:   service,
		publicURL: publicURL,
	}
}

// HandleUpload handles an upload request. Stores the given file to the filesystem and returns UploadResponse
// if successful. It also can return 400, 500 in case of errors.
func (h *handler) HandleUpload(r *http.Request) *http_handler.Response {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed to get uploaded file with error: %s", err.Error())
		return http_handler.JSON(http.StatusBadRequest, &responses.ErrorResponse{
			Message: "`file` is a required field",
		})
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("failed to read the file with error: %s", err.Error())
		return http_handler.JSON(http.StatusBadRequest, &responses.ErrorResponse{
			Message: "failed to read the file",
		})
	}

	if !image.IsSupportedImage(fileBytes) {
		return http_handler.JSON(http.StatusBadRequest, &responses.ErrorResponse{
			Message: fmt.Sprintf("unsupported image type, the API supports only %s", image.GetSupportedTypes()),
		})
	}

	err = h.service.StoreFile(fileHeader.Filename, fileBytes)
	if err != nil {
		log.Printf("failed to store the file with error: %s", err.Error())
		return http_handler.JSON(http.StatusInternalServerError, &responses.ErrorResponse{
			Message: "failed to store file",
		})
	}

	return http_handler.JSON(http.StatusOK, &UploadResponse{
		FileName:  fileHeader.Filename,
		AssertURL: fmt.Sprintf("%s%s/%s", h.publicURL, V1FileAPIPath, url.PathEscape(fileHeader.Filename)),
	})
}

// HandleGetFile fetches the file based on the given name. The filename is taken from the request URL Path.
// Returns a stream of bytes if the image is found, if it is not found returns 404 with empty json body.
func (h *handler) HandleGetFile(r *http.Request) *http_handler.Response {
	fileName := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/", V1FileAPIPath))
	unescapedPath, err := url.PathUnescape(fileName)
	if err != nil {
		log.Printf("failed to unescape the filename with error: %s", err.Error())
		return http_handler.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "invalid filename",
		})
	}

	fileBytes, found := h.service.GetFile(unescapedPath)
	if !found {
		return http_handler.JSON(http.StatusNotFound, responses.EmptyResponse{})
	}

	return http_handler.Image(http.StatusOK, fileBytes)
}
