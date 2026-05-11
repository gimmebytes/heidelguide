# Tailwind CDN Dark Mode — Critical Implementation Notes

## The Problem
The locally-served Tailwind CDN script (`static/js/tailwind-cdn.js`) **overwrites `window.tailwind` with its own Proxy object** at load time. Any `window.tailwind.config` set *before* the script loads is lost.

This means `darkMode: 'class'` is NOT active by default — the CDN defaults to `darkMode: "media"`, which means dark styles respond to the OS `prefers-color-scheme` setting, NOT to a `dark` class on `<html>`. Toggling the `dark` class has no visual effect without this fix.

## The Fix
Set `window.tailwind.config` in a `<script>` tag **after** the CDN script loads:

```html
<script src="/static/js/tailwind-cdn.js"></script>
<script>
  // Must be set AFTER the CDN script — the CDN overwrites window.tailwind on load
  window.tailwind.config = { darkMode: 'class' };
</script>
```

Do NOT set it before the CDN script — it will be ignored.

## FOUC Prevention
To prevent flash of incorrect theme on page load, add an inline IIFE in `<head>` **before** any scripts. It must account for both localStorage preference AND system preference:

```html
<script>
  (function() {
    var theme = localStorage.getItem('theme');
    if (theme === 'dark' || (!theme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
      document.documentElement.classList.add('dark');
    }
  })();
</script>
```

## Alpine.js State Initialization
The Alpine `x-data` on `<body>` must also respect system preference as fallback:

```html
<body x-data="{ darkMode: localStorage.getItem('theme') === 'dark' || (!localStorage.getItem('theme') && window.matchMedia('(prefers-color-scheme: dark)').matches) }">
```

## Toggle Button
Put all toggle logic directly in the `@click` handler — do NOT use `x-effect` or `x-init`+`$watch` on `<body>`, as these have known reliability issues in Firefox/Safari:

```html
<button @click="darkMode = !darkMode; localStorage.setItem('theme', darkMode ? 'dark' : 'light'); document.documentElement.classList.toggle('dark', darkMode)">
```

## Summary of What NOT to Do
- ❌ `window.tailwind.config = ...` before the CDN script
- ❌ `x-effect` on `<body>` for theme persistence (unreliable cross-browser)
- ❌ `x-init` + `$watch` on `<body>` (unreliable in Firefox/Safari)
- ❌ Only checking `localStorage` in FOUC script (ignores system preference)
