package store

import "testing"

func TestUpsertRating_Insert(t *testing.T) {
	s := setupTestStore(t)

	if err := s.UpsertRating(1, "device-aaa", 4); err != nil {
		t.Fatalf("UpsertRating failed: %v", err)
	}

	summary, err := s.GetRatingSummary(1, "device-aaa")
	if err != nil {
		t.Fatalf("GetRatingSummary failed: %v", err)
	}
	if summary.Count != 1 {
		t.Errorf("expected count 1, got %d", summary.Count)
	}
	if summary.Average != 4.0 {
		t.Errorf("expected average 4.0, got %f", summary.Average)
	}
	if summary.UserScore != 4 {
		t.Errorf("expected user score 4, got %d", summary.UserScore)
	}
}

func TestUpsertRating_Update(t *testing.T) {
	s := setupTestStore(t)

	s.UpsertRating(1, "device-bbb", 3)
	s.UpsertRating(1, "device-bbb", 5) // update

	summary, _ := s.GetRatingSummary(1, "device-bbb")
	if summary.UserScore != 5 {
		t.Errorf("expected updated score 5, got %d", summary.UserScore)
	}
	if summary.Count != 1 {
		t.Errorf("expected count 1 (no duplicate), got %d", summary.Count)
	}
}

func TestGetRatingSummary_MultipleDevices(t *testing.T) {
	s := setupTestStore(t)

	s.UpsertRating(2, "device-aaa", 2)
	s.UpsertRating(2, "device-bbb", 4)
	s.UpsertRating(2, "device-ccc", 3)

	summary, err := s.GetRatingSummary(2, "device-aaa")
	if err != nil {
		t.Fatalf("GetRatingSummary failed: %v", err)
	}
	if summary.Count != 3 {
		t.Errorf("expected count 3, got %d", summary.Count)
	}
	if summary.Average != 3.0 {
		t.Errorf("expected average 3.0, got %f", summary.Average)
	}
	if summary.UserScore != 2 {
		t.Errorf("expected user score 2 for device-aaa, got %d", summary.UserScore)
	}
}

func TestGetRatingSummary_NoRatings(t *testing.T) {
	s := setupTestStore(t)

	summary, err := s.GetRatingSummary(1, "device-zzz")
	if err != nil {
		t.Fatalf("GetRatingSummary failed: %v", err)
	}
	if summary.Count != 0 {
		t.Errorf("expected count 0, got %d", summary.Count)
	}
	if summary.Average != 0.0 {
		t.Errorf("expected average 0.0, got %f", summary.Average)
	}
	if summary.UserScore != 0 {
		t.Errorf("expected user score 0, got %d", summary.UserScore)
	}
}
