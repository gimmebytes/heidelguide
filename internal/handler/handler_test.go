package handler_test

import (
	"context"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/heidelguide/heidelguide/internal/handler"
	"github.com/heidelguide/heidelguide/internal/i18n"
	"github.com/heidelguide/heidelguide/internal/store"
)

// setupTestHandler creates a Handler backed by an in-memory SQLite database
// with migrations and seed data applied, and templates parsed from disk.
func setupTestHandler(t *testing.T) *handler.Handler {
	t.Helper()

	s, err := store.Open(":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}
	t.Cleanup(func() { s.Close() })

	if err := s.Migrate(); err != nil {
		t.Fatalf("failed to run migrations: %v", err)
	}
	if err := s.Seed(); err != nil {
		t.Fatalf("failed to seed data: %v", err)
	}

	if err := s.SeedCategories(); err != nil {
		t.Fatalf("failed to seed categories: %v", err)
	}
	if err := s.AssignDefaultCategories(); err != nil {
		t.Fatalf("failed to assign default categories: %v", err)
	}

	// Define template functions (must match main.go).
	funcMap := template.FuncMap{
		"categoryColorClass": handler.CategoryColorClass,
		"seq":                handler.Seq,
		"ratingFor":          handler.RatingFor,
	}

	// Parse templates — paths relative to project root (tests run from package dir).
	templatesDir := filepath.Join("..", "..", "templates")
	baseFile := filepath.Join(templatesDir, "base.html")
	ratingWidget := filepath.Join(templatesDir, "partials", "rating_widget.html")

	templates := make(map[string]*template.Template)
	pages := []string{"landing.html", "detail.html", "404.html"}
	for _, page := range pages {
		pageFile := filepath.Join(templatesDir, page)
		parseFiles := []string{baseFile, pageFile}
		if page == "detail.html" {
			parseFiles = append(parseFiles, ratingWidget)
		}
		tmpl, err := template.New(page).Funcs(funcMap).ParseFiles(parseFiles...)
		if err != nil {
			t.Fatalf("failed to parse template %s: %v", page, err)
		}
		templates[page] = tmpl
	}

	return handler.New(s, templates, i18n.Labels())
}

// withChiURLParam adds a chi URL parameter to the request context.
func withChiURLParam(r *http.Request, key, value string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, value)
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestLanding_Returns200WithLandmarkContent(t *testing.T) {
	h := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	h.Landing(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// Verify landmark content is present (German by default).
	if !strings.Contains(body, "Heidelberger Schloss") {
		t.Error("expected landing page to contain 'Heidelberger Schloss'")
	}
	if !strings.Contains(body, "Alte Brücke") {
		t.Error("expected landing page to contain 'Alte Brücke'")
	}
	// Verify the page title is rendered.
	if !strings.Contains(body, "Heidelguide") {
		t.Error("expected landing page to contain 'Heidelguide'")
	}
}

func TestLanding_RespectsLocaleCookie(t *testing.T) {
	h := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "lang", Value: "en"})
	rec := httptest.NewRecorder()

	h.Landing(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// English content should be rendered.
	if !strings.Contains(body, "Heidelberg Castle") {
		t.Error("expected landing page to contain 'Heidelberg Castle' in English locale")
	}
}

func TestDetail_Returns200ForValidID(t *testing.T) {
	h := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/landmarks/1", nil)
	req = withChiURLParam(req, "id", "1")
	rec := httptest.NewRecorder()

	h.Detail(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	// Landmark 1 is Heidelberg Castle.
	if !strings.Contains(body, "Heidelberger Schloss") {
		t.Error("expected detail page to contain 'Heidelberger Schloss'")
	}
}

func TestDetail_Returns404ForInvalidID(t *testing.T) {
	h := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/landmarks/9999", nil)
	req = withChiURLParam(req, "id", "9999")
	rec := httptest.NewRecorder()

	h.Detail(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}
}

func TestDetail_Returns404ForNonNumericID(t *testing.T) {
	h := setupTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/landmarks/abc", nil)
	req = withChiURLParam(req, "id", "abc")
	rec := httptest.NewRecorder()

	h.Detail(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}
}

func TestSwitchLanguage_SetsCookieAndRedirects(t *testing.T) {
	h := setupTestHandler(t)

	form := url.Values{}
	form.Set("locale", "en")
	req := httptest.NewRequest(http.MethodPost, "/language", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "/landmarks/2")
	rec := httptest.NewRecorder()

	h.SwitchLanguage(rec, req)

	// Should redirect (303 See Other).
	if rec.Code != http.StatusSeeOther {
		t.Fatalf("expected status 303, got %d", rec.Code)
	}

	// Should redirect back to the referer.
	location := rec.Header().Get("Location")
	if location != "/landmarks/2" {
		t.Errorf("expected redirect to /landmarks/2, got %q", location)
	}

	// Should set the lang cookie.
	cookies := rec.Result().Cookies()
	var langCookie *http.Cookie
	for _, c := range cookies {
		if c.Name == "lang" {
			langCookie = c
			break
		}
	}
	if langCookie == nil {
		t.Fatal("expected lang cookie to be set")
	}
	if langCookie.Value != "en" {
		t.Errorf("expected lang cookie value 'en', got %q", langCookie.Value)
	}
	if !langCookie.HttpOnly {
		t.Error("expected lang cookie to be HttpOnly")
	}
	if langCookie.SameSite != http.SameSiteLaxMode {
		t.Error("expected lang cookie SameSite=Lax")
	}
}

func TestSwitchLanguage_DefaultsToDeForInvalidLocale(t *testing.T) {
	h := setupTestHandler(t)

	form := url.Values{}
	form.Set("locale", "fr")
	req := httptest.NewRequest(http.MethodPost, "/language", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	h.SwitchLanguage(rec, req)

	cookies := rec.Result().Cookies()
	var langCookie *http.Cookie
	for _, c := range cookies {
		if c.Name == "lang" {
			langCookie = c
			break
		}
	}
	if langCookie == nil {
		t.Fatal("expected lang cookie to be set")
	}
	if langCookie.Value != "de" {
		t.Errorf("expected lang cookie value 'de' for invalid locale, got %q", langCookie.Value)
	}
}

func TestSwitchLanguage_RedirectsToRootWhenNoReferer(t *testing.T) {
	h := setupTestHandler(t)

	form := url.Values{}
	form.Set("locale", "de")
	req := httptest.NewRequest(http.MethodPost, "/language", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	h.SwitchLanguage(rec, req)

	location := rec.Header().Get("Location")
	if location != "/" {
		t.Errorf("expected redirect to /, got %q", location)
	}
}
