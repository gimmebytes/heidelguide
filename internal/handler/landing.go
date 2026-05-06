package handler

import "net/http"

// Landing handles GET / — renders the landing page with all landmarks.
func (h *Handler) Landing(w http.ResponseWriter, r *http.Request) {
	locale := getLocale(r)
	landmarks, err := h.Store.ListLandmarks(locale)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := PageData{
		Labels:    h.Labels[locale],
		Locale:    locale,
		Landmarks: landmarks,
	}
	if h.Templates == nil {
		http.Error(w, "templates not loaded", http.StatusInternalServerError)
		return
	}
	if err := h.Templates.ExecuteTemplate(w, "landing.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
