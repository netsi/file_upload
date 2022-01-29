package http_handler

import "net/http"

// RequestHandler receives http.Request and returns a Response.
type RequestHandler func(r *http.Request) *Response
