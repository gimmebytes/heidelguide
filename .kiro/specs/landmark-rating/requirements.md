# Requirements Document

## Introduction

Landmark Rating System for the HeidelGuide app. Users can rate landmarks on a 1–5 star scale directly on the detail page. No login is required — device identification is handled via a UUID stored in localStorage. Each device can submit one rating per landmark and update it later, but cannot create duplicate ratings. Ratings are persisted server-side in SQLite. The average rating and total count are displayed on the detail page.

## Glossary

- **Rating**: A 1–5 star score assigned by a device to a specific landmark
- **Device_ID**: A UUID generated client-side on first visit and stored in localStorage, used to identify the rating device
- **Average_Rating**: The arithmetic mean of all ratings for a given landmark
- **Rating_Count**: The total number of ratings submitted for a given landmark
- **Detail_Page**: The page displaying full information about a single landmark

## Requirements

### Requirement 1: Device Identification

**User Story:** As a visitor, I want my device to be automatically identified without creating an account, so that my ratings are remembered across page visits.

#### Acceptance Criteria

1. WHEN a user visits the HeidelGuide app for the first time, THE app SHALL generate a UUIDv4 and store it in localStorage under the key `device_id`
2. WHEN a user visits the HeidelGuide app and a valid `device_id` already exists in localStorage, THE app SHALL reuse the existing value
3. THE Device_ID SHALL be sent to the server as a request header (`X-Device-ID`) with every rating submission
4. IF no `device_id` is present in localStorage at the time of a rating interaction, THE app SHALL generate one before submitting the request

### Requirement 2: Submit a Rating

**User Story:** As a visitor, I want to rate a landmark with 1–5 stars, so that I can express how much I liked it.

#### Acceptance Criteria

1. THE Detail_Page SHALL display a star rating widget with 5 clickable stars
2. WHEN the user clicks a star, THE app SHALL submit the rating (1–5) to the server via an HTMX POST request to `/api/landmarks/{id}/rating` including the `X-Device-ID` header
3. THE server SHALL store the rating in the database associated with the landmark ID and Device_ID
4. IF a rating from the same Device_ID for the same landmark already exists, THE server SHALL update the existing rating value instead of creating a duplicate
5. THE server SHALL validate that the rating value is an integer between 1 and 5 inclusive; invalid values SHALL return HTTP 400
6. AFTER a successful submission, THE server SHALL return an HTML fragment (HTMX partial) containing the updated average rating, count, and the user's current rating state

### Requirement 3: Display Rating Summary

**User Story:** As a visitor, I want to see the average rating and number of ratings for a landmark, so that I can gauge others' opinions.

#### Acceptance Criteria

1. THE Detail_Page SHALL display the Average_Rating as a numeric value (one decimal place) alongside a visual star representation
2. THE Detail_Page SHALL display the Rating_Count (total number of ratings) next to the average
3. WHEN the page loads, THE server SHALL compute the average and count from all ratings stored for that landmark
4. IF no ratings exist for a landmark, THE Detail_Page SHALL display an empty star row with text indicating no ratings yet

### Requirement 4: Show User's Own Rating

**User Story:** As a visitor, I want to see my previous rating highlighted, so that I know what I rated and can change it.

#### Acceptance Criteria

1. WHEN the detail page loads, THE app SHALL send the Device_ID to the server (via query parameter or header) to retrieve the user's existing rating for that landmark
2. IF the user has previously rated the landmark, THE star widget SHALL display the user's rating as filled/highlighted stars
3. IF the user has not previously rated the landmark, THE star widget SHALL display all stars as empty/unfilled
4. WHEN the user clicks a different star value, THE app SHALL update their existing rating (not create a new one)

### Requirement 5: Rating Persistence

**User Story:** As a developer, I want ratings stored reliably in SQLite, so that they survive server restarts.

#### Acceptance Criteria

1. THE database SHALL contain a `ratings` table with columns: `id`, `landmark_id`, `device_id`, `score`, `created_at`, `updated_at`
2. THE `ratings` table SHALL enforce a UNIQUE constraint on (`landmark_id`, `device_id`) to prevent duplicate ratings
3. THE server SHALL use INSERT OR REPLACE (upsert) semantics when processing a rating submission
4. THE `score` column SHALL only accept integer values between 1 and 5

### Requirement 6: Star Widget Interactivity

**User Story:** As a visitor, I want the star widget to feel responsive, so that rating feels smooth and intuitive.

#### Acceptance Criteria

1. WHEN the user hovers over a star, THE widget SHALL highlight all stars from 1 up to and including the hovered star with a visual hover state (e.g., color change)
2. WHEN the user moves the cursor away from the widget, THE widget SHALL revert to showing the user's submitted rating (or empty if none)
3. THE star widget SHALL use Alpine.js for hover state and HTMX for server submission
4. THE star widget SHALL be keyboard-accessible (operable via Tab and Enter/Space keys)
