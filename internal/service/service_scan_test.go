package service

import (
	"path/filepath"
	"testing"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/storage"
)

// TestScanPersistsDeduction verifies that scanning a record's 5S remediation
// items saves the deduction score, so a later read reflects the scanned state.
func TestScanPersistsDeduction(t *testing.T) {
	st, e := storage.Open(filepath.Join(t.TempDir(), "t.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer st.Close()

	svc := &Service{Store: st, Clock: FixedClock{Value: "2026-01-01T00:00:00Z"}}

	items := []domain.Item{
		{ID: "S001", Area: "area-1", Description: "sort", Complete: false, Points: 1},
		{ID: "S002", Area: "area-2", Description: "set", Complete: false, Points: 2},
	}
	r, e := svc.Create("R1", "site-a", "inspector-a", items)
	if e != nil {
		t.Fatal(e)
	}

	// Stale in-memory score before scan: Create already recalculated, so the
	// record on disk has the correct deduction. To exercise the persist path,
	// corrupt the stored score, then scan should correct and save it.
	r.Score = 0
	if e := st.SaveRecord(r); e != nil {
		t.Fatal(e)
	}

	res, e := svc.Scan("R1")
	if e != nil {
		t.Fatal(e)
	}
	if res.Score != 3 {
		t.Fatalf("expected scanned deduction 3, got %d", res.Score)
	}
	if len(res.Pending) != 2 {
		t.Fatalf("expected 2 pending remediation items, got %d", len(res.Pending))
	}

	// Reload from storage to confirm the deduction was saved.
	got, e := st.GetRecord("R1")
	if e != nil {
		t.Fatal(e)
	}
	if got.Score != 3 {
		t.Fatalf("stored score = %d, want 3 after scan", got.Score)
	}
}

// TestScanOnlyPendingItems verifies that a completed item is dropped from the
// pending list after a scan.
func TestScanOnlyPendingItems(t *testing.T) {
	st, e := storage.Open(filepath.Join(t.TempDir(), "t.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer st.Close()

	svc := &Service{Store: st, Clock: FixedClock{Value: "2026-01-01T00:00:00Z"}}

	items := []domain.Item{
		{ID: "S001", Area: "area-1", Description: "sort", Complete: false, Points: 1},
		{ID: "S002", Area: "area-2", Description: "set", Complete: true, Points: 2},
	}
	if _, e := svc.Create("R2", "site-a", "inspector-a", items); e != nil {
		t.Fatal(e)
	}

	res, e := svc.Scan("R2")
	if e != nil {
		t.Fatal(e)
	}
	if len(res.Pending) != 1 {
		t.Fatalf("expected 1 pending item (S002 completed), got %d", len(res.Pending))
	}
	if res.Pending[0].ID != "S001" {
		t.Fatalf("expected pending S001, got %s", res.Pending[0].ID)
	}
	if res.Score != 1 {
		t.Fatalf("expected deduction 1, got %d", res.Score)
	}
}
