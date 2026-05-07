package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/heidelguide/heidelguide/internal/handler"
	"github.com/heidelguide/heidelguide/internal/i18n"
	"github.com/heidelguide/heidelguide/internal/middleware"
	"github.com/heidelguide/heidelguide/internal/store"
)

func main() {
	// Set up structured logger.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Open SQLite store.
	s, err := store.Open("heidelberg.db")
	if err != nil {
		slog.Error("failed to open database", "error", err)
		os.Exit(1)
	}
	defer s.Close()
	slog.Info("database opened", "path", "heidelberg.db")

	// Run migrations.
	if err := s.Migrate(); err != nil {
		slog.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}
	slog.Info("migrations complete")

	// Seed data.
	if err := s.Seed(); err != nil {
		slog.Error("failed to seed data", "error", err)
		os.Exit(1)
	}
	slog.Info("seed data ready")

	// Seed categories.
	if err := s.SeedCategories(); err != nil {
		slog.Error("failed to seed categories", "error", err)
		os.Exit(1)
	}
	slog.Info("category seed data ready")

	// Assign default categories to landmarks.
	if err := s.AssignDefaultCategories(); err != nil {
		slog.Error("failed to assign default categories", "error", err)
		os.Exit(1)
	}
	slog.Info("default categories assigned")

	// Define template functions.
	funcMap := template.FuncMap{
		"categoryColorClass": handler.CategoryColorClass,
	}

	// Parse templates - each page template is combined with base.html.
	templates := make(map[string]*template.Template)
	baseFile := filepath.Join("templates", "base.html")

	pages := []string{"landing.html", "detail.html", "404.html"}
	for _, page := range pages {
		pageFile := filepath.Join("templates", page)
		t, err := template.New(page).Funcs(funcMap).ParseFiles(baseFile, pageFile)
		if err != nil {
			slog.Error("failed to parse template", "file", page, "error", err)
			os.Exit(1)
		}
		templates[page] = t
	}

	// Create handler with dependencies.
	h := handler.New(s, templates, i18n.Labels())

	// Create chi router and register routes.
	r := chi.NewRouter()

	// Wire request logging middleware.
	r.Use(middleware.RequestLogger)

	r.Get("/", h.Landing)
	r.Get("/landmarks/{id}", h.Detail)
	r.Post("/language", h.SwitchLanguage)

	// Serve static files.
	fileServer := http.FileServer(http.Dir("static"))
	r.Get("/static/*", http.StripPrefix("/static", fileServer).ServeHTTP)

	// Determine port.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	slog.Info("server listening", "addr", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
