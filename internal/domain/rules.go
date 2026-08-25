package domain

import "errors"

func ValidateSite(site string) error {
	if len(site) < 2 {
		return errors.New("site too short")
	}
	return nil
}
func ValidateInspector(name string) error {
	if len(name) < 2 {
		return errors.New("inspector too short")
	}
	return nil
}
func ValidItem(i Item) bool {
	return i.ID != "" && i.Area != "" && i.Description != "" && i.Points >= 0
}
func NormalizeItems(items []Item) []Item {
	out := make([]Item, 0, len(items))
	seen := map[string]bool{}
	for _, i := range items {
		if ValidItem(i) && !seen[i.ID] {
			seen[i.ID] = true
			out = append(out, i)
		}
	}
	return out
}
func Score(items []Item) int {
	n := 0
	for _, i := range items {
		if !i.Complete {
			n += i.Points
		}
	}
	return n
}
