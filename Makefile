.PHONY: run dev build test test-visual deps clean docker lint deploy

# Application
APP_NAME := heidelguide
BINARY := bin/server
CMD := ./cmd/server

# Deployment
DEPLOY_HOST ?= gimmebytes-apps-prod
CONTAINER_NAME := $(APP_NAME)

# URLs for frontend dependencies
HTMX_URL := https://unpkg.com/htmx.org/dist/htmx.min.js
ALPINE_URL := https://unpkg.com/alpinejs/dist/cdn.min.js
TAILWIND_URL := https://cdn.tailwindcss.com/3.4.17

## run: Compile and start the application
run:
	go run $(CMD)

## dev: Start the application with hot-reload (requires Air), opens browser when ready
dev:
	@(until curl -s http://localhost:8080 > /dev/null 2>&1; do sleep 0.2; done; open http://localhost:8080) &
	air

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
	curl -sL -o static/js/tailwind-cdn.js $(TAILWIND_URL)

## clean: Remove build artifacts and database
clean:
	rm -rf bin/
	rm -f heidelberg.db

## test-visual: Run Playwright visual regression tests
test-visual:
	cd tests && npx playwright test

## docker: Build the Docker image
docker:
	docker build -t $(APP_NAME) .

## deploy: Build Docker image and deploy to remote host via SSH
deploy: docker
	docker save $(APP_NAME):latest | gzip | ssh $(DEPLOY_HOST) "gunzip | docker load"
	ssh $(DEPLOY_HOST) "docker stop $(CONTAINER_NAME) 2>/dev/null || true"
	ssh $(DEPLOY_HOST) "docker rm $(CONTAINER_NAME) 2>/dev/null || true"
	ssh $(DEPLOY_HOST) "docker run -d --restart unless-stopped --name $(CONTAINER_NAME) -p 8080:8080 $(APP_NAME):latest"
