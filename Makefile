.PHONY: run build test deps clean docker lint

# Application
APP_NAME := heidelberg-guide
BINARY := bin/server
CMD := ./cmd/server

# URLs for frontend dependencies
HTMX_URL := https://unpkg.com/htmx.org/dist/htmx.min.js
ALPINE_URL := https://unpkg.com/alpinejs/dist/cdn.min.js
TAILWIND_URL := https://cdn.jsdelivr.net/npm/tailwindcss@3/dist/tailwind.min.css

## run: Compile and start the application
run:
	go run $(CMD)

## build: Compile the Go binary to bin/server
build:
	@mkdir -p bin
	go build -o $(BINARY) $(CMD)

## test: Run all Go unit tests
test:
	go test ./...

## lint: Run golangci-lint
lint:
	golangci-lint run

## deps: Download frontend dependencies (HTMX, Alpine.js, Tailwind CSS)
deps:
	@mkdir -p static/js static/css
	curl -sL -o static/js/htmx.min.js $(HTMX_URL)
	curl -sL -o static/js/alpine.min.js $(ALPINE_URL)
	curl -sL -o static/css/tailwind.min.css $(TAILWIND_URL)

## clean: Remove build artifacts and database
clean:
	rm -rf bin/
	rm -f heidelberg.db

## docker: Build the Docker image
docker:
	docker build -t $(APP_NAME) .
