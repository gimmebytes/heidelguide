package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Detail handles GET /landmarks/{id} — renders the detail page for a single landmark.
func (h *Handler) Detail(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.renderNotFound(w)
		return
	}

	locale := getLocale(r)
	landmark, err := h.Store.GetLandmark(id, locale)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if landmark == nil {
		h.renderNotFound(w)
		return
	}

	deviceID := r.Header.Get("X-Device-ID")
	rating, _ := h.Store.GetRatingSummary(id, deviceID)

	data := PageData{
		Labels:   h.Labels[locale],
		Locale:   locale,
		Landmark: landmark,
		Rating:   rating,
	}

	tmpl, ok := h.Templates["detail.html"]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// renderNotFound renders the 404 page template, falling back to a plain text response.
func (h *Handler) renderNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	tmpl, ok := h.Templates["404.html"]
	if ok {
		locale := "de"
		data := PageData{
			Labels: h.Labels[locale],
			Locale: locale,
		}
		if err := tmpl.ExecuteTemplate(w, "base.html", data); err == nil {
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
