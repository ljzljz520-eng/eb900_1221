package domain

import (
	"errors"
	"fmt"
)

type Item struct {
	ID          string `json:"id"`
	Area        string `json:"area"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Points      int    `json:"points"`
}
type Record struct {
	ID        string `json:"id"`
	Site      string `json:"site"`
	Inspector string `json:"inspector"`
	Status    string `json:"status"`
	Score     int    `json:"score"`
	Items     []Item `json:"items"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type AuditEvent struct {
	ID       string
	RecordID string
	Kind     string
	Message  string
	At       string
}
type Workflow struct {
	ID       string
	RecordID string
	State    string
	At       string
}
type Attachment struct {
	ID       string
	RecordID string
	Name     string
	Digest   string
}

var statuses = map[string]bool{"draft": true, "review": true, "approved": true, "archived": true}

func NewRecord(id, site, inspector string, items []Item) (Record, error) {
	if id == "" || site == "" || inspector == "" {
		return Record{}, errors.New("id, site and inspector required")
	}
	if len(items) == 0 {
		return Record{}, errors.New("items required")
	}
	r := Record{ID: id, Site: site, Inspector: inspector, Status: "draft", Items: items}
	r.Recalculate()
	return r, nil
}
func (r *Record) Recalculate() {
	r.Score = 0
	for _, i := range r.Items {
		if !i.Complete {
			r.Score += i.Points
		}
	}
}
func (r Record) Validate() error {
	if r.ID == "" {
		return errors.New("missing id")
	}
	if !statuses[r.Status] {
		return fmt.Errorf("invalid status %s", r.Status)
	}
	if len(r.Items) == 0 {
		return errors.New("no items")
	}
	return nil
}
func (r *Record) Transition(next string) error {
	if !statuses[next] {
		return errors.New("unknown status")
	}
	allowed := map[string][]string{"draft": {"review"}, "review": {"approved", "draft"}, "approved": {"archived"}, "archived": {}}
	ok := false
	for _, v := range allowed[r.Status] {
		if v == next {
			ok = true
		}
	}
	if !ok {
		return fmt.Errorf("cannot transition %s to %s", r.Status, next)
	}
	r.Status = next
	return nil
}
func (r *Record) UpdateItem(id string, complete bool) error {
	for n := range r.Items {
		if r.Items[n].ID == id {
			r.Items[n].Complete = complete
			r.Recalculate()
			return nil
		}
	}
	return errors.New("item not found")
}
func (r Record) IncompleteItems() []Item {
	out := make([]Item, 0)
	for _, i := range r.Items {
		if !i.Complete {
			out = append(out, i)
		}
	}
	return out
}
func (r Record) IsArchived() bool { return r.Status == "archived" }
func (r Record) IsApproved() bool { return r.Status == "approved" }
func (r Record) Summary() string  { return fmt.Sprintf("%s:%s score=%d", r.ID, r.Site, r.Score) }
