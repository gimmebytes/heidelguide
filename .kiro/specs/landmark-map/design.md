# Design: Landmark Detail Map

## Architecture

### Approach
Minimal integration: add Leaflet library files to `static/`, include them only in the detail template, and initialize the map with coordinates already available in the template data.

No backend changes required — `Latitude` and `Longitude` are already part of `model.Landmark` and passed to the detail template via `PageData.Landmark`.

### File Changes

| File | Change |
|------|--------|
| `static/js/leaflet.js` | Add Leaflet JS library (v1.9.4, minified) |
| `static/css/leaflet.css` | Add Leaflet CSS (v1.9.4) |
| `static/img/marker-icon.png` | Leaflet default marker icon |
| `static/img/marker-icon-2x.png` | Leaflet retina marker icon |
| `static/img/marker-shadow.png` | Leaflet marker shadow |
| `templates/detail.html` | Add map section with Leaflet init |

### Template Integration

The map section is placed between the metadata pills and the description section in `detail.html`:

```html
<!-- Map -->
<section class="mb-8">
    <link rel="stylesheet" href="/static/css/leaflet.css">
    <div id="map" class="h-64 w-full rounded-xl border border-stone-200 dark:border-stone-700 z-0"></div>
    <script src="/static/js/leaflet.js"></script>
    <script>
        (function() {
            var map = L.map('map').setView([{{.Landmark.Latitude}}, {{.Landmark.Longitude}}], 16);
            L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                attribution: '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
                maxZoom: 19
            }).addTo(map);
            L.marker([{{.Landmark.Latitude}}, {{.Landmark.Longitude}}]).addTo(map);
        })();
    </script>
</section>
```

### Asset Loading Strategy

Leaflet CSS and JS are loaded inline in the detail template only (not in `base.html`). This ensures:
- Landing page is not affected (REQ-8)
- No unnecessary asset loading on other pages
- Simple, self-contained approach

### Marker Icons

Leaflet expects marker icons at a specific path. Configure `L.Icon.Default` to point to `/static/img/`:

```javascript
L.Icon.Default.imagePath = '/static/img/';
```

### Dark Mode

The map tiles themselves cannot be dark-themed (they come from OSM). The container uses `dark:border-stone-700` for border consistency. This is acceptable — map tiles in light mode within a dark UI is a common pattern (Google Maps, Apple Maps do the same).

### Responsive Behavior

The map container uses `w-full h-64` — full width of the content column, fixed 256px height. This works well on all screen sizes since the content column is already max-width constrained (`max-w-4xl`).
