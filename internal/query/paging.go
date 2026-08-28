package query

import "warehouse5s/internal/domain"

func Page(records []domain.Record, offset, limit int) []domain.Record {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = len(records)
	}
	if offset >= len(records) {
		return []domain.Record{}
	}
	end := offset + limit
	if end > len(records) {
		end = len(records)
	}
	return records[offset:end]
}
func Statuses(records []domain.Record) map[string]int {
	out := map[string]int{}
	for _, r := range records {
		out[r.Status]++
	}
	return out
}
