package model

// Landmark represents a point of interest in Heidelberg.
type Landmark struct {
	ID            int64
	Latitude      float64
	Longitude     float64
	ImageFilename string
	YearBuilt     int
	YearDestroyed *int // nullable — nil means still standing
}

// LandmarkTranslation holds locale-specific content for a landmark.
type LandmarkTranslation struct {
	ID          int64
	LandmarkID  int64
	Locale      string
	Name        string
	Description string
	History     string
}

// LandmarkWithTranslation is the joined view used by handlers.
type LandmarkWithTranslation struct {
	Landmark
	LandmarkTranslation
}
