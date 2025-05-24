package entity

import (
	"time"
)

// Todo represents a todo item entity
type Todo struct {
	ID          int64     `bun:"id,pk,autoincrement"`
	Title       string    `bun:"title,notnull"`
	Description string    `bun:"description"`
	Completed   bool      `bun:"completed,notnull,default:false"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}