package api

import (
	uuid "github.com/satori/go.uuid"
)

// Storage - in-memory simple storage
type Storage struct {
	data []Answer
}

// NewStorage returns storage client
func NewStorage() *Storage {
	return &Storage{}
}

// Save answer
func (s *Storage) Save(a Answer) {
	a.ID = uuid.NewV4().String()
	s.data = append(s.data, a)
}

// Get answers
func (s *Storage) Get() []Answer {
	return s.data
}
