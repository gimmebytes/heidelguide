# Requirements Document

## Introduction

This document defines the requirements for the foundational infrastructure of the Heidelberg Tourism Guide demo application. The application is a conference demo that showcases a tourism guide for the city of Heidelberg, Germany. This spec covers the project skeleton, build system, database setup, template rendering, seed data, and containerization — everything needed to run `make run` and see a working app with meaningful content.

## Glossary

- **App**: The Heidelberg Tourism Guide Go application
- **Server**: The HTTP server binary located in `cmd/server/`
- **Router**: The chi HTTP router that dispatches requests to handlers
- **Store**: The SQLite persistence layer in `internal/store/`
- **Handler**: HTTP request handlers in `internal/handler/`
- **Template_Engine**: Go's `html/template` package used to render HTML pages
- **Static_File_Server**: The HTTP file server that serves assets from the `static/` directory
- **Landmark**: A point of interest in Heidelberg stored in the database (e.g., castle, bridge), with multilingual content
- **Seed_Data**: Pre-populated landmark records inserted into the database on first run, including real historic information and images
- **Locale**: A language identifier (e.g., "de", "en") used to select translated content
- **i18n**: Internationalization — the system's ability to serve content in multiple languages
- **Makefile**: The central build automation file at the project root
- **Dockerfile**: The multi-stage container build definition

## Requirements

### Requirement 1: Project Structure

**User Story:** As a developer, I want the project to follow the canonical Go layout, so that the codebase is immediately navigable and idiomatic.

#### Acceptance Criteria

1. THE App SHALL organize source code into `cmd/server/`, `internal/handler/`, `internal/model/`, and `internal/store/` directories
2. THE App SHALL serve frontend assets from a `static/` directory containing `js/`, `css/`, and `img/` subdirectories
3. THE App SHALL store HTML templates in a `templates/` directory at the project root
4. THE App SHALL include a `go.mod` file at the project root defining the Go module
5. THE App SHALL include a `Makefile` at the project root as the central command interface

### Requirement 2: Go Module and Dependencies

**User Story:** As a developer, I want the Go module properly initialized with minimal dependencies, so that the project builds reliably with standard tooling.

#### Acceptance Criteria

1. THE App SHALL declare a Go module in `go.mod` with a valid module path
2. THE App SHALL use the chi router package as the sole third-party routing dependency
3. THE App SHALL use a SQLite driver package as the sole third-party database dependency
4. THE App SHALL compile successfully using `go build ./...` after dependency resolution

### Requirement 3: HTTP Server and Routing

**User Story:** As a developer, I want a configured HTTP server with chi routing, so that the app can serve pages and static assets on startup.

#### Acceptance Criteria

1. WHEN the Server starts, THE Server SHALL listen on a configurable port defaulting to 8080
2. WHEN a request arrives at the root path, THE Router SHALL route the request to the landing page handler
3. WHEN a request arrives at `/landmarks/:id`, THE Router SHALL route the request to the landmark detail handler
4. WHEN a request arrives at the static asset path prefix, THE Router SHALL route the request to the Static_File_Server
5. THE Router SHALL return HTTP 404 for undefined routes
6. WHEN the Server starts successfully, THE Server SHALL log the listening address to standard output

### Requirement 4: Static File Serving

**User Story:** As a user, I want all frontend assets served locally by the app, so that the app works without any external network dependencies.

#### Acceptance Criteria

1. THE Static_File_Server SHALL serve files from the `static/` directory under the `/static/` URL path
2. THE App SHALL include HTMX JavaScript library in `static/js/`
3. THE App SHALL include Alpine.js JavaScript library in `static/js/`
4. THE App SHALL include Tailwind CSS stylesheet in `static/css/`
5. WHEN a requested static file does not exist, THE Static_File_Server SHALL return HTTP 404

### Requirement 5: SQLite Database Setup

**User Story:** As a developer, I want a SQLite database with landmarks and translations tables created automatically, so that the app has persistent multilingual storage without manual setup.

#### Acceptance Criteria

1. WHEN the App starts, THE Store SHALL create the SQLite database file if the file does not exist
2. WHEN the Store initializes, THE Store SHALL create a `landmarks` table with columns for id, latitude, longitude, image_filename, year_built, and year_destroyed (nullable)
3. WHEN the Store initializes, THE Store SHALL create a `landmark_translations` table with columns for id, landmark_id, locale, name, description, and history
4. WHEN the Store initializes, THE Store SHALL use an auto-incrementing integer primary key for the landmark id column
5. THE `landmark_translations` table SHALL have a unique constraint on (landmark_id, locale)
6. IF the database file already exists with the correct schema, THEN THE Store SHALL reuse the existing database without modification
7. IF the database cannot be opened or created, THEN THE App SHALL exit with a descriptive error message

### Requirement 6: Seed Data

**User Story:** As a user, I want to see real Heidelberg landmarks with authentic historic information and images immediately on first run, so that the app demonstrates meaningful, locally recognizable content.

#### Acceptance Criteria

1. WHEN the Store initializes with an empty landmarks table, THE Store SHALL insert seed data for at least 8 Heidelberg landmarks
2. THE Seed_Data SHALL include Heidelberg Castle (Heidelberger Schloss), Old Bridge (Alte Brücke / Karl-Theodor-Brücke), Philosophers' Walk (Philosophenweg), Church of the Holy Spirit (Heiliggeistkirche), Student Prison (Studentenkarzer), University Library (Universitätsbibliothek), Königstuhl, and Neckar Meadow (Neckarwiese) as landmarks
3. THE Seed_Data SHALL provide for each landmark: latitude, longitude, year_built, year_destroyed (if applicable), and an image file
4. THE Seed_Data SHALL provide translations in both "de" and "en" locales for each landmark, including name, description, and a brief history text
5. THE Seed_Data history text SHALL contain real, factually accurate historic information (e.g., "Erbaut im 13. Jahrhundert, im Pfälzischen Erbfolgekrieg 1693 zerstört")
6. THE Seed_Data SHALL include a representative image for each landmark stored in `static/img/landmarks/`
7. IF the landmarks table already contains data, THEN THE Store SHALL skip seed data insertion

### Requirement 6a: Internationalization (i18n)

**User Story:** As a user, I want to switch the app language between German and English, so that both local and international visitors can use the guide.

#### Acceptance Criteria

1. THE App SHALL use German ("de") as the default language
2. THE App SHALL support English ("en") as an alternative language
3. WHEN a user switches language, THE App SHALL display all landmark content (name, description, history) in the selected language
4. THE App SHALL provide a language switcher UI element visible on all pages
5. THE App SHALL persist the selected language preference for the duration of the browser session (e.g., via cookie or URL parameter)
6. THE App SHALL serve UI labels (navigation, headings, buttons) in the selected language using a simple translation mechanism (e.g., a map of key-value pairs per locale)

### Requirement 7: HTML Template Rendering

**User Story:** As a user, I want to see a rendered landing page with landmark content, so that the app displays visually compelling content immediately.

#### Acceptance Criteria

1. THE Template_Engine SHALL parse all HTML templates from the `templates/` directory on startup
2. WHEN the landing page is requested, THE Handler SHALL render a template displaying all landmarks from the database
3. THE landing page template SHALL include HTMX and Alpine.js script references pointing to locally served files
4. THE landing page template SHALL include the Tailwind CSS stylesheet reference pointing to the locally served file
5. IF a template fails to parse, THEN THE App SHALL exit with a descriptive error message indicating the template file and parse error

### Requirement 8: Makefile Build System

**User Story:** As a developer, I want a comprehensive Makefile, so that all common development tasks are accessible through a single consistent interface.

#### Acceptance Criteria

1. THE Makefile SHALL provide a `run` target that compiles and starts the App
2. THE Makefile SHALL provide a `build` target that compiles the Go binary to a `bin/` directory
3. THE Makefile SHALL provide a `test` target that runs all Go unit tests
4. THE Makefile SHALL provide a `deps` target that downloads HTMX, Alpine.js, and Tailwind CSS to the `static/` directory
5. THE Makefile SHALL provide a `clean` target that removes build artifacts and the `bin/` directory
6. THE Makefile SHALL provide a `docker` target that builds the Docker container image
7. THE Makefile SHALL provide a `lint` target that runs Go linting tools

### Requirement 9: Dockerfile and Containerization

**User Story:** As a developer, I want a multi-stage Dockerfile producing a minimal image, so that the app is easily distributable and runnable anywhere.

#### Acceptance Criteria

1. THE Dockerfile SHALL use a multi-stage build with a Go builder stage and a minimal runtime stage
2. THE Dockerfile SHALL produce a final image containing only the compiled binary, static assets, templates, and SQLite database
3. WHEN the container starts, THE App SHALL listen on port 8080 and serve the landing page
4. THE Dockerfile SHALL use a distroless or Alpine-based runtime image to minimize image size
5. THE Dockerfile SHALL copy the SQLite database file into the container image so the app starts with seed data

### Requirement 10: Landing Page Content

**User Story:** As a conference attendee viewing the demo, I want to see an attractive landing page with Heidelberg landmarks in German by default, so that the app immediately communicates its purpose and looks polished.

#### Acceptance Criteria

1. THE landing page SHALL display a header with the application title "Heidelberg Guide"
2. THE landing page SHALL display each landmark as a card showing the landmark name, description, and a thumbnail image
3. THE landing page SHALL display content in the currently selected language (German by default)
4. THE landing page SHALL include a language switcher to toggle between German and English
5. THE landing page SHALL use Tailwind CSS utility classes for styling with a warm, inviting color palette
6. THE landing page SHALL be responsive and render correctly on both desktop and mobile viewport widths
7. THE landing page SHALL load without any external network requests

### Requirement 12: Structured Logging and Access Logs

**User Story:** As a developer, I want structured application logs and HTTP access logs printed to the console, so that I can observe request traffic and application behavior side by side during development and demos.

#### Acceptance Criteria

1. THE App SHALL use Go's standard `log/slog` package for all application-level logging (startup, errors, informational messages)
2. THE App SHALL output structured log entries in text format to standard output, including timestamp, level, and message fields
3. THE App SHALL include a chi middleware that logs every HTTP request with method, path, status code, and response duration
4. THE App SHALL log access entries at INFO level using the same `slog` logger so that access logs and app logs appear interleaved in a single console stream
5. THE App SHALL log application startup events (DB opened, migrations run, seed complete, server listening) at INFO level
6. THE App SHALL log application errors (failed DB queries, template render errors) at ERROR level
7. THE App SHALL NOT use the legacy `log` package — all logging goes through `slog`

### Requirement 11: Landmark Detail Page

**User Story:** As a user, I want to click on a landmark card and see a detailed page with full information, so that I can learn more about a specific point of interest.

#### Acceptance Criteria

1. WHEN a user clicks a landmark card on the landing page, THE App SHALL navigate to a detail page at `/landmarks/:id`
2. THE detail page SHALL display the landmark's full name, description, history text, year_built, year_destroyed (if applicable), and a full-size image
3. THE detail page SHALL display content in the currently selected language
4. THE detail page SHALL include a back/home navigation link to return to the landing page
5. IF the requested landmark ID does not exist, THE Handler SHALL return HTTP 404 with a user-friendly error page
6. THE detail page SHALL use the same Tailwind CSS styling and layout conventions as the landing page
7. THE detail page SHALL include the language switcher consistent with the landing page
