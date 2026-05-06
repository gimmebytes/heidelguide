---
inclusion: always
---

# Dev Setup

## Makefile
- All core commands must be accessible via a central `Makefile` at the project root
- Common targets to include:
  - `make run` — start the application locally
  - `make build` — compile the Go binary
  - `make test` — run unit tests
  - `make test-visual` — run Playwright visual regression checks
  - `make lint` — run linters
  - `make deps` — download/fetch frontend assets (HTMX, Alpine.js, Tailwind CSS)
  - `make clean` — remove build artifacts

## Local Development
- The app must run locally with zero external service dependencies
- All frontend libraries (HTMX, Alpine.js) are vendored/downloaded locally
- Use `go mod` for Go dependency management
