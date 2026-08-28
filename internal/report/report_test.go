package report

import (
	"path/filepath"
	"testing"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/service"
	"warehouse5s/internal/storage"
)

func TestImport(t *testing.T) {
	st, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer st.Close()
	s := &service.Service{Store: st, Clock: service.FixedClock{Value: "t"}}
	b := []byte(`[{"id":"a","site":"dock","inspector":"Li","items":[{"id":"i","area":"x","description":"y","points":1}]},{"site":"bad"}]`)
	r, e := Import(s, b)
	if e != nil || r.Accepted != 1 || r.Rejected != 1 {
		t.Fatal(r, e)
	}
}
func TestSummary(t *testing.T) {
	r, _ := domain.NewRecord("a", "d", "i", []domain.Item{{ID: "x", Area: "a", Description: "b", Points: 2}})
	if Summarize([]domain.Record{r}).Points != 2 {
		t.Fatal()
	}
}
