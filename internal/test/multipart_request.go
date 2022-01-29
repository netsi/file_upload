package test

import (
	"bytes"
	http_handler "file_upload/pkg/http-handler"
	"image"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
)

func CreateMultipartRequest(key, filename string) *http.Request {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	part, _ := writer.CreateFormFile(key, filename)

	width := 10
	height := 10

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
	_ = png.Encode(part, img)
	request := httptest.NewRequest(http.MethodPost, "/", body)
	request.Header.Add(http_handler.ContentTypeHeader, writer.FormDataContentType())

	return request
}
