# Implementation Plan: Dark Mode

## Overview

Add a toggleable dark color scheme to the HeidelGuide app using Tailwind CSS class-based dark mode, Alpine.js for state management, and localStorage for persistence. Tasks are ordered for maximum visual impact during a live demo with hot-reload (Air) — each task produces an immediately visible change when the page refreshes.

## Tasks

- [x] 1. Foundation: Tailwind dark mode config, FOUC prevention, and Alpine.js body setup
  - [x] 1.1 Add Tailwind darkMode config, FOUC prevention script, and Alpine.js x-data/x-init on body in `templates/base.html`
    - Add `tailwind.config` script block **after** the Tailwind CDN script to set `darkMode: 'class'` — IMPORTANT: must be placed after the CDN script, not before, because the CDN overwrites `window.tailwind` on load
    - Add inline FOUC prevention IIFE after meta tags that reads localStorage and adds `dark` class to `<html>` — must also check `window.matchMedia('(prefers-color-scheme: dark)')` as fallback when no localStorage value exists
    - Add `x-data="{ darkMode: localStorage.getItem('theme') === 'dark' || (!localStorage.getItem('theme') && window.matchMedia('(prefers-color-scheme: dark)').matches) }"` to the `<body>` element — put all toggle logic in the `@click` handler (see Task 2.1), do NOT use `x-effect` or `x-init`+`$watch` (unreliable cross-browser)
    - Add `transition-colors duration-200 ease-in-out` to the body class list for smooth theme transitions
    - _Requirements: 1.1, 1.2, 1.3, 2.3, 2.4, 2.8, 3.1, 3.2, 3.3, 3.4, 3.5, 8.1, 8.2, 8.3_

- [x] 2. Toggle button in navbar
  - [x] 2.1 Add the dark mode toggle button with sun/moon SVG icons to the navigation bar in `templates/base.html`
    - Insert a `<button>` with `@click="darkMode = !darkMode"` in the nav's flex container, before the language switcher
    - Add moon SVG icon with `x-show="!darkMode"` (shown in light mode, indicates "switch to dark")
    - Add sun SVG icon with `x-show="darkMode"` (shown in dark mode, indicates "switch to light")
    - Add dynamic `aria-label` via `:aria-label="darkMode ? 'Switch to light mode' : 'Switch to dark mode'"`
    - Style with `px-2 py-1 rounded bg-emerald-800 hover:bg-emerald-700 transition-colors` (matching language switcher button sizing per navbar-buttons steering guideline)
    - _Requirements: 2.1, 2.2, 2.5, 2.7_

- [x] 3. Dark styling for base layout (body, nav, footer)
  - [x] 3.1 Add dark variant classes to body, navigation, and footer elements in `templates/base.html`
    - Add `dark:bg-stone-900 dark:text-stone-100` to the `<body>` element
    - Add `dark:bg-stone-800` to the `<nav>` element (neutral dark background for better contrast with buttons)
    - Add dark variants to language-switcher buttons: active state `dark:bg-amber-500 dark:text-white` and inactive state `dark:bg-stone-700 dark:text-stone-300 dark:hover:bg-stone-600`
    - Add `dark:bg-stone-800 dark:text-stone-400` to the `<footer>` element
    - Ensure footer links remain distinguishable with `dark:text-amber-400 dark:hover:text-amber-300`
    - _Requirements: 4.1, 4.2, 4.3, 4.4_

- [x] 4. Dark styling for landing page
  - [x] 4.1 Add dark variant classes to landmark cards and hero section in `templates/landing.html`
    - Add `dark:bg-stone-800 dark:border-stone-700` to each card link element
    - Add `dark:bg-stone-700` to the image placeholder area
    - Add `dark:text-stone-100 dark:group-hover:text-amber-400` to card title `<h2>`
    - Add `dark:text-stone-300` to card description `<p>`
    - Add `dark:text-amber-400 dark:group-hover:text-amber-300` to the "learn more" span
    - Add `dark:shadow-lg dark:shadow-stone-900/50` to card hover state
    - Ensure hero gradient overlay remains legible in dark mode (existing `from-stone-900/70` already works)
    - _Requirements: 5.1, 5.2, 5.3, 5.4, 5.5_

- [x] 5. Dark styling for detail page
  - [x] 5.1 Add dark variant classes to headings, text, badges, breadcrumbs, and borders in `templates/detail.html`
    - Add `dark:text-stone-100` to `<h1>` heading
    - Add `dark:text-stone-100` to `<h2>` section headings (Description, History)
    - Add `dark:text-stone-300` to body text paragraphs (description, history)
    - Add `dark:bg-stone-700 dark:text-stone-200` to metadata badge spans
    - Add `dark:text-amber-400 dark:hover:text-amber-300` to breadcrumb link
    - Add `dark:text-stone-500` to breadcrumb separator SVG
    - Add `dark:text-stone-300` to current-page breadcrumb span
    - Add `dark:border-stone-700` to the bottom border separator
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5, 6.6_

- [x] 6. Dark styling for error page and final polish
  - [x] 6.1 Add dark variant classes to the 404 error page in `templates/404.html`
    - Add `dark:text-stone-100` to the heading `<h1>`
    - Add `dark:text-stone-300` to the description `<p>`
    - The "404" number already uses `text-amber-600/30` which works in both modes
    - The "back to home" button already uses `bg-amber-600` which works in both modes
    - Verify contrast ratios meet 4.5:1 minimum in dark mode
    - _Requirements: 7.1, 7.2, 7.3, 7.4_

- [x] 7. (Optional) Playwright visual regression tests for dark mode
  - [x] 7.1 Extend `tests/visual.spec.ts` with dark mode screenshots for landing, detail, and 404 pages
    - Add test cases that set `localStorage.setItem('theme', 'dark')` before navigating to each page
    - Capture dark mode screenshots as new baselines
    - Add a test for toggle interaction (click toggle, verify `dark` class on `<html>`)
    - Run via `make test-visual` — update snapshots with `--update-snapshots` flag on first run
    - _Note: Skip this task during the live demo. Run it afterwards to establish dark mode baselines._

## Notes

- No Go backend changes are needed — this feature is entirely client-side (templates + Tailwind + Alpine.js)
- The app runs with Air hot-reload, so each saved template change is immediately visible in the browser
- No property-based tests apply (the feature is UI rendering with a 2-value boolean state)
- Visual regression tests can be extended in `tests/visual.spec.ts` to capture dark mode screenshots
- Tasks are ordered for maximum demo impact: foundation → interactive toggle → pages going dark one by one

## Task Dependency Graph

```json
{
  "waves": [
    { "id": 0, "tasks": ["1.1"] },
    { "id": 1, "tasks": ["2.1"] },
    { "id": 2, "tasks": ["3.1"] },
    { "id": 3, "tasks": ["4.1", "5.1", "6.1"] }
  ]
}
```
