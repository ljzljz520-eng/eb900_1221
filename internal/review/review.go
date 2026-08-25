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

// Inspect scans a record for its 5S remediation items. Pending only retains
// items that are still incomplete, so completed items are never reported as
// awaiting remediation. Score is recomputed from the current item state so a
// scan reflects the actual deduction rather than any stale stored value.
func Inspect(r domain.Record) Result {
	pending := r.IncompleteItems()
	return Result{RecordID: r.ID, Pending: pending, Score: domain.Score(r.Items), Ready: len(pending) == 0}
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
