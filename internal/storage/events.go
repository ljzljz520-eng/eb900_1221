package storage

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"warehouse5s/internal/domain"
)

func (s *Store) Events(recordID string) ([]domain.AuditEvent, error) {
	out := []domain.AuditEvent{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(eventsBucket).ForEach(func(_, v []byte) error {
			var x domain.AuditEvent
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			if x.RecordID == recordID {
				out = append(out, x)
			}
			return nil
		})
	})
	return out, e
}
