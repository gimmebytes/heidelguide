package handler

import (
	"fmt"

	"github.com/heidelguide/heidelguide/internal/store"
)

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

// Seq returns a slice of integers from start to end (inclusive).
func Seq(start, end int) []int {
	s := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		s = append(s, i)
	}
	return s
}

// RatingFor returns the formatted average rating for a landmark ID, or empty string if none.
func RatingFor(ratings map[int64]*store.RatingSummary, id int64) string {
	if ratings == nil {
		return ""
	}
	r, ok := ratings[id]
	if !ok || r.Count == 0 {
		return ""
	}
	return fmt.Sprintf("%.1f", r.Average)
}
