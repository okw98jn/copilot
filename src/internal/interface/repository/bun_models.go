package repository

import (
	"copilot/internal/domain/entity"
	"time"
)

// BunTodo represents a todo item entity for Bun ORM
type BunTodo struct {
	ID          int64     `bun:"id,pk,autoincrement"`
	Title       string    `bun:"title,notnull"`
	Description string    `bun:"description"`
	Completed   bool      `bun:"completed,notnull,default:false"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

// ToDomainEntity converts a BunTodo to a domain entity Todo
func (b *BunTodo) ToDomainEntity() *entity.Todo {
	return &entity.Todo{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Completed:   b.Completed,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

// FromDomainEntity converts a domain entity Todo to BunTodo
func FromDomainEntity(todo *entity.Todo) *BunTodo {
	return &BunTodo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}