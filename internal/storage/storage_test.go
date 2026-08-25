package storage

import (
	"path/filepath"
	"testing"
	"warehouse5s/internal/domain"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "x.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r, _ := domain.NewRecord("r", "dock", "Li", []domain.Item{{ID: "i", Area: "set", Description: "x", Points: 1}})
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("r")
	if e != nil || got.ID != "r" {
		t.Fatal(got, e)
	}
}
