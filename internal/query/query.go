package query

import (
	"strings"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/storage"
)

type Filter struct {
	Site      string
	Inspector string
	Status    string
	MinScore  int
}

func Search(s *storage.Store, f Filter) ([]domain.Record, error) {
	all, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []domain.Record{}
	for _, r := range all {
		if f.Site != "" && !strings.Contains(strings.ToLower(r.Site), strings.ToLower(f.Site)) {
			continue
		}
		if f.Inspector != "" && r.Inspector != f.Inspector {
			continue
		}
		if f.Status != "" && r.Status != f.Status {
			continue
		}
		if r.Score < f.MinScore {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}
func ByID(s *storage.Store, id string) (domain.Record, error) { return s.GetRecord(id) }
func Count(records []domain.Record) int                       { return len(records) }
func SortByScore(records []domain.Record) []domain.Record {
	out := append([]domain.Record{}, records...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Score > out[i].Score {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
