package data

import "sync"

type Provider interface {
	Get(k string) ([]string, bool)
	Set(k string, v []string)
	Add(k string, v string)
	List() map[string][]string
}

type Storage struct {
	mu   sync.RWMutex
	data map[string][]string
}

func New() *Storage {
	return &Storage{
		mu:   sync.RWMutex{},
		data: make(map[string][]string),
	}
}

func (s *Storage) Get(k string) ([]string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.get(k)
}

func (s *Storage) Set(k string, v []string) {
	s.mu.Lock()
	s.set(k, v)
	s.mu.Unlock()
}

func (s *Storage) Add(k string, v string) {
	s.mu.Lock()
	s.add(k, v)
	s.mu.Unlock()
}

func (s *Storage) get(k string) ([]string, bool) {
	res, ok := s.data[k]
	return res, ok
}

func (s *Storage) set(k string, v []string) {
	s.data[k] = v
}

func (s *Storage) add(k string, v string) {
	if _, ok := s.data[k]; !ok {
		s.data[k] = make([]string, 0)
	}
	s.data[k] = append(s.data[k], v)
}

func (s *Storage) List() map[string][]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.list()
}

func (s *Storage) list() map[string][]string {
	res := s.data
	return res
}
