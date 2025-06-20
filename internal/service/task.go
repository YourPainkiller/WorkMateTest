package service

import (
	"context"
	"math/rand"

	"time"

	"github.com/YourPainkiller/WorkMateTest/internal/model"
	"github.com/google/uuid"
)

type IMemoryStorage interface {
	Create(*model.Task)
	Get(string) (*model.Task, bool)
	Delete(string)
}

type TaskService struct {
	store IMemoryStorage
}

func NewTaskService(store IMemoryStorage) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) CreateTask(ctx context.Context) *model.Task {
	task := &model.Task{
		ID:        uuid.New().String(),
		Status:    model.StatusCreated,
		CreatedAt: time.Now(),
	}

	s.store.Create(task)

	go s.processTask(task)

	return task
}

func (s *TaskService) processTask(task *model.Task) {
	task.Status = model.StatusRunning
	task.StartedAt = time.Now()

	// I/O bound simulation
	duration := time.Duration(rand.Intn(120)+180) * time.Second
	time.Sleep(duration)

	task.Status = model.StatusCompleted
	task.Result = "success"

	task.CompletedAt = time.Now()
}

func (s *TaskService) GetTask(id string) (*model.Task, error) {
	task, ok := s.store.Get(id)
	if !ok {
		return nil, ErrNoSuchTask
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {
	if _, ok := s.store.Get(id); !ok {
		return ErrNoSuchTask
	}

	s.store.Delete(id)
	return nil
}
