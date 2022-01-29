package http_handler

import (
	"log"
	"net/http"
	"sync"
)

type handler struct {
	methodHandlers sync.Map
}

// NewHandler initializes the http handler.
func NewHandler() *handler {
	return &handler{
		methodHandlers: sync.Map{},
	}
}

// POST adds a POST RequestHandler.
func (h *handler) POST(methodHandler RequestHandler) {
	h.methodHandlers.Store(http.MethodPost, methodHandler)
}

// GET adds a GET RequestHandler.
func (h *handler) GET(requestHandler RequestHandler) {
	h.methodHandlers.Store(http.MethodGet, requestHandler)
}

// Handle handles an HTTP request by using the configured Request Handlers for an HTTP Method.
// If the RequestHandler is not found for this HTTP Method then 404 is returned.
// If the RequestHandler doesn't return a response then 204 is returned.
// Otherwise returns the Headers, StatusCode and the response that the RequestHandler returned.
func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	handlerItem, found := h.methodHandlers.Load(r.Method)
	if !found {
		log.Print("no handler found for this method")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	methodHandler, ok := handlerItem.(RequestHandler)
	if !ok {
		log.Print("method handler was not of type RequestHandler")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := methodHandler(r)
	if response == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for headerKey, headerVal := range response.Headers {
		w.Header().Set(headerKey, headerVal)
	}

	w.WriteHeader(response.StatusCode)
	_, err := w.Write(response.Response)
	if err != nil {
		log.Printf("failed to encode the response object with error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
