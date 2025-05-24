package repository

import (
	"context"
	"copilot/internal/domain/entity"
)

// TodoRepository defines the interface for todo repository operations
type TodoRepository interface {
	// FindAll returns all todo items
	FindAll(ctx context.Context) ([]*entity.Todo, error)
}