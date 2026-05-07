package handler_test

import (
	"testing"

	"github.com/heidelguide/heidelguide/internal/handler"
)

func TestCategoryColorClass(t *testing.T) {
	tests := []struct {
		name  string
		color string
		want  string
	}{
		{name: "rose", color: "rose", want: "bg-rose-100 text-rose-700"},
		{name: "emerald", color: "emerald", want: "bg-emerald-100 text-emerald-700"},
		{name: "amber", color: "amber", want: "bg-amber-100 text-amber-700"},
		{name: "violet", color: "violet", want: "bg-violet-100 text-violet-700"},
		{name: "unknown color returns fallback", color: "blue", want: "bg-stone-100 text-stone-700"},
		{name: "empty string returns fallback", color: "", want: "bg-stone-100 text-stone-700"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handler.CategoryColorClass(tt.color)
			if got != tt.want {
				t.Errorf("CategoryColorClass(%q) = %q, want %q", tt.color, got, tt.want)
			}
		})
	}
}
