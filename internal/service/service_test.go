package service

import (
	"path/filepath"
	"testing"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/storage"
)

func svc(t *testing.T) *Service {
	s, e := storage.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return &Service{Store: s, Clock: FixedClock{Value: "t"}}
}
func TestCreateTransitions(t *testing.T) {
	s := svc(t)
	r, e := s.Create("r", "dock", "Li", []domain.Item{{ID: "i", Area: "sort", Description: "x", Points: 1}})
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.Review(r.ID); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Approve(r.ID); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Archive(r.ID); e != nil {
		t.Fatal(e)
	}
}
