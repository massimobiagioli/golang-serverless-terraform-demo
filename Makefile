.PHONY: build start-local format test cover help
.DEFAULT_GOAL := help

build: # Build the package
	go build -o bin/server cmd/server/main.go

start-local: # Start the server locally
	go run cmd/server/main.go

format: # Format the code
	go fmt ./...

test: # Run tests
	@if [ -z "$(filter)" ]; then \
        go test ./... -v; \
    else \
        go test -run $(filter) ./... -v; \
    fi

cover: # Run tests with coverage
	go test ./... -cover	

help: # make help
	@awk 'BEGIN {FS = ":.*#"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?#/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
