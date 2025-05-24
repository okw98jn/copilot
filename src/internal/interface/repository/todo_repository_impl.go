package repository

import (
	"context"
	"copilot/internal/domain/entity"
	"copilot/internal/domain/repository"

	"github.com/uptrace/bun"
)

// todoRepositoryImpl implements the TodoRepository interface
type todoRepositoryImpl struct {
	db *bun.DB
}

// NewTodoRepository creates a new todo repository implementation
func NewTodoRepository(db *bun.DB) repository.TodoRepository {
	return &todoRepositoryImpl{
		db: db,
	}
}

// FindAll returns all todo items
func (r *todoRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Todo, error) {
	var todos []*entity.Todo
	err := r.db.NewSelect().
		Model(&todos).
		OrderExpr("created_at DESC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}