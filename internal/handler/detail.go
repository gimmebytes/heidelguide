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

	data := PageData{
		Labels:   h.Labels[locale],
		Locale:   locale,
		Landmark: landmark,
	}

	if h.Templates == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if err := h.Templates.ExecuteTemplate(w, "detail.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// renderNotFound renders the 404 page template, falling back to a plain text response.
func (h *Handler) renderNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	if h.Templates != nil {
		locale := "de"
		data := PageData{
			Labels: h.Labels[locale],
			Locale: locale,
		}
		if err := h.Templates.ExecuteTemplate(w, "404.html", data); err == nil {
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
