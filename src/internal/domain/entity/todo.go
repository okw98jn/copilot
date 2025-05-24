package entity

import (
	"time"
)

// Todo represents a todo item entity in the domain layer
type Todo struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}