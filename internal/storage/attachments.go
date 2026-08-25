package storage

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"warehouse5s/internal/domain"
)

func (s *Store) Attachments(recordID string) ([]domain.Attachment, error) {
	out := []domain.Attachment{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(attachmentsBucket).ForEach(func(_, v []byte) error {
			var x domain.Attachment
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
