package store

import (
	"database/sql"

	"github.com/heidelguide/heidelguide/internal/model"
)

// ListLandmarks returns all landmarks joined with their translations and category info for the given locale.
func (s *Store) ListLandmarks(locale string) ([]model.LandmarkView, error) {
	rows, err := s.db.Query(`
		SELECT l.id, l.latitude, l.longitude, l.image_filename, l.year_built, l.year_destroyed,
		       l.category_id, l.highlighted,
		       t.id, t.landmark_id, t.locale, t.name, t.description, t.history,
		       COALESCE(c.slug, '') as category_slug,
		       COALESCE(c.color, '') as category_color,
		       COALESCE(ct.name, c.slug, '') as category_name
		FROM landmarks l
		JOIN landmark_translations t ON l.id = t.landmark_id AND t.locale = ?
		LEFT JOIN categories c ON l.category_id = c.id
		LEFT JOIN category_translations ct ON c.id = ct.category_id AND ct.locale = ?
		ORDER BY l.highlighted DESC, l.id ASC`,
		locale, locale)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var landmarks []model.LandmarkView
	for rows.Next() {
		var lv model.LandmarkView
		var highlighted int
		if err := rows.Scan(
			&lv.Landmark.ID,
			&lv.Landmark.Latitude,
			&lv.Landmark.Longitude,
			&lv.Landmark.ImageFilename,
			&lv.Landmark.YearBuilt,
			&lv.Landmark.YearDestroyed,
			&lv.Landmark.CategoryID,
			&highlighted,
			&lv.LandmarkTranslation.ID,
			&lv.LandmarkTranslation.LandmarkID,
			&lv.LandmarkTranslation.Locale,
			&lv.LandmarkTranslation.Name,
			&lv.LandmarkTranslation.Description,
			&lv.LandmarkTranslation.History,
			&lv.CategorySlug,
			&lv.CategoryColor,
			&lv.CategoryName,
		); err != nil {
			return nil, err
		}
		lv.Landmark.Highlighted = highlighted != 0
		landmarks = append(landmarks, lv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return landmarks, nil
}

// GetLandmark returns a single landmark joined with its translation for the given locale.
// Returns nil (not an error) if the landmark is not found.
func (s *Store) GetLandmark(id int64, locale string) (*model.LandmarkWithTranslation, error) {
	row := s.db.QueryRow(`
		SELECT l.id, l.latitude, l.longitude, l.image_filename, l.year_built, l.year_destroyed,
		       t.id, t.landmark_id, t.locale, t.name, t.description, t.history
		FROM landmarks l
		JOIN landmark_translations t ON l.id = t.landmark_id
		WHERE t.locale = ? AND l.id = ?`, locale, id)

	var lw model.LandmarkWithTranslation
	err := row.Scan(
		&lw.Landmark.ID,
		&lw.Landmark.Latitude,
		&lw.Landmark.Longitude,
		&lw.Landmark.ImageFilename,
		&lw.Landmark.YearBuilt,
		&lw.Landmark.YearDestroyed,
		&lw.LandmarkTranslation.ID,
		&lw.LandmarkTranslation.LandmarkID,
		&lw.LandmarkTranslation.Locale,
		&lw.LandmarkTranslation.Name,
		&lw.LandmarkTranslation.Description,
		&lw.LandmarkTranslation.History,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &lw, nil
}
