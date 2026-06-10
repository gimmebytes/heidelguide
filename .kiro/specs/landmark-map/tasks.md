# Tasks: Landmark Detail Map

## Task 1: Download Leaflet assets
- [ ] Download Leaflet v1.9.4 minified JS → `static/js/leaflet.js`
- [ ] Download Leaflet v1.9.4 CSS → `static/css/leaflet.css`
- [ ] Download marker icons (marker-icon.png, marker-icon-2x.png, marker-shadow.png) → `static/img/`

**Requirements:** REQ-4

## Task 2: Add map section to detail template
- [ ] Add Leaflet CSS link in the map section of `templates/detail.html`
- [ ] Add map container `<div>` with responsive styling and dark mode border
- [ ] Add Leaflet JS script tag
- [ ] Add initialization script: set view, add tile layer with OSM attribution, add marker
- [ ] Configure `L.Icon.Default.imagePath` to `/static/img/`
- [ ] Place the map section between metadata pills and description

**Requirements:** REQ-1, REQ-2, REQ-3, REQ-5, REQ-6, REQ-7

## Task 3: Verify no landing page impact
- [ ] Confirm Leaflet CSS/JS are NOT loaded on the landing page
- [ ] Run existing tests (`make test`)

**Requirements:** REQ-8

## Task 4: Manual smoke test
- [ ] Start the app (`make run`)
- [ ] Navigate to a landmark detail page
- [ ] Verify map displays with correct marker position
- [ ] Verify map works in dark mode
- [ ] Verify landing page loads without Leaflet assets
