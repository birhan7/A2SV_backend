package models

import (
	"time"
)

// Task represents a task with its properties.
type Task struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
}
