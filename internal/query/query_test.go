package query

import (
	"path/filepath"
	"testing"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/storage"
)

func TestSearch(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer s.Close()
	r, _ := domain.NewRecord("a", "Dock", "Li", []domain.Item{{ID: "i", Area: "x", Description: "y", Points: 3}})
	s.SaveRecord(r)
	v, e := Search(s, Filter{Site: "dock"})
	if e != nil || len(v) != 1 {
		t.Fatal(v, e)
	}
}
