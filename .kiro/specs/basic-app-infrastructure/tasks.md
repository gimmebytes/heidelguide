# Implementation Plan: Basic App Infrastructure

## Overview

Build the foundational infrastructure for the Heidelberg Tourism Guide — a self-contained Go web application with chi routing, SQLite persistence, HTMX/Alpine.js interactivity, Tailwind CSS styling, and multilingual support (German/English). Tasks are ordered so each step produces a runnable increment, culminating in a fully functional demo app launchable via `make run`.

## Tasks

- [ ] 1. Initialize project structure and Go module
  - [-] 1.1 Create directory structure and initialize Go module
    - Create `cmd/server/`, `internal/handler/`, `internal/model/`, `internal/store/`, `internal/i18n/`, `static/js/`, `static/css/`, `static/img/landmarks/`, `templates/` directories
    - Initialize `go.mod` with a valid module path
    - Add chi router (`github.com/go-chi/chi/v5`) and SQLite driver (`modernc.org/sqlite`) dependencies
    - Create a minimal `cmd/server/main.go` that compiles and prints "starting server"
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 2.1, 2.2, 2.3, 2.4_

  - [~] 1.2 Create Makefile with all targets
    - Implement `run`, `build`, `test`, `deps`, `clean`, `docker`, `lint` targets
    - `deps` target downloads HTMX, Alpine.js, and Tailwind CSS standalone CLI output to `static/`
    - `build` compiles to `bin/server`
    - `clean` removes `bin/` and `heidelberg.db`
    - _Requirements: 1.5, 8.1, 8.2, 8.3, 8.4, 8.5, 8.6, 8.7_

- [ ] 2. Implement data layer (models, store, migrations, seed)
  - [~] 2.1 Define domain models
    - Create `internal/model/landmark.go` with `Landmark`, `LandmarkTranslation`, and `LandmarkWithTranslation` structs
    - _Requirements: 5.2, 5.3_

  - [~] 2.2 Implement SQLite store with migrations
    - Create `internal/store/store.go` with `Store` struct, `Open()`, `Close()`, `Migrate()` methods
    - Create `internal/store/migrations.go` with `CREATE TABLE IF NOT EXISTS` for `landmarks` and `landmark_translations`
    - Enforce unique constraint on `(landmark_id, locale)`
    - Auto-incrementing integer primary key on `landmarks.id`
    - Handle DB open/create errors with descriptive messages
    - _Requirements: 5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7_

  - [~] 2.3 Implement seed data
    - Create `internal/store/seed.go` with `Seed()` method
    - Insert 8 Heidelberg landmarks with real coordinates, year_built, year_destroyed, image filenames
    - Insert German and English translations with factually accurate name, description, and history text
    - Only seed when `landmarks` table is empty (check via `SELECT COUNT(*)`)
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5, 6.7_

  - [~] 2.4 Implement query methods
    - Create `internal/store/landmarks.go` with `ListLandmarks(locale)` and `GetLandmark(id, locale)` methods
    - Join `landmarks` and `landmark_translations` tables filtered by locale
    - _Requirements: 6a.3_

  - [ ]* 2.5 Write property tests for data layer
    - **Property 1: Unique constraint on (landmark_id, locale) rejects duplicates**
    - **Validates: Requirements 5.5**

  - [ ]* 2.6 Write property test for seed data completeness
    - **Property 2: Seed data completeness**
    - **Validates: Requirements 6.3, 6.4**

  - [ ]* 2.7 Write property test for locale-correct query results
    - **Property 3: Locale-correct query results**
    - **Validates: Requirements 6a.3**

  - [ ]* 2.8 Write unit tests for store layer
    - Test `Migrate()` creates tables correctly
    - Test `Seed()` inserts data only when table is empty
    - Test `ListLandmarks()` returns all landmarks for a given locale
    - Test `GetLandmark()` returns correct landmark or nil for missing ID
    - Use in-memory SQLite for tests
    - _Requirements: 5.1, 5.2, 5.6, 6.1, 6.7_

- [ ] 3. Implement i18n labels
  - [~] 3.1 Create i18n package with UI labels
    - Create `internal/i18n/i18n.go` with `Labels()` function returning `map[string]map[string]string`
    - Include all UI labels for "de" and "en" locales (app_title, nav_home, back, year_built, year_destroyed, language, description, history)
    - _Requirements: 6a.1, 6a.2, 6a.6_

  - [ ]* 3.2 Write property test for UI labels completeness
    - **Property 4: UI labels completeness**
    - **Validates: Requirements 6a.6**

- [~] 4. Checkpoint
  - Ensure all tests pass with `go test ./...`, ask the user if questions arise.

