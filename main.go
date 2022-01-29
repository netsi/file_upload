package main

import (
	"file_upload/api/v1/file"
	"file_upload/internal/env"
	"file_upload/pkg/http-handler"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	publicURLKey   = "PUBLIC_URL"
	httpServerPort = 8080
	defaultHost    = "http://localhost"

	defaultReadTimeout  = 60 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

func main() {
	defaultPublicURL := fmt.Sprintf("%s:%d", defaultHost, httpServerPort)
	publicURL := env.LookupWithDefault(publicURLKey, defaultPublicURL)

	fileHandler := file.NewFileHandler(publicURL)

	exactPathRequestHandler := http_handler.NewHandler()
	exactPathRequestHandler.POST(fileHandler.HandleUpload)

	pathParamRequestHandler := http_handler.NewHandler()
	pathParamRequestHandler.GET(fileHandler.HandleGetFile)

	http.HandleFunc(file.V1FileAPIPath, exactPathRequestHandler.Handle)
	http.HandleFunc(fmt.Sprintf("%s/", file.V1FileAPIPath), pathParamRequestHandler.Handle)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", httpServerPort),
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start the http server on port %d with error %s", httpServerPort, err.Error())
	}
}
