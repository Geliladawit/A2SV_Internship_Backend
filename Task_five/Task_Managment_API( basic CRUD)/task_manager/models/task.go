package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Status      string    `json:"status"`
}
