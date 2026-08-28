package service

import (
	"warehouse5s/internal/domain"
	"warehouse5s/internal/storage"
)

func SaveAttachment(st *storage.Store, a domain.Attachment) error { return st.SaveAttachment(a) }
func (s *Service) AddAudit(id, kind, message string) error {
	return s.Store.SaveEvent(domain.AuditEvent{ID: id + "-" + kind, RecordID: id, Kind: kind, Message: message, At: s.Clock.Now()})
}
func (s *Service) Snapshot(id string) (domain.Record, error) { return s.Store.GetRecord(id) }
