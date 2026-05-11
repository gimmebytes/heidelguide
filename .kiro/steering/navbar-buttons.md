# Navbar Button Styling

## Consistency Rule
All action buttons in the navigation bar (dark mode toggle, language switcher, any future buttons) must be visually consistent in size and shape.

## Sizing
- Use `px-2 py-1` padding for all navbar buttons (both icon-only and text buttons)
- Icon-only buttons: use `h-4 w-4` for SVG icons inside navbar buttons
- Text buttons: use `text-sm` font size
- All buttons use `rounded` border radius (not `rounded-lg`)

## Spacing
- Group related buttons with `gap-1` (e.g., language switcher DE/EN)
- Separate button groups with `gap-3` in the parent flex container

## Dark Mode Toggle Specifics
- The toggle button should match the visual weight of the language switcher buttons
- Use the same background/hover colors as inactive language buttons: `bg-emerald-800 hover:bg-emerald-700`
