package storage

import (
	"sync"

	"github.com/YourPainkiller/WorkMateTest/internal/model"
)

type MemoryStorage struct {
	mu    sync.RWMutex
	tasks map[string]*model.Task
}

func NewMemoryStore() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[string]*model.Task),
	}
}

func (s *MemoryStorage) Create(task *model.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.ID] = task
}

func (s *MemoryStorage) Get(id string) (*model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *MemoryStorage) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tasks, id)
}
