package storage

import "fmt"

// Storage represents the simple storage interface
type Storage interface {
	Add(key, value string) string
	Get(key string) (string, error)
	List() ItemType
	Delete(key string) error
}

type storage struct {
	s ItemType
}

// ItemType is a general type of the data to store
type ItemType map[string]string

// New returns a new instance of the Storage interface.
func New() Storage {
	return &storage{
		s: make(ItemType),
	}
}

func (s *storage) Add(key, value string) string {
	s.s[key] = value
	return key
}

func (s *storage) Get(key string) (string, error) {
	v, ok := s.s[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}
	return v, nil
}

func (s *storage) List() ItemType {
	copied := make(ItemType, len(s.s))
	for k, v := range s.s {
		copied[k] = v
	}
	return copied
}

func (s *storage) Delete(key string) error {
	delete(s.s, key)
	return nil
}
