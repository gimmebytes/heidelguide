# Tasks: UX Refinement

## Task 1: Database Schema & Seed Data

- [x] 1.1 Add new migration SQL statements: CREATE TABLE categories, CREATE TABLE category_translations, ALTER TABLE landmarks ADD COLUMN category_id, ALTER TABLE landmarks ADD COLUMN highlighted
- [x] 1.2 Update the `model` package: add `Category`, `CategoryTranslation` structs; extend `Landmark` struct with `CategoryID` and `Highlighted` fields; create `LandmarkView` struct
- [x] 1.3 Implement `SeedCategories()` in the store: insert 4 categories with translations for de/en, idempotent
- [x] 1.4 Implement `AssignDefaultCategories()` in the store: assign category_id and highlighted flag to existing landmarks
- [x] 1.5 Update application startup to call `SeedCategories()` and `AssignDefaultCategories()` after migrations

## Task 2: Store Layer Updates

- [x] 2.1 Update `ListLandmarks` query to JOIN categories and category_translations, return `[]model.LandmarkView`, order by highlighted DESC then id ASC
- [x] 2.2 Update `GetLandmark` query to JOIN categories and category_translations, return `*model.LandmarkView`
- [x] 2.3 Write unit tests for `ListLandmarks` (verify category data populated, highlight ordering) and `GetLandmark` (verify category data)

## Task 3: Handler & Template Infrastructure

- [x] 3.1 Register `categoryColorClass` template function in the handler's template FuncMap
- [x] 3.2 Update `PageData` struct to use `[]model.LandmarkView` and `*model.LandmarkView`
- [x] 3.3 Add new i18n labels: `learn_more`, `home`, `highlight` for both de and en locales
- [x] 3.4 Write unit test for `categoryColorClass` (all valid colors + unknown fallback)

## Task 4: Landing Page Template

- [x] 4.1 Replace the gradient hero section with a full-width scenic image hero (img tag + gradient overlay + text)
- [x] 4.2 Add hero image file to `static/img/` (scenic Heidelberg photograph)
- [x] 4.3 Update the landmark grid from 3-column to 4-column layout (lg:grid-cols-4, md:grid-cols-2)
- [x] 4.4 Redesign landmark cards: add category pill, highlight badge, HG monogram placeholder, "Learn more" link, updated styling

## Task 5: Detail Page Template

- [x] 5.1 Add breadcrumb navigation bar below the nav: "Home > Landmark Name" with Home as a link to /
- [x] 5.2 Add category pill and highlight badge display to the detail page content area
- [x] 5.3 Apply warmer color palette to detail page elements (consistent with landing page overhaul)

## Task 6: Color Palette & Base Template

- [x] 6.1 Update `base.html` nav and body colors to use the warmer, more vibrant palette
- [x] 6.2 Update footer styling to match the new palette
- [x] 6.3 Verify Tailwind classes used (rose, emerald, amber, violet) are available in the existing tailwind.min.css

## Task 7: Testing & Verification

- [x] 7.1 Run existing unit tests to verify no regressions
- [x] 7.2 Verify the application builds and starts successfully with `make run`
- [x] 7.3 Add Playwright visual regression tests: hero section, 4-column grid, category pills, highlight badges, breadcrumb on detail page
