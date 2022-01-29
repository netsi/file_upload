SHELL := /bin/bash

define USAGE
File upload

Commands:
	help				Shows this.
	start				Builds and starts the application in a docker container on port 8080
	stop 				Stops the application
endef

help:
	@echo "${USAGE}"

start:
	docker build -t file_upload -f ./build/dev/Dockerfile .
	docker run -d --rm --name file_upload -v uploads:/uploads -p 8080:8080 file_upload:latest

stop:
	docker stop file_upload