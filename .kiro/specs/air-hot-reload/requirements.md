# Requirements Document

## Introduction

This feature integrates [Air](https://github.com/air-verse/air) as a development-time hot-reload tool for the Heidelberg tourism demo application. Air watches for file changes (Go source, HTML templates, static assets) and automatically rebuilds and restarts the server, enabling a seamless live-coding experience during conference demos. The goal is zero-friction iteration: save a file, see the result in the browser immediately without manual restarts.

## Glossary

- **Air**: A live-reloading tool for Go applications that watches for file changes, rebuilds the binary, and restarts the process automatically.
- **Air_Configuration**: The `.air.toml` configuration file that defines Air's watch paths, build command, exclusions, and reload behavior.
- **Developer**: A person running the application locally for development or live demo purposes.
- **Build_Command**: The shell command Air executes to compile the Go application after detecting a change.
- **Watch_Path**: A directory that Air monitors for file changes.
- **Exclude_Path**: A directory or pattern that Air ignores when monitoring for changes.
- **Makefile**: The project's central command interface for common development tasks.

## Requirements

### Requirement 1: Air Configuration File

**User Story:** As a developer, I want a pre-configured Air configuration file in the project, so that I can start hot-reloading without manual setup.

#### Acceptance Criteria

1. THE Air_Configuration SHALL exist as a `.air.toml` file at the project root.
2. THE Air_Configuration SHALL specify `./cmd/server` as the build target using `go build -o ./tmp/main ./cmd/server`.
3. THE Air_Configuration SHALL store the compiled binary in the `./tmp/` directory.
4. THE Air_Configuration SHALL set the build output binary name to `main`.

### Requirement 2: Go Source File Watching

**User Story:** As a developer, I want Air to detect changes to Go source files, so that the server automatically rebuilds and restarts when I modify backend code.

#### Acceptance Criteria

1. THE Air_Configuration SHALL include `.go` in the list of watched file extensions.
2. WHEN a `.go` file in the project is modified, THE Air SHALL trigger a rebuild and restart of the server.
3. THE Air_Configuration SHALL watch the `cmd/` and `internal/` directories for Go file changes.

### Requirement 3: Template File Watching

**User Story:** As a developer, I want Air to detect changes to HTML template files, so that template edits trigger a server restart and reflect immediately.

#### Acceptance Criteria

1. THE Air_Configuration SHALL include `.html` in the list of watched file extensions.
2. THE Air_Configuration SHALL watch the `templates/` directory for template changes.
3. WHEN an HTML template file is modified, THE Air SHALL trigger a rebuild and restart of the server.

### Requirement 4: Static Asset Watching

**User Story:** As a developer, I want Air to detect changes to static assets (CSS, JS), so that style and script modifications are picked up without manual restarts.

#### Acceptance Criteria

1. THE Air_Configuration SHALL include `.css` and `.js` in the list of watched file extensions.
2. THE Air_Configuration SHALL watch the `static/` directory for asset changes.
3. WHEN a static asset file is modified, THE Air SHALL trigger a rebuild and restart of the server.

### Requirement 5: Exclude Non-Source Directories

**User Story:** As a developer, I want Air to ignore generated files, dependencies, and build artifacts, so that unnecessary rebuilds are avoided.

#### Acceptance Criteria

1. THE Air_Configuration SHALL exclude the `tmp/` directory from watching.
2. THE Air_Configuration SHALL exclude the `tests/` directory from watching.
3. THE Air_Configuration SHALL exclude the `node_modules/` directory pattern from watching.
4. THE Air_Configuration SHALL exclude the `.git/` directory from watching.
5. THE Air_Configuration SHALL exclude the `vendor/` directory from watching.

### Requirement 6: Makefile Integration

**User Story:** As a developer, I want a `make dev` target that starts Air, so that hot-reload development follows the same Makefile-driven workflow as other project commands.

#### Acceptance Criteria

1. THE Makefile SHALL provide a `dev` target that executes the `air` command.
2. WHEN the Developer runs `make dev`, THE Makefile SHALL start Air with the project's `.air.toml` configuration.
3. THE Makefile SHALL include a help comment for the `dev` target describing its purpose.

### Requirement 7: Git Ignore Air Artifacts

**User Story:** As a developer, I want Air's temporary build artifacts excluded from version control, so that generated files do not pollute the repository.

#### Acceptance Criteria

1. THE `.gitignore` file SHALL include the `tmp/` directory used by Air for build output.

### Requirement 8: Development Documentation

**User Story:** As a developer, I want clear instructions on how to install Air and use the hot-reload workflow, so that anyone cloning the project can get started quickly.

#### Acceptance Criteria

1. THE README SHALL document how to install Air (via `go install github.com/air-verse/air@latest`).
2. THE README SHALL document the `make dev` command as the way to start the application with hot-reload.
3. THE README SHALL explain which file types and directories are watched by Air.
