---
inclusion: always
---

# Tech Stack

## Backend
- **Language**: Go (Golang)
- **Routing**: Standard library `net/http` or `chi` router
- **Database**: SQLite — lightweight, file-based, no external database server required
- **Principles**: Use Go best practices, prefer the standard library over third-party packages unless there's a clear functional need

## Frontend
- **Interactivity**: HTMX for server-driven interactions, Alpine.js for lightweight client-side state
- **Styling**: Tailwind CSS
- **No React, no TypeScript** — keep the frontend simple and server-rendered
- **Local assets**: HTMX and Alpine.js must be downloaded/fetched and served locally — no CDN dependencies at runtime

## Static Assets
- All static assets (images, JS libraries, CSS files) are served by the Go application itself
- Use Go's `embed` package or a dedicated `static/` folder served via `http.FileServer`
- Directory structure: `static/js/`, `static/css/`, `static/img/` (or similar)
- No external asset hosting or CDN — everything is self-contained

## General
- Follow the KISS principle throughout
- Do not introduce dependencies unless they are strictly required for functionality
- The application must be fully executable locally without external service dependencies
