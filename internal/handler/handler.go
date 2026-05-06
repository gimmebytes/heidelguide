package handler

import (
	"html/template"
	"net/http"

	"github.com/heidelguide/heidelguide/internal/model"
	"github.com/heidelguide/heidelguide/internal/store"
)

// Handler holds dependencies for HTTP handlers.
type Handler struct {
	Store     *store.Store
	Templates *template.Template
	Labels    map[string]map[string]string // locale -> key -> label
}

// New creates a new Handler with the given dependencies.
func New(s *store.Store, tmpl *template.Template, labels map[string]map[string]string) *Handler {
	return &Handler{
		Store:     s,
		Templates: tmpl,
		Labels:    labels,
	}
}

// PageData is the data structure passed to templates.
type PageData struct {
	Labels    map[string]string
	Locale    string
	Landmarks []model.LandmarkWithTranslation
	Landmark  *model.LandmarkWithTranslation
}

// getLocale reads the "lang" cookie from the request, defaulting to "de".
func getLocale(r *http.Request) string {
	cookie, err := r.Cookie("lang")
	if err != nil || (cookie.Value != "de" && cookie.Value != "en") {
		return "de"
	}
	return cookie.Value
}
