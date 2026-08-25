package review

import (
	"testing"
	"warehouse5s/internal/domain"
)

func TestInspect(t *testing.T) {
	r, _ := domain.NewRecord("a", "d", "i", []domain.Item{{ID: "x", Area: "a", Description: "b", Points: 2}})
	if len(Inspect(r).Pending) != 1 {
		t.Fatal()
	}
}
