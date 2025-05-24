package repository

import (
	"context"
	"copilot/internal/domain/entity"
	"copilot/internal/domain/repository"

	"github.com/uptrace/bun"
)

// todoRepositoryImpl はTodoRepositoryインターフェースを実装する
type todoRepositoryImpl struct {
	db *bun.DB
}

// NewTodoRepository は新しいTodoリポジトリの実装を作成する
func NewTodoRepository(db *bun.DB) repository.TodoRepository {
	return &todoRepositoryImpl{
		db: db,
	}
}

// FindAll は全てのTodoアイテムを返す
func (r *todoRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Todo, error) {
	var bunTodos []*BunTodo
	err := r.db.NewSelect().
		Model(&bunTodos).
		OrderExpr("created_at DESC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	// BunTodoスライスをドメインエンティティスライスに変換する
	todos := make([]*entity.Todo, len(bunTodos))
	for i, bunTodo := range bunTodos {
		todos[i] = bunTodo.ToDomainEntity()
	}

	return todos, nil
}