package review

import (
	"errors"
	"warehouse5s/internal/domain"
)

type Result struct {
	RecordID string
	Pending  []domain.Item
	Score    int
	Ready    bool
}

func Inspect(r domain.Record) Result {
	pending := r.IncompleteItems()
	if r.Status == "draft" {
		pending = append([]domain.Item{}, r.Items...)
	}
	return Result{RecordID: r.ID, Pending: pending, Score: r.Score, Ready: len(pending) == 0}
}
func Apply(r *domain.Record, completed []string) error {
	if r == nil {
		return errors.New("nil record")
	}
	for _, id := range completed {
		if e := r.UpdateItem(id, true); e != nil {
			return e
		}
	}
	return nil
}
func RequireReview(r domain.Record) error {
	if r.Status != "review" {
		return errors.New("record is not in review")
	}
	return nil
}
func PendingIDs(r domain.Record) []string {
	out := []string{}
	for _, i := range r.IncompleteItems() {
		out = append(out, i.ID)
	}
	return out
}
