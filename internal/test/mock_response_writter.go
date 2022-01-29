package test

import "net/http"

type MockResponseWritter struct {
	Headers    http.Header
	Output     []byte
	StatusCode int
}

func NewMockResponseWritter() *MockResponseWritter {
	return &MockResponseWritter{
		Headers: map[string][]string{},
	}
}

func (m *MockResponseWritter) Header() http.Header {
	return m.Headers
}

func (m *MockResponseWritter) Write(bytes []byte) (int, error) {
	m.Output = bytes
	return len(bytes), nil
}

func (m *MockResponseWritter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}
