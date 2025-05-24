package entity

import (
	"time"
)

// Todo はドメイン層のTodoアイテムエンティティを表す
type Todo struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}