# Requirements Document

## Introduction

Dark Mode for the HeidelGuide app. This feature adds a toggleable dark color scheme to all pages, using Tailwind CSS class-based dark mode and Alpine.js for client-side state management. The user preference is persisted in localStorage. This feature is designed as a live-coding demo for the trade/off ai summit in Heidelberg, with tasks structured for incremental, visually impressive changes during hot-reload development.

## Glossary

- **Dark_Mode_Toggle**: A UI button in the navigation bar that switches between light and dark color schemes
- **Theme_State**: The current color scheme preference (light or dark), managed client-side via Alpine.js
- **HeidelGuide_App**: The Go-based web application serving Heidelberg landmark information
- **Navigation_Bar**: The top-level navigation component rendered on every page
- **Landing_Page**: The main page displaying a hero image and landmark card grid
- **Detail_Page**: The page displaying full information about a single landmark
- **Error_Page**: The 404 page displayed when a route is not found
- **LocalStorage**: The browser's Web Storage API used to persist the theme preference across sessions

## Requirements

### Requirement 1: Tailwind Dark Mode Configuration

**User Story:** As a developer, I want Tailwind CSS configured for class-based dark mode, so that dark variants of utility classes are activated by a `dark` class on the HTML element.

#### Acceptance Criteria

1. THE HeidelGuide_App SHALL configure Tailwind CSS to use class-based dark mode strategy (darkMode: 'class') via the Tailwind CDN script configuration
2. WHEN the `dark` class is present on the `<html>` element, THE HeidelGuide_App SHALL apply styles defined by Tailwind `dark:` variant utility classes to their respective elements
3. IF the `dark` class is not present on the `<html>` element, THEN THE HeidelGuide_App SHALL not apply any styles defined by Tailwind `dark:` variant utility classes

### Requirement 2: Dark Mode Toggle Component

**User Story:** As a user, I want a visible toggle button in the navigation bar, so that I can switch between light and dark mode.

#### Acceptance Criteria

1. THE Navigation_Bar SHALL display a Dark_Mode_Toggle button with an inline SVG sun icon (for switching to light) and an inline SVG moon icon (for switching to dark)
2. WHEN the user clicks the Dark_Mode_Toggle, THE HeidelGuide_App SHALL switch the Theme_State from light to dark or from dark to light
3. WHEN the Theme_State changes to dark, THE HeidelGuide_App SHALL add the `dark` class to the `<html>` element
4. WHEN the Theme_State changes to light, THE HeidelGuide_App SHALL remove the `dark` class from the `<html>` element
5. THE Dark_Mode_Toggle SHALL display the moon icon when Theme_State is light and the sun icon when Theme_State is dark
6. WHEN the Theme_State changes, THE HeidelGuide_App SHALL persist the selected theme in localStorage so that the preference is retained across page navigations and browser sessions, defaulting to light mode when no stored preference exists
7. THE Dark_Mode_Toggle SHALL include an `aria-label` attribute that reflects the current action (e.g., indicating "switch to dark mode" or "switch to light mode"), a `role` of button, and shall be operable via keyboard (Enter and Space keys)
8. WHEN the page loads and a dark theme preference exists in localStorage, THE HeidelGuide_App SHALL apply the `dark` class to the `<html>` element before the page content renders to prevent a flash of the incorrect theme

### Requirement 3: Theme Preference Persistence

**User Story:** As a user, I want my dark mode preference remembered, so that the app respects my choice when I return.

#### Acceptance Criteria

1. WHEN the user changes the Theme_State, THE HeidelGuide_App SHALL store the selected value in LocalStorage under the key `theme` with the value `dark` or `light` corresponding to the chosen state
2. WHEN the page loads and a theme preference with a valid value (`dark` or `light`) exists in LocalStorage, THE HeidelGuide_App SHALL apply the stored preference before the page content becomes visible to the user
3. WHEN the page loads and no theme preference exists in LocalStorage, THE HeidelGuide_App SHALL default to light mode
4. IF the page loads and the LocalStorage key `theme` contains a value other than `dark` or `light`, THEN THE HeidelGuide_App SHALL discard the invalid value and default to light mode
5. THE HeidelGuide_App SHALL apply the theme preference via an inline script in the document head so that no flash of incorrect theme is visible during page load

### Requirement 4: Dark Mode Styling for Base Layout

**User Story:** As a user, I want the overall page background, text colors, navigation bar, and footer to adapt to dark mode, so that the entire app feels cohesive in both themes.

#### Acceptance Criteria

1. WHILE Theme_State is dark, THE HeidelGuide_App SHALL render the page body with background color stone-900 and text color stone-100
2. WHILE Theme_State is dark, THE Navigation_Bar SHALL render with background color emerald-950 and all text elements (logo, links, language buttons) at a minimum contrast ratio of 4.5:1 against the background
3. WHILE Theme_State is dark, THE HeidelGuide_App SHALL render the footer with background color emerald-950 and text color stone-400, and footer links SHALL remain visually distinguishable from surrounding text
4. WHILE Theme_State is dark, THE Navigation_Bar SHALL render language-switcher buttons with background and text colors that remain distinguishable between the active-locale state and the inactive-locale state

### Requirement 5: Dark Mode Styling for Landing Page

**User Story:** As a user, I want the landing page cards and hero section to look good in dark mode, so that browsing landmarks is comfortable at night.

#### Acceptance Criteria

1. WHILE Theme_State is dark, THE Landing_Page SHALL render landmark cards with a dark background (stone-800), light text (stone-100), a dark border (stone-700), and a dark image placeholder area (stone-700)
2. WHILE Theme_State is dark, THE Landing_Page SHALL render card description text in stone-300 and the "learn more" link in amber-400
3. WHILE Theme_State is dark, THE Landing_Page SHALL render the hero section gradient overlay with a minimum opacity of 70% on the dark end (stone-900/70) so that all overlaid text remains legible against the background image
4. WHEN Theme_State changes to dark, THE Landing_Page SHALL apply dark styling via Tailwind dark-variant classes that activate when the `dark` class is present on the root `<html>` element
5. WHILE Theme_State is dark, THE Landing_Page SHALL render card hover states with an elevated shadow (shadow-lg using a dark-appropriate shadow color) and the title text transitioning to amber-400

### Requirement 6: Dark Mode Styling for Detail Page

**User Story:** As a user, I want the landmark detail page to be readable in dark mode, so that I can comfortably read descriptions and history.

#### Acceptance Criteria

1. WHILE Theme_State is dark, THE Detail_Page SHALL render heading text (h1, h2) in stone-100
2. WHILE Theme_State is dark, THE Detail_Page SHALL render body text (description paragraph, history paragraph) in stone-300
3. WHILE Theme_State is dark, THE Detail_Page SHALL render metadata badges with a stone-700 background and stone-200 text color
4. WHILE Theme_State is dark, THE Detail_Page SHALL render the breadcrumb link in amber-400, the breadcrumb separator in stone-500, and the current-page breadcrumb text in stone-300
5. WHILE Theme_State is dark, THE Detail_Page SHALL render the bottom border separator in stone-700
6. WHILE Theme_State is dark, THE Detail_Page SHALL render the page background in stone-900

### Requirement 7: Dark Mode Styling for Error Page

**User Story:** As a user, I want the 404 page to match the dark theme, so that the experience is consistent even on error pages.

#### Acceptance Criteria

1. WHILE the `dark` class is present on the HTML root element, THE Error_Page SHALL render the heading text in stone-100 and the description text in stone-300
2. WHILE the `dark` class is present on the HTML root element, THE Error_Page SHALL render the page background in a dark color consistent with the site-wide dark theme background
3. WHILE the `dark` class is present on the HTML root element, THE Error_Page SHALL render the "404" number in amber-600 with 30% opacity and the "back to home" button with an amber-600 background
4. WHILE the `dark` class is present on the HTML root element, THE Error_Page SHALL maintain a minimum contrast ratio of 4.5:1 between all text elements and their background

### Requirement 8: Transition Animation

**User Story:** As a user, I want a smooth visual transition when switching themes, so that the change feels polished rather than jarring.

#### Acceptance Criteria

1. WHEN the Theme_State changes, THE HeidelGuide_App SHALL apply a CSS transition to background-color and color properties with a duration of 200ms and ease-in-out timing
2. THE transition SHALL be applied globally via a Tailwind utility class on the body element
3. THE transition SHALL only animate color-related properties (background-color, color, border-color) to avoid layout performance issues
