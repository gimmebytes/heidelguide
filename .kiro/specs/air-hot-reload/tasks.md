# Tasks

## Task 1: Create Air configuration file

- [x] 1.1 Create `.air.toml` at the project root with the `[build]` section: set `cmd` to `go build -o ./tmp/main ./cmd/server`, set `bin` to `./tmp/main`, set `include_ext` to `["go", "html", "css", "js"]`, set `include_dir` to `["cmd", "internal", "templates", "static"]`, and set `exclude_dir` to `["tmp", "tests", "node_modules", ".git", "vendor"]`
- [x] 1.2 Add `[log]` and `[misc]` sections with sensible defaults (clean_on_exit = true)

## Task 2: Update .gitignore

- [x] 2.1 Add `tmp/` entry to `.gitignore` to exclude Air's build output directory

## Task 3: Add Makefile dev target

- [~] 3.1 Add a `dev` phony target to the Makefile that runs `air`, with a help comment `## dev: Start the application with hot-reload (requires Air)`

## Task 4: Update README with hot-reload documentation

- [~] 4.1 Add a "Development with Hot-Reload" section to README.md documenting Air installation (`go install github.com/air-verse/air@latest`), the `make dev` command, and which file types/directories are watched
