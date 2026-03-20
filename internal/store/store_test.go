package store_test

import (
	"testing"

	"github.com/adjendrzcCG/test2/internal/store"
)

func TestCreateAndGet(t *testing.T) {
	s := store.New()

	task := s.Create("Buy coffee", "dark roast")
	if task.ID != 1 {
		t.Errorf("expected id=1, got %d", task.ID)
	}

	got, err := s.Get(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Title != "Buy coffee" {
		t.Errorf("expected 'Buy coffee', got %q", got.Title)
	}
}

func TestGetNotFound(t *testing.T) {
	s := store.New()

	_, err := s.Get(42)
	if err != store.ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestList(t *testing.T) {
	s := store.New()
	s.Create("Task 1", "")
	s.Create("Task 2", "")

	tasks := s.List()
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestUpdate(t *testing.T) {
	s := store.New()
	s.Create("Old title", "")

	task, err := s.Update(1, "New title", "new desc", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.Title != "New title" {
		t.Errorf("expected 'New title', got %q", task.Title)
	}
	if !task.Done {
		t.Error("expected task to be done")
	}
}

func TestUpdateNotFound(t *testing.T) {
	s := store.New()

	_, err := s.Update(99, "x", "", false)
	if err != store.ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestDelete(t *testing.T) {
	s := store.New()
	s.Create("To delete", "")

	if err := s.Delete(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := s.Get(1)
	if err != store.ErrNotFound {
		t.Errorf("expected ErrNotFound after delete, got %v", err)
	}
}

func TestDeleteNotFound(t *testing.T) {
	s := store.New()

	if err := s.Delete(99); err != store.ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}
