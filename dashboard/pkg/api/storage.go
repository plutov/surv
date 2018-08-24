package api

import (
	"sync"
)

// Storage - in-memory simple storage
type Storage struct {
	data []DashboardRow
	ids  map[string]bool
	mu   *sync.Mutex
}

// NewStorage returns storage client
func NewStorage() *Storage {
	return &Storage{
		ids: make(map[string]bool),
		mu:  &sync.Mutex{},
	}
}

// Save DashboardRow
func (s *Storage) Save(r DashboardRow) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.ids[r.AnswerID]
	if !exists {
		s.data = append(s.data, r)
		s.ids[r.AnswerID] = true
	}
}

// Get DashboardRows
func (s *Storage) Get(limit int, offset int) []DashboardRow {
	s.mu.Lock()
	defer s.mu.Unlock()

	to := offset + limit
	if to > len(s.data) {
		to = len(s.data)
	}
	if offset >= len(s.data) {
		offset = len(s.data)
	}

	return s.data[offset:to]
}
