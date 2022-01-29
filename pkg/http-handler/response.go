package http_handler

import (
	"encoding/json"
	"log"
)

const (
	ContentTypeHeader = "Content-Type"
	JsonContentType   = "application/json"
	StreamContent     = "application/octet-stream"
)

// Response returned by the RequestHandler.
type Response struct {
	Headers    map[string]string
	StatusCode int
	Response   []byte
}

// JSON marshals the resp and returns Response object with json Content-Type.
func JSON(statusCode int, resp interface{}) *Response {
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("could not marshal the response object %s", err.Error())
	}

	return &Response{
		Headers:    map[string]string{ContentTypeHeader: JsonContentType},
		StatusCode: statusCode,
		Response:   jsonResponse,
	}
}

// Image returns the byte slice with octet-stream Content-Type.
func Image(statusCode int, bytes []byte) *Response {
	return &Response{
		Headers:    map[string]string{ContentTypeHeader: StreamContent},
		StatusCode: statusCode,
		Response:   bytes,
	}
}
