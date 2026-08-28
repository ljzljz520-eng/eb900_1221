package review

import "warehouse5s/internal/domain"

func CompleteAll(r *domain.Record) error {
	for _, i := range r.IncompleteItems() {
		if e := r.UpdateItem(i.ID, true); e != nil {
			return e
		}
	}
	return nil
}
func CompletionRate(r domain.Record) float64 {
	if len(r.Items) == 0 {
		return 0
	}
	return float64(len(r.Items)-len(r.IncompleteItems())) / float64(len(r.Items))
}
