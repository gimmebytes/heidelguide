# Requirements: UX Refinement

## Requirement 1: Category Data Model

### Description
Introduce a category system for landmarks with full i18n support. Each landmark belongs to exactly one category. Categories have a machine-readable slug, a display color for UI pills, and translated names for each supported locale.

### Acceptance Criteria

- 1.1 A `categories` table exists with columns: `id` (PK), `slug` (TEXT UNIQUE NOT NULL), `color` (TEXT NOT NULL)
- 1.2 A `category_translations` table exists with columns: `id` (PK), `category_id` (FK), `locale` (TEXT NOT NULL), `name` (TEXT NOT NULL), with a UNIQUE constraint on (category_id, locale)
- 1.3 The `landmarks` table has a `category_id` column (INTEGER, FK to categories) and a `highlighted` column (INTEGER NOT NULL DEFAULT 0)
- 1.4 Four categories are seeded: architecture/rose, nature/emerald, history/amber, culture/violet
- 1.5 Each category has translations for both supported locales (de, en): ARCHITEKTUR/ARCHITECTURE, NATUR/NATURE, GESCHICHTE/HISTORY, KULTUR/CULTURE
- 1.6 Every landmark has a non-null `category_id` after seeding
- 1.7 Three landmarks are marked as highlighted: Heidelberger Schloss, Alte Brücke, Studentenkarzer
- 1.8 Category and landmark seeding is idempotent — running it multiple times does not create duplicates

---

## Requirement 2: Landing Page Visual Overhaul

### Description
Replace the current gradient hero section with a full-width scenic photograph of Heidelberg. Redesign the landmark grid to use 4 columns on desktop with playful, colorful cards featuring category pills, highlight badges, and a "Learn more" call-to-action.

### Acceptance Criteria

- 2.1 The hero section displays a full-width scenic image of Heidelberg with a gradient overlay for text readability
- 2.2 The hero section shows a localized heading ("Entdecke Heidelberg" / "Discover Heidelberg") and subtitle over the image
- 2.3 The landmark grid uses 4 columns on desktop viewports (≥1024px)
- 2.4 The landmark grid uses 2 columns on tablet viewports (768px–1023px)
- 2.5 The landmark grid uses 1 column on mobile viewports (<768px)
- 2.6 Each landmark card displays a colored category pill showing the translated category name (e.g., "ARCHITEKTUR" in German)
- 2.7 Landmark cards for highlighted landmarks display a "Highlight" badge
- 2.8 Each card shows an "HG" monogram behind the image as a loading placeholder
- 2.9 Each card displays a localized "Mehr erfahren →" / "Learn more →" link at the bottom
- 2.10 Highlighted landmarks appear before non-highlighted landmarks in the grid

---

## Requirement 3: Detail Page Overhaul

### Description
Apply the colorful design language to the detail page. Add a breadcrumb navigation bar for easy back-and-forth navigation. Display category and highlight information consistently with the landing page cards.

### Acceptance Criteria

- 3.1 The detail page shows a breadcrumb bar: "Home > Landmark Name" (localized "Startseite" in German)
- 3.2 The "Home" element in the breadcrumb is a clickable link that navigates to `/`
- 3.3 The landmark name in the breadcrumb is plain text (not a link)
- 3.4 The detail page displays the landmark's category pill with the same color coding as on the landing page
- 3.5 The detail page shows a "Highlight" badge if the landmark is highlighted
- 3.6 The top navigation bar remains visible and accessible on the detail page (no scroll-based hiding)

---

## Requirement 4: Backend & Template Infrastructure

### Description
Update the Go backend to support the new data model, extend the store queries to join category data, add template helper functions for color mapping, and add new i18n labels.

### Acceptance Criteria

- 4.1 The color palette uses warmer, more vibrant Tailwind CSS classes (rose, emerald, amber, violet) for category-specific elements
- 4.2 A template function `categoryColorClass` maps color slugs ("rose", "emerald", "amber", "violet") to corresponding Tailwind background + text classes
- 4.3 The `categoryColorClass` function returns a neutral fallback (stone) for unknown color values
- 4.4 New i18n labels are added for both locales: `learn_more` ("Mehr erfahren" / "Learn more"), `home` ("Startseite" / "Home"), `highlight` ("Highlight" / "Highlight")
- 4.5 `ListLandmarks` returns `LandmarkView` structs with `CategoryName`, `CategorySlug`, and `CategoryColor` populated from the joined category/translation data
- 4.6 `GetLandmark` returns a `LandmarkView` struct with category data populated for the detail page
