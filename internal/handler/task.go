// Package handler provides HTTP handlers for the task API.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/adjendrzcCG/test2/internal/store"
)

type createRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type updateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Tasks returns an http.Handler that handles task CRUD operations.
// Routes:
//
//	GET    /tasks       – list all tasks
//	POST   /tasks       – create a task
//	GET    /tasks/{id}  – get a task
//	PUT    /tasks/{id}  – update a task
//	DELETE /tasks/{id}  – delete a task
func Tasks(s *store.Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listTasks(w, s)
		case http.MethodPost:
			createTask(w, r, s)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseID(r.URL.Path, "/tasks/")
		if err != nil {
			http.Error(w, "invalid task id", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			getTask(w, id, s)
		case http.MethodPut:
			updateTask(w, r, id, s)
		case http.MethodDelete:
			deleteTask(w, id, s)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}

func listTasks(w http.ResponseWriter, s *store.Store) {
	tasks := s.List()
	writeJSON(w, http.StatusOK, tasks)
}

func createTask(w http.ResponseWriter, r *http.Request, s *store.Store) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(req.Title) == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	task := s.Create(req.Title, req.Description)
	writeJSON(w, http.StatusCreated, task)
}

func getTask(w http.ResponseWriter, id int, s *store.Store) {
	task, err := s.Get(id)
	if errors.Is(err, store.ErrNotFound) {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, task)
}

func updateTask(w http.ResponseWriter, r *http.Request, id int, s *store.Store) {
	var req updateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(req.Title) == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	task, err := s.Update(id, req.Title, req.Description, req.Done)
	if errors.Is(err, store.ErrNotFound) {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, task)
}

func deleteTask(w http.ResponseWriter, id int, s *store.Store) {
	if err := s.Delete(id); errors.Is(err, store.ErrNotFound) {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func parseID(path, prefix string) (int, error) {
	raw := strings.TrimPrefix(path, prefix)
	raw = strings.TrimSuffix(raw, "/")
	return strconv.Atoi(raw)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
