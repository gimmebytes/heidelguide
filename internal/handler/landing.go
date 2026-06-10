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
	ratings, _ := h.Store.GetAllRatingSummaries()
	data := PageData{
		Labels:    h.Labels[locale],
		Locale:    locale,
		Landmarks: landmarks,
		Ratings:   ratings,
	}
	tmpl, ok := h.Templates["landing.html"]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
