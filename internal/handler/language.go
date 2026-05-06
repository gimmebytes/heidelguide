package handler

import (
	"net/http"
)

// SwitchLanguage handles POST /language — sets the language cookie and redirects back.
func (h *Handler) SwitchLanguage(w http.ResponseWriter, r *http.Request) {
	locale := r.FormValue("locale")
	if locale != "de" && locale != "en" {
		locale = "de"
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "lang",
		Value:    locale,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}

	http.Redirect(w, r, referer, http.StatusSeeOther)
}
