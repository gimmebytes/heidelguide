package store

import "fmt"

// SeedCategories inserts the 4 landmark categories with translations for de and en.
// It is idempotent: if categories already exist, it returns nil without changes.
func (s *Store) SeedCategories() error {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check categories count: %w", err)
	}
	if count > 0 {
		return nil
	}

	type seedCategory struct {
		Slug  string
		Color string
		Names map[string]string // locale -> name
	}

	categories := []seedCategory{
		{"architecture", "rose", map[string]string{"de": "ARCHITEKTUR", "en": "ARCHITECTURE"}},
		{"nature", "emerald", map[string]string{"de": "NATUR", "en": "NATURE"}},
		{"history", "amber", map[string]string{"de": "GESCHICHTE", "en": "HISTORY"}},
		{"culture", "violet", map[string]string{"de": "KULTUR", "en": "CULTURE"}},
	}

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin category seed transaction: %w", err)
	}
	defer tx.Rollback()

	for _, cat := range categories {
		result, err := tx.Exec(
			"INSERT INTO categories (slug, color) VALUES (?, ?)",
			cat.Slug, cat.Color)
		if err != nil {
			return fmt.Errorf("failed to insert category %q: %w", cat.Slug, err)
		}
		catID, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get category id for %q: %w", cat.Slug, err)
		}
		for locale, name := range cat.Names {
			_, err = tx.Exec(
				"INSERT INTO category_translations (category_id, locale, name) VALUES (?, ?, ?)",
				catID, locale, name)
			if err != nil {
				return fmt.Errorf("failed to insert category translation (%s/%s): %w", cat.Slug, locale, err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit category seed transaction: %w", err)
	}
	return nil
}
