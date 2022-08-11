.PHONY: help test compile
.DEFAULT_GOAL := help

test:
	go test -race -shuffle=on ./...

compile:
	protoc api/v1/*.proto \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative  \
	--proto_path=.

help:
	@grep -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST)