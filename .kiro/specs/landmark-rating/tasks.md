# Tasks

## Task 1: Add ratings table migration

- **File**: `internal/store/migrations.go`
- **Requirements**: REQ-5
- **Description**: Add the `CREATE TABLE IF NOT EXISTS ratings (...)` statement to the `migrations` slice with the schema defined in design.md (id, landmark_id, device_id, score with CHECK constraint, created_at, updated_at, UNIQUE on landmark_id+device_id).

## Task 2: Implement store rating functions

- **File**: `internal/store/ratings.go` (new)
- **Requirements**: REQ-2, REQ-3, REQ-5
- **Description**: Create `RatingSummary` struct and implement `GetRatingSummary(landmarkID, deviceID)` and `UpsertRating(landmarkID, deviceID, score)` using the SQL from design.md.

## Task 3: Add `seq` template helper function

- **File**: `internal/handler/template_funcs.go`
- **Requirements**: REQ-6
- **Description**: Add a `seq` function to the template FuncMap that generates a slice of integers from start to end (inclusive), needed for iterating 1–5 stars in the template.

## Task 4: Create rating handler

- **File**: `internal/handler/rating.go` (new)
- **Requirements**: REQ-2, REQ-3, REQ-4
- **Description**: Implement `SubmitRating` handler — parse landmark ID, read X-Device-ID header, parse score from form body, validate (1–5), call store upsert, query updated summary, render partial template response.

## Task 5: Extend PageData and detail handler

- **File**: `internal/handler/handler.go`, `internal/handler/detail.go`
- **Requirements**: REQ-3, REQ-4
- **Description**: Add `Rating *store.RatingSummary` field to PageData (or use the store type). Update the Detail handler to read `X-Device-ID` from request header, call `GetRatingSummary`, and pass data to template.

## Task 6: Create star widget partial template

- **File**: `templates/partials/rating_widget.html` (new)
- **Requirements**: REQ-2, REQ-3, REQ-4, REQ-6
- **Description**: Build the Alpine.js + HTMX star widget as described in design.md. 5 star buttons with hover state, click submission, display of average and count, and the user's own rating highlighted.

## Task 7: Integrate widget into detail page

- **File**: `templates/detail.html`
- **Requirements**: REQ-3, REQ-6
- **Description**: Include the rating widget partial in the detail template, placed below the metadata badges section.

## Task 8: Add device_id initialization script

- **File**: `templates/base.html`
- **Requirements**: REQ-1
- **Description**: Add inline `<script>` that generates and stores a UUID in localStorage if not already present, using `crypto.randomUUID()`.

## Task 9: Register rating route

- **File**: `cmd/server/main.go`
- **Requirements**: REQ-2
- **Description**: Add `r.Post("/api/landmarks/{id}/rating", h.SubmitRating)` to the chi router.

## Task 10: Write unit tests for store layer

- **File**: `internal/store/ratings_test.go` (new)
- **Requirements**: REQ-2, REQ-5
- **Description**: Test `UpsertRating` (insert + update) and `GetRatingSummary` (empty, single rating, multiple ratings, correct average calculation).
