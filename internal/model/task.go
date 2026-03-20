// Package model defines the data structures used in the application.
package model

import "time"

// Task represents a task with a title, description, and completion status.
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}
