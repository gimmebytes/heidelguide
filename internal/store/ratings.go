package store

// RatingSummary holds aggregate rating data for a landmark.
type RatingSummary struct {
	Average   float64
	Count     int
	UserScore int // 0 means no rating from this device
}

// GetRatingSummary returns avg, count, and the device's own score for a landmark.
func (s *Store) GetRatingSummary(landmarkID int64, deviceID string) (*RatingSummary, error) {
	var summary RatingSummary
	err := s.db.QueryRow(`
		SELECT
			COALESCE(AVG(score), 0),
			COUNT(*),
			COALESCE((SELECT score FROM ratings WHERE landmark_id = ? AND device_id = ?), 0)
		FROM ratings WHERE landmark_id = ?`,
		landmarkID, deviceID, landmarkID,
	).Scan(&summary.Average, &summary.Count, &summary.UserScore)
	if err != nil {
		return nil, err
	}
	return &summary, nil
}

// GetAllRatingSummaries returns average and count for all landmarks that have ratings.
func (s *Store) GetAllRatingSummaries() (map[int64]*RatingSummary, error) {
	rows, err := s.db.Query(`
		SELECT landmark_id, AVG(score), COUNT(*)
		FROM ratings GROUP BY landmark_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int64]*RatingSummary)
	for rows.Next() {
		var id int64
		var rs RatingSummary
		if err := rows.Scan(&id, &rs.Average, &rs.Count); err != nil {
			return nil, err
		}
		result[id] = &rs
	}
	return result, rows.Err()
}

// UpsertRating creates or updates a rating for a device+landmark pair.
func (s *Store) UpsertRating(landmarkID int64, deviceID string, score int) error {
	_, err := s.db.Exec(`
		INSERT INTO ratings (landmark_id, device_id, score)
		VALUES (?, ?, ?)
		ON CONFLICT(landmark_id, device_id)
		DO UPDATE SET score = excluded.score, updated_at = datetime('now')`,
		landmarkID, deviceID, score,
	)
	return err
}