- [ ] 5. Implement HTTP server, routing, and handlers
  - [~] 5.1 Set up HTTP server with chi router
    - Wire `cmd/server/main.go` to open the store, run migrations, seed data, parse templates, and start the HTTP server on port 8080
    - Register routes: `GET /`, `GET /landmarks/{id}`, `POST /language`, `GET /static/*`
    - Serve static files via `http.FileServer`
    - Log listening address to stdout on successful start
    - Exit with descriptive error on startup failures (DB, templates, port)
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 4.1, 4.5_

  - [~] 5.2 Implement handler struct and landing page handler
    - Create `internal/handler/handler.go` with `Handler` struct holding store, templates, and labels
    - Create `internal/handler/landing.go` with `Landing()` method
    - Read locale from cookie (default "de"), query all landmarks, render landing template
    - _Requirements: 7.2, 10.1, 10.2, 10.3, 6a.1_

  - [~] 5.3 Implement detail page handler
    - Create `internal/handler/detail.go` with `Detail()` method
    - Parse landmark ID from URL, query landmark by ID and locale
    - Return 404 with error page for missing/invalid IDs
    - Render detail template with full landmark data
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5_

  - [~] 5.4 Implement language switching handler
    - Create `internal/handler/language.go` with `SwitchLanguage()` method
    - Accept POST with locale value, set `lang` cookie (HttpOnly, SameSite=Lax, Path="/")
    - Redirect back to referring page
    - _Requirements: 6a.3, 6a.4, 6a.5_

  - [ ]* 5.5 Write unit tests for handlers
    - Test landing handler returns 200 with landmark content
    - Test detail handler returns 200 for valid ID and 404 for invalid ID
    - Test language handler sets cookie and redirects
    - Use `httptest.NewRecorder` with a seeded test database
    - _Requirements: 3.2, 3.3, 3.5, 11.5_

- [ ] 6. Implement HTML templates
  - [~] 6.1 Create base layout template
    - Create `templates/base.html` with shared HTML structure (head, nav with app title, language switcher, footer)
    - Include `<script>` tags for HTMX and Alpine.js pointing to `/static/js/`
    - Include `<link>` tag for Tailwind CSS pointing to `/static/css/`
    - Define `{{block "content" .}}` for page-specific content
    - _Requirements: 7.1, 7.3, 7.4, 10.1, 10.4, 10.7_

  - [~] 6.2 Create landing page template
    - Create `templates/landing.html` defining the "content" block
    - Render landmark cards in a responsive grid with name, description snippet, and thumbnail image
    - Each card links to `/landmarks/{id}`
    - Use Tailwind CSS utility classes for warm, inviting styling
    - Responsive layout for desktop and mobile
    - _Requirements: 10.2, 10.3, 10.5, 10.6, 11.1_

  - [~] 6.3 Create detail page template
    - Create `templates/detail.html` defining the "content" block
    - Display full landmark name, description, history, year_built, year_destroyed (if applicable), and full-size image
    - Include back/home navigation link
    - Use consistent Tailwind styling with landing page
    - _Requirements: 11.2, 11.3, 11.4, 11.6, 11.7_

  - [~] 6.4 Create 404 error page template
    - Create `templates/404.html` with user-friendly error message
    - Include navigation back to home
    - _Requirements: 3.5, 11.5_

  - [ ]* 6.5 Write property test for landing page rendering
    - **Property 5: Landing page renders all landmarks**
    - **Validates: Requirements 10.2**

  - [ ]* 6.6 Write property test for no external network dependencies
    - **Property 6: No external network dependencies in rendered pages**
    - **Validates: Requirements 10.7**

  - [ ]* 6.7 Write property test for detail page rendering
    - **Property 7: Detail page renders all landmark fields**
    - **Validates: Requirements 11.2**

- [~] 7. Checkpoint
  - Ensure all tests pass with `go test ./...`, ask the user if questions arise.

- [ ] 8. Static assets and Dockerfile
  - [~] 8.1 Download and configure static assets
    - Run `make deps` to download HTMX, Alpine.js, and Tailwind CSS to `static/js/` and `static/css/`
    - Add placeholder landmark images to `static/img/landmarks/` (one per seeded landmark)
    - _Requirements: 4.2, 4.3, 4.4, 6.6_

  - [~] 8.2 Create Dockerfile
    - Implement multi-stage build: Go builder stage + Alpine runtime stage
    - Copy compiled binary, templates, and static assets into final image
    - Expose port 8080, set CMD to run the server
    - SQLite DB created at runtime via auto-migration and seeding
    - _Requirements: 9.1, 9.2, 9.3, 9.4, 9.5_

- [~] 9. Final checkpoint
  - Ensure all tests pass with `go test ./...`, verify `make run` starts the app and displays the landing page with landmarks, ask the user if questions arise.

## Notes

- Tasks marked with `*` are optional and can be skipped for faster MVP
- Each task references specific requirements for traceability
- Checkpoints ensure incremental validation
- Property tests use `pgregory.net/rapid` as specified in the design
- Unit tests use Go's standard `testing` package with in-memory SQLite
- The app should be fully functional after task 7 (templates + handlers wired together)
