package api

import (
	"net/http"
	"warehouse5s/internal/domain"
)

func method(w http.ResponseWriter, r *http.Request, allowed string) bool {
	if r.Method != allowed {
		w.Header().Set("Allow", allowed)
		http.Error(w, "method not allowed", 405)
		return false
	}
	return true
}
func validItems(items []domain.Item) bool {
	for _, i := range items {
		if !domain.ValidItem(i) {
			return false
		}
	}
	return len(items) > 0
}
