all: run
run:
	go build -v ./cmd/api && ./api