SHELL := /bin/bash

.PHONY: test api docs

test:
	echo "running tests"
	go test ./... -v

api:
	echo "running api"
	go run main.go

docs:
	echo "starting swagger documentation"
	swagger serve -F=swagger swagger.yaml