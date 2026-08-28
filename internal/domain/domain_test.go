package domain

import "testing"

func TestNewRecord(t *testing.T) {
	r, e := NewRecord("1", "A", "Bob", []Item{{ID: "i", Area: "sort", Description: "clear", Points: 2}})
	if e != nil || r.Score != 2 {
		t.Fatal(r, e)
	}
}
