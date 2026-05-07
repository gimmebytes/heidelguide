# 🏰 heidelguide

A tourism guide web app for the city of Heidelberg — built live on stage as a demo for the [trade/off Summit](https://tradeoff.sh) MasterClass **"Agentic AI im Realitätscheck"** (2026).

## What is this?

heidelguide is a server-rendered web application that showcases Heidelberg's landmarks, history, and points of interest. It serves as a practical example of building a real application with AI-assisted development tooling — specifically [Kiro](https://kiro.dev), an AI-powered IDE by AWS.

The app is intentionally kept simple and self-contained: one binary, one database file, no external services. Start it, open the browser, explore Heidelberg.

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go |
| Routing | [chi](https://github.com/go-chi/chi) |
| Database | SQLite (via [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)) |
| Templates | Go `html/template` |
| Interactivity | [HTMX](https://htmx.org) + [Alpine.js](https://alpinejs.dev) |
| Styling | [Tailwind CSS](https://tailwindcss.com) |
| Hot-Reload | [Air](https://github.com/air-verse/air) |

All frontend libraries are served locally — no CDN dependencies at runtime.

## Built with Kiro

This project was developed using [Kiro](https://kiro.dev), an AI-powered development environment. Kiro's spec-driven workflow was used to plan, design, and implement features through structured requirements, design documents, and task lists. The specs live in `.kiro/specs/` if you're curious about the process.

## Getting Started

### Prerequisites

- Go 1.22+
- [Air](https://github.com/air-verse/air) (for hot-reload, optional)

### Run the app

```bash
make run
```

The server starts on [http://localhost:8080](http://localhost:8080).

### Development with hot-reload

Install Air once:

```bash
go install github.com/air-verse/air@latest
```

Then start the dev server:

```bash
make dev
```

Air watches `.go`, `.html`, `.css`, and `.js` files in `cmd/`, `internal/`, `templates/`, and `static/`. Every save triggers an automatic rebuild and restart.

## Available Commands

| Command | Description |
|---------|-------------|
| `make run` | Start the application |
| `make dev` | Start with hot-reload (requires Air) |
| `make build` | Compile the binary to `bin/server` |
| `make test` | Run unit tests |
| `make test-visual` | Run Playwright visual regression tests |
| `make lint` | Run golangci-lint |
| `make deps` | Download frontend dependencies |
| `make docker` | Build the Docker image |
| `make clean` | Remove build artifacts |

## Docker

```bash
make docker
docker run -p 8080:8080 heidelberg-guide
```

No compose file needed — the image is fully self-contained.

## License

[MIT](LICENSE)
