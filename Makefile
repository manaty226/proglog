.PHONY: help test proto
.DEFAULT_GOAL := help

test:
	go test -race -shuffle=on ./...

proto:
	protoc api/v1/*.proto --go_out=. --go_opt=paths=source_relative --proto_path=.

help:
	@grep -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST)