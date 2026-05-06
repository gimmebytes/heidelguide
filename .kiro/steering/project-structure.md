---
inclusion: always
---

# Project Structure

## Layout
Follow the canonical Go project layout with a clean, minimal root directory.

```
├── cmd/                  # Application entry points (main packages)
│   └── server/           # Main server binary
├── internal/             # Private application code (not importable by other projects)
│   ├── handler/          # HTTP handlers
│   ├── model/            # Domain models / structs
│   ├── store/            # Database / persistence layer (SQLite)
│   └── ...               # Additional packages as needed
├── static/               # Static assets served by the app
│   ├── js/               # HTMX, Alpine.js, and other JS
│   ├── css/              # Tailwind output and other styles
│   └── img/              # Images and icons
├── templates/            # Go HTML templates
├── docs/                 # All application documentation
├── Makefile              # Central command interface
├── go.mod
├── go.sum
└── README.md
```

## Principles
- **Minimal root** — only essential config files (Makefile, go.mod, go.sum, README) live at the root
- **`cmd/`** — each subdirectory is a separate main package / binary entry point
- **`internal/`** — all private application logic; enforces encapsulation via Go's import rules
- **`static/`** — all frontend assets; served directly by the Go app
- **`templates/`** — server-rendered HTML templates (Go `html/template`)
- **`docs/`** — project documentation, architecture decisions, guides
- Do not create packages prematurely — add structure as complexity grows
