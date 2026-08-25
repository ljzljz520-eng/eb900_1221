package report

import "warehouse5s/internal/domain"

type Summary struct {
	Total    int
	Open     int
	Archived int
	Points   int
}

func Summarize(rs []domain.Record) Summary {
	out := Summary{Total: len(rs)}
	for _, r := range rs {
		if r.IsArchived() {
			out.Archived++
		} else {
			out.Open++
		}
		out.Points += r.Score
	}
	return out
}
func NeedsAttention(r domain.Record) bool { return r.Score > 0 && r.Status != "archived" }
