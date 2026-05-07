package model

// Landmark represents a point of interest in Heidelberg.
type Landmark struct {
	ID            int64
	Latitude      float64
	Longitude     float64
	ImageFilename string
	YearBuilt     int
	YearDestroyed *int  // nullable — nil means still standing
	CategoryID    int64 // FK to categories table
	Highlighted   bool  // true = show "Highlight" badge on card
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

// LandmarkWithTranslation is the joined view used by handlers (kept for backward compatibility).
type LandmarkWithTranslation struct {
	Landmark
	LandmarkTranslation
}

// Category represents a landmark classification (e.g., Architecture, Nature).
type Category struct {
	ID    int64
	Slug  string // machine-readable identifier: "architecture", "nature", "history", "culture"
	Color string // Tailwind color class for the pill: "rose", "emerald", "amber", "violet"
}

// CategoryTranslation holds locale-specific name for a category.
type CategoryTranslation struct {
	ID         int64
	CategoryID int64
	Locale     string
	Name       string
}

// LandmarkView is the full view used by templates (replaces LandmarkWithTranslation for new code).
type LandmarkView struct {
	Landmark
	LandmarkTranslation
	CategoryName  string // translated category name for current locale
	CategorySlug  string // for color mapping
	CategoryColor string // Tailwind color token
}
