package handler

// CategoryColorClass maps a category color slug to Tailwind CSS classes for the pill.
func CategoryColorClass(color string) string {
	colors := map[string]string{
		"rose":    "bg-rose-100 text-rose-700",
		"emerald": "bg-emerald-100 text-emerald-700",
		"amber":   "bg-amber-100 text-amber-700",
		"violet":  "bg-violet-100 text-violet-700",
	}
	if cls, ok := colors[color]; ok {
		return cls
	}
	return "bg-stone-100 text-stone-700"
}
