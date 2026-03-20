// Package store provides an in-memory store for tasks.
package store

import (
	"errors"
	"sync"
	"time"

	"github.com/adjendrzcCG/test2/internal/model"
)

// ErrNotFound is returned when a task is not found.
var ErrNotFound = errors.New("task not found")

// Store is an in-memory store for tasks.
type Store struct {
	mu     sync.RWMutex
	tasks  map[int]*model.Task
	nextID int
}

// New creates a new Store.
func New() *Store {
	return &Store{
		tasks:  make(map[int]*model.Task),
		nextID: 1,
	}
}

// Create adds a new task to the store and returns it.
func (s *Store) Create(title, description string) *model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &model.Task{
		ID:          s.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now().UTC(),
	}
	s.tasks[s.nextID] = task
	s.nextID++
	return task
}

// Get retrieves a task by ID. Returns ErrNotFound if the task does not exist.
func (s *Store) Get(id int) (*model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.tasks[id]
	if !ok {
		return nil, ErrNotFound
	}
	return task, nil
}

// List returns all tasks in the store.
func (s *Store) List() []*model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*model.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

// Update modifies an existing task. Returns ErrNotFound if the task does not exist.
func (s *Store) Update(id int, title, description string, done bool) (*model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return nil, ErrNotFound
	}
	task.Title = title
	task.Description = description
	task.Done = done
	return task, nil
}

// Delete removes a task by ID. Returns ErrNotFound if the task does not exist.
func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}
	delete(s.tasks, id)
	return nil
}
