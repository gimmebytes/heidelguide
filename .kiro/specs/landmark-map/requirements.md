# Requirements: Landmark Detail Map

## Overview
Display an interactive OpenStreetMap map on the landmark detail page showing the landmark's location. Uses Leaflet.js served locally (no CDN).

## Requirements

### REQ-1: Map Display
- **Given** a user navigates to a landmark detail page
- **When** the page loads
- **Then** an interactive map is displayed showing the landmark's geographic location with a marker

### REQ-2: Map Configuration
- **Given** the map is displayed
- **When** it initializes
- **Then** it is centered on the landmark's coordinates at a zoom level appropriate for a city-level POI (zoom ~16)

### REQ-3: Marker
- **Given** the map is displayed
- **When** it initializes
- **Then** a single marker is placed at the landmark's exact latitude/longitude

### REQ-4: Self-Contained Assets
- **Given** the app serves all assets locally
- **When** the map loads
- **Then** Leaflet JS and CSS are served from `/static/js/` and `/static/css/` (no external CDN for library code)

### REQ-5: Map Tiles
- **Given** the map needs tile imagery
- **When** rendering
- **Then** it uses OpenStreetMap tile servers (`tile.openstreetmap.org`) with proper attribution

### REQ-6: Responsive Layout
- **Given** the map is displayed on various screen sizes
- **When** the viewport changes
- **Then** the map container adapts responsively (full width, fixed height)

### REQ-7: Dark Mode Compatibility
- **Given** the user has dark mode enabled
- **When** the map is displayed
- **Then** the map container styling (border, rounded corners) matches the dark mode theme

### REQ-8: No Impact on Landing Page
- **Given** Leaflet assets are added to the project
- **When** the landing page loads
- **Then** Leaflet JS/CSS are NOT loaded (only on detail pages)
