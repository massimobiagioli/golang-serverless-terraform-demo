.PHONY: build start-local format test cover terraform-create terraform-destroy terraform-format help
.DEFAULT_GOAL := help

build: # Build the package
	go build -o bin/server cmd/server/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/bootstrap cmd/lambda_main/main.go

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

terraform-create: build # Create infrastructure
	cd infra && terraform init && terraform plan && terraform apply -auto-approve

terraform-destroy: # Destroy infrastructure
	cd infra && terraform destroy -auto-approve

terraform-format: # Format terraform files
	cd infra && terraform fmt

help: # make help
	@awk 'BEGIN {FS = ":.*#"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?#/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
