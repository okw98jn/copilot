package dto

import (
	"copilot/internal/domain/entity"
	"time"
)

// TodoResponse はTodoアイテムのレスポンスを表す
type TodoResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FromEntity はTodoエンティティをTodoResponse DTOに変換する
func (t *TodoResponse) FromEntity(todo *entity.Todo) {
	t.ID = todo.ID
	t.Title = todo.Title
	t.Description = todo.Description
	t.Completed = todo.Completed
	t.CreatedAt = todo.CreatedAt
	t.UpdatedAt = todo.UpdatedAt
}

// NewTodoResponse はTodoエンティティから新しいTodoResponseを作成する
func NewTodoResponse(todo *entity.Todo) *TodoResponse {
	resp := &TodoResponse{}
	resp.FromEntity(todo)
	return resp
}