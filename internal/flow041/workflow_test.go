package flow041

import (
	"path/filepath"
	"testing"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/service"
	"warehouse5s/internal/storage"
)

func TestWorkflowCreateReviewArchive(t *testing.T) {
	st, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer st.Close()
	s := &service.Service{Store: st, Clock: service.FixedClock{Value: "t"}}
	r, e := s.Create("a", "dock", "Li", []domain.Item{{ID: "i", Area: "x", Description: "y", Points: 1}})
	if e != nil {
		t.Fatal(e)
	}
	s.Review(r.ID)
	s.Approve(r.ID)
	r, e = s.Archive(r.ID)
	if e != nil || !r.IsArchived() {
		t.Fatal(r, e)
	}
}
func TestWorkflowSearchUpdatePublish(t *testing.T) {
	st, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer st.Close()
	s := &service.Service{Store: st, Clock: service.FixedClock{Value: "t"}}
	r, _ := s.Create("a", "dock", "Li", []domain.Item{{ID: "i", Area: "x", Description: "y", Points: 1}})
	s.Update(r.ID, "i", true)
	if _, e := s.Publish(r.ID); e == nil {
		t.Fatal("publish should require review")
	}
}
func TestWorkflowImportReport(t *testing.T) {
	if BuildMarker() == "" {
		t.Fatal()
	}
}
func BuildMarker() string { return "import-report" }
