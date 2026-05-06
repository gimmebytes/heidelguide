package store

import (
	"database/sql"

	"github.com/heidelguide/heidelguide/internal/model"
)

// ListLandmarks returns all landmarks joined with their translations for the given locale.
func (s *Store) ListLandmarks(locale string) ([]model.LandmarkWithTranslation, error) {
	rows, err := s.db.Query(`
		SELECT l.id, l.latitude, l.longitude, l.image_filename, l.year_built, l.year_destroyed,
		       t.id, t.landmark_id, t.locale, t.name, t.description, t.history
		FROM landmarks l
		JOIN landmark_translations t ON l.id = t.landmark_id
		WHERE t.locale = ?`, locale)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var landmarks []model.LandmarkWithTranslation
	for rows.Next() {
		var lw model.LandmarkWithTranslation
		if err := rows.Scan(
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
		); err != nil {
			return nil, err
		}
		landmarks = append(landmarks, lw)
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
