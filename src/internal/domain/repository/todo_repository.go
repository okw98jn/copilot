package repository

import (
	"context"
	"copilot/internal/domain/entity"
)

// TodoRepository はTodoリポジトリ操作のためのインターフェースを定義する
type TodoRepository interface {
	// FindAll は全てのTodoアイテムを返す
	FindAll(ctx context.Context) ([]*entity.Todo, error)
}