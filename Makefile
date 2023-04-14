.PHONY: build
build:
	go build -v ./cmd/api && ./api

.DEFAULT_GOAL:= build