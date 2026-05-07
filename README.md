# heidelguide
Demo app for trade/off Summit MasterClass "Agentic AI im Realitätscheck" (2026)

## Development with Hot-Reload

This project uses [Air](https://github.com/air-verse/air) for live-reloading during development. Air watches for file changes, rebuilds the Go binary, and restarts the server automatically.

### Install Air

```bash
go install github.com/air-verse/air@latest
```

### Start development server

```bash
make dev
```

### Watched files and directories

Air monitors the following for changes:

- **Extensions**: `.go`, `.html`, `.css`, `.js`
- **Directories**: `cmd/`, `internal/`, `templates/`, `static/`

Changes to any of these trigger an automatic rebuild and restart. The `tmp/`, `tests/`, `node_modules/`, `.git/`, and `vendor/` directories are excluded.