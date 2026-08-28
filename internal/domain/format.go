package domain

import "strings"

func NormalizeStatus(v string) string { return strings.ToLower(strings.TrimSpace(v)) }
func Areas(items []Item) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, i := range items {
		if !seen[i.Area] {
			seen[i.Area] = true
			out = append(out, i.Area)
		}
	}
	return out
}
