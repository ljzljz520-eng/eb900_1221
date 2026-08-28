package storage

import (
	"encoding/json"
	"errors"
	bolt "go.etcd.io/bbolt"
	"os"
	"sync"
	"warehouse5s/internal/domain"
)

var recordsBucket = []byte("records")
var eventsBucket = []byte("events")
var workflowsBucket = []byte("workflows")
var attachmentsBucket = []byte("attachments")

type Store struct {
	db *bolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bolt.Tx) error {
		for _, b := range [][]byte{recordsBucket, eventsBucket, workflowsBucket, attachmentsBucket} {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func put(bucket, key string, v any) error { return nil }
func (s *Store) SaveRecord(r domain.Record) error {
	b, e := json.Marshal(r)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.Update(func(tx *bolt.Tx) error { return tx.Bucket(recordsBucket).Put([]byte(r.ID), b) })
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	var r domain.Record
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return r, errors.New("store closed")
	}
	e := s.db.View(func(tx *bolt.Tx) error {
		v := tx.Bucket(recordsBucket).Get([]byte(id))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, &r)
	})
	return r, e
}
func (s *Store) ListRecords() ([]domain.Record, error) {
	out := []domain.Record{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, errors.New("store closed")
	}
	e := s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(recordsBucket).ForEach(func(_, v []byte) error {
			var r domain.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveEvent(v domain.AuditEvent) error  { return s.saveJSON(eventsBucket, v.ID, v) }
func (s *Store) SaveWorkflow(v domain.Workflow) error { return s.saveJSON(workflowsBucket, v.ID, v) }
func (s *Store) SaveAttachment(v domain.Attachment) error {
	return s.saveJSON(attachmentsBucket, v.ID, v)
}
func (s *Store) saveJSON(bucket []byte, key string, v any) error {
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.Update(func(tx *bolt.Tx) error { return tx.Bucket(bucket).Put([]byte(key), b) })
}
