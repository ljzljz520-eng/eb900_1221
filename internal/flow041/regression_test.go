package flow041

import (
	"testing"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/review"
)

func Test900BusinessRegression(t *testing.T) {
	r, _ := domain.NewRecord("bug", "dock", "Li", []domain.Item{{ID: "done", Area: "sort", Description: "clear", Complete: true, Points: 1}, {ID: "open", Area: "set", Description: "label", Points: 2}})
	got := review.Inspect(r)
	if len(got.Pending) != 1 || got.Pending[0].ID != "open" {
		t.Fatalf("pending items: %#v", got.Pending)
	}
}
