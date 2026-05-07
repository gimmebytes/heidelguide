package store

import "testing"

func setupTestStore(t *testing.T) *Store {
	t.Helper()
	s, err := Open(":memory:")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	t.Cleanup(func() { s.Close() })
	if err := s.Migrate(); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	if err := s.Seed(); err != nil {
		t.Fatalf("failed to seed: %v", err)
	}
	if err := s.SeedCategories(); err != nil {
		t.Fatalf("failed to seed categories: %v", err)
	}
	if err := s.AssignDefaultCategories(); err != nil {
		t.Fatalf("failed to assign categories: %v", err)
	}
	return s
}

func TestListLandmarks_CategoryDataPopulated(t *testing.T) {
	s := setupTestStore(t)

	landmarks, err := s.ListLandmarks("de")
	if err != nil {
		t.Fatalf("ListLandmarks returned error: %v", err)
	}

	if len(landmarks) != 8 {
		t.Fatalf("expected 8 landmarks, got %d", len(landmarks))
	}

	// First landmark should be Heidelberger Schloss (ID 1, highlighted, ordered first).
	first := landmarks[0]
	if first.Landmark.ID != 1 {
		t.Errorf("expected first landmark ID=1, got %d", first.Landmark.ID)
	}
	if first.CategoryName != "ARCHITEKTUR" {
		t.Errorf("expected CategoryName 'ARCHITEKTUR', got %q", first.CategoryName)
	}
	if first.CategorySlug != "architecture" {
		t.Errorf("expected CategorySlug 'architecture', got %q", first.CategorySlug)
	}
	if first.CategoryColor != "rose" {
		t.Errorf("expected CategoryColor 'rose', got %q", first.CategoryColor)
	}

	// All landmarks must have non-empty category fields.
	for i, lv := range landmarks {
		if lv.CategoryName == "" {
			t.Errorf("landmark[%d] (ID=%d) has empty CategoryName", i, lv.Landmark.ID)
		}
		if lv.CategorySlug == "" {
			t.Errorf("landmark[%d] (ID=%d) has empty CategorySlug", i, lv.Landmark.ID)
		}
		if lv.CategoryColor == "" {
			t.Errorf("landmark[%d] (ID=%d) has empty CategoryColor", i, lv.Landmark.ID)
		}
	}
}

func TestListLandmarks_HighlightOrdering(t *testing.T) {
	s := setupTestStore(t)

	landmarks, err := s.ListLandmarks("de")
	if err != nil {
		t.Fatalf("ListLandmarks returned error: %v", err)
	}

	if len(landmarks) != 8 {
		t.Fatalf("expected 8 landmarks, got %d", len(landmarks))
	}

	// First 3 should be highlighted (IDs 1, 2, 5).
	for i := 0; i < 3; i++ {
		if !landmarks[i].Landmark.Highlighted {
			t.Errorf("expected landmarks[%d] (ID=%d) to be highlighted", i, landmarks[i].Landmark.ID)
		}
	}

	// Remaining 5 should not be highlighted.
	for i := 3; i < 8; i++ {
		if landmarks[i].Landmark.Highlighted {
			t.Errorf("expected landmarks[%d] (ID=%d) to NOT be highlighted", i, landmarks[i].Landmark.ID)
		}
	}
}

func TestGetLandmark_CategoryData(t *testing.T) {
	s := setupTestStore(t)

	lv, err := s.GetLandmark(1, "en")
	if err != nil {
		t.Fatalf("GetLandmark returned error: %v", err)
	}
	if lv == nil {
		t.Fatal("expected non-nil LandmarkView, got nil")
	}

	if lv.CategoryName != "ARCHITECTURE" {
		t.Errorf("expected CategoryName 'ARCHITECTURE', got %q", lv.CategoryName)
	}
	if lv.CategorySlug != "architecture" {
		t.Errorf("expected CategorySlug 'architecture', got %q", lv.CategorySlug)
	}
	if lv.CategoryColor != "rose" {
		t.Errorf("expected CategoryColor 'rose', got %q", lv.CategoryColor)
	}
	if !lv.Landmark.Highlighted {
		t.Error("expected Highlighted to be true")
	}
}

func TestGetLandmark_NotFound(t *testing.T) {
	s := setupTestStore(t)

	lv, err := s.GetLandmark(9999, "de")
	if err != nil {
		t.Fatalf("GetLandmark returned error: %v", err)
	}
	if lv != nil {
		t.Errorf("expected nil for non-existent landmark, got ID=%d", lv.Landmark.ID)
	}
}
