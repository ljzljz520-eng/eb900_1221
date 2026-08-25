package service

import (
	"errors"
	"fmt"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/storage"
)

type Clock interface{ Now() string }
type FixedClock struct{ Value string }

func (c FixedClock) Now() string { return c.Value }

type Service struct {
	Store *storage.Store
	Clock Clock
}

func (s *Service) Create(id, site, inspector string, items []domain.Item) (domain.Record, error) {
	if s.Clock == nil {
		s.Clock = FixedClock{Value: "2026-01-01T00:00:00Z"}
	}
	items = domain.NormalizeItems(items)
	r, e := domain.NewRecord(id, site, inspector, items)
	if e != nil {
		return r, e
	}
	r.CreatedAt = s.Clock.Now()
	r.UpdatedAt = r.CreatedAt
	e = s.Store.SaveRecord(r)
	if e == nil {
		e = s.Store.SaveEvent(domain.AuditEvent{ID: id + "-create", RecordID: id, Kind: "create", Message: "record created", At: r.CreatedAt})
	}
	return r, e
}
func (s *Service) Review(id string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if e = r.Transition("review"); e != nil {
		return r, e
	}
	r.UpdatedAt = s.Clock.Now()
	e = s.Store.SaveRecord(r)
	return r, e
}
func (s *Service) Approve(id string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if e = r.Transition("approved"); e != nil {
		return r, e
	}
	r.UpdatedAt = s.Clock.Now()
	return r, s.Store.SaveRecord(r)
}
func (s *Service) Archive(id string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if e = r.Transition("archived"); e != nil {
		return r, e
	}
	r.UpdatedAt = s.Clock.Now()
	return r, s.Store.SaveRecord(r)
}
func (s *Service) Update(id, item string, complete bool) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.IsArchived() {
		return r, errors.New("archived record immutable")
	}
	if e = r.UpdateItem(item, complete); e != nil {
		return r, e
	}
	r.UpdatedAt = s.Clock.Now()
	return r, s.Store.SaveRecord(r)
}
func (s *Service) Publish(id string) (domain.Record, error) {
	r, e := s.Approve(id)
	if e != nil {
		return r, e
	}
	return r, s.Store.SaveEvent(domain.AuditEvent{ID: fmt.Sprintf("%s-publish", id), RecordID: id, Kind: "publish", Message: r.Summary(), At: r.UpdatedAt})
}
