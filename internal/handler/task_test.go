package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adjendrzcCG/test2/internal/handler"
	"github.com/adjendrzcCG/test2/internal/model"
	"github.com/adjendrzcCG/test2/internal/store"
)

func newHandler() http.Handler {
	return handler.Tasks(store.New())
}

func TestCreateAndListTasks(t *testing.T) {
	h := newHandler()

	body := `{"title":"Buy milk","description":"2% milk"}`
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(body)))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}

	var created model.Task
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if created.ID != 1 {
		t.Errorf("expected id=1, got %d", created.ID)
	}
	if created.Title != "Buy milk" {
		t.Errorf("expected title='Buy milk', got %q", created.Title)
	}

	rec2 := httptest.NewRecorder()
	h.ServeHTTP(rec2, httptest.NewRequest(http.MethodGet, "/tasks", nil))
	if rec2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec2.Code)
	}

	var list []*model.Task
	if err := json.NewDecoder(rec2.Body).Decode(&list); err != nil {
		t.Fatalf("decode list: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 task, got %d", len(list))
	}
}

func TestGetTask(t *testing.T) {
	h := newHandler()

	h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/tasks",
		bytes.NewBufferString(`{"title":"Read book"}`)))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tasks/1", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var task model.Task
	if err := json.NewDecoder(rec.Body).Decode(&task); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if task.Title != "Read book" {
		t.Errorf("expected 'Read book', got %q", task.Title)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	h := newHandler()

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tasks/99", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	h := newHandler()

	h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/tasks",
		bytes.NewBufferString(`{"title":"Write tests"}`)))

	body := `{"title":"Write more tests","description":"cover edge cases","done":true}`
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBufferString(body)))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var task model.Task
	if err := json.NewDecoder(rec.Body).Decode(&task); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if task.Title != "Write more tests" {
		t.Errorf("expected updated title, got %q", task.Title)
	}
	if !task.Done {
		t.Error("expected task to be done")
	}
}

func TestDeleteTask(t *testing.T) {
	h := newHandler()

	h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/tasks",
		bytes.NewBufferString(`{"title":"Clean desk"}`)))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodDelete, "/tasks/1", nil))
	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", rec.Code)
	}

	rec2 := httptest.NewRecorder()
	h.ServeHTTP(rec2, httptest.NewRequest(http.MethodGet, "/tasks/1", nil))
	if rec2.Code != http.StatusNotFound {
		t.Fatalf("expected 404 after delete, got %d", rec2.Code)
	}
}

func TestCreateTaskMissingTitle(t *testing.T) {
	h := newHandler()

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/tasks",
		bytes.NewBufferString(`{"description":"no title"}`)))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	h := newHandler()

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodPatch, "/tasks", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rec.Code)
	}
}
