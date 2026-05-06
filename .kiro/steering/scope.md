---
inclusion: always
---

# Scope & Mindset

## MVP / Demo-First
- This is a conference demo application — prioritize showability over production readiness
- Favor "good enough" over perfect: skip auth, rate limiting, and other production concerns unless they serve the demo
- The app should be startable with a single `make run` and immediately show meaningful, visually compelling content
- Keep the data model simple and focused on what's needed to demonstrate the concept
- Do not over-engineer — if it looks good and works on stage, it's done

## Distribution
- The app must be distributable as a small, self-contained Docker image
- Use multi-stage builds to keep the image minimal (Go binary + static assets + SQLite DB)
- A `make docker` target should build the image
- `docker run` should be all that's needed to start the app — no external dependencies, no compose files required
