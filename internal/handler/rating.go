package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/heidelguide/heidelguide/internal/model"
)

// SubmitRating handles POST /api/landmarks/{id}/rating.
func (h *Handler) SubmitRating(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid landmark id", http.StatusBadRequest)
		return
	}

	deviceID := r.Header.Get("X-Device-ID")
	if deviceID == "" {
		http.Error(w, "missing X-Device-ID header", http.StatusBadRequest)
		return
	}

	scoreStr := r.FormValue("score")
	score, err := strconv.Atoi(scoreStr)
	if err != nil || score < 1 || score > 5 {
		http.Error(w, "score must be 1-5", http.StatusBadRequest)
		return
	}

	if err := h.Store.UpsertRating(id, deviceID, score); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	summary, err := h.Store.GetRatingSummary(id, deviceID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get locale for labels
	locale := getLocale(r)

	data := PageData{
		Labels: h.Labels[locale],
		Locale: locale,
		Landmark: &model.LandmarkView{
			Landmark: model.Landmark{ID: id},
		},
		Rating: summary,
	}

	tmpl, ok := h.Templates["rating_widget.html"]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "rating_widget", data)
}
