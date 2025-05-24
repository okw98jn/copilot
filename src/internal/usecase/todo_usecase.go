package usecase

import (
	"context"
	"copilot/internal/domain/repository"
	"copilot/internal/usecase/dto"
)

// TodoUseCase defines the interface for todo use cases
type TodoUseCase interface {
	// GetAllTodos returns all todo items
	GetAllTodos(ctx context.Context) ([]*dto.TodoResponse, error)
}

// todoUseCase implements the TodoUseCase interface
type todoUseCase struct {
	todoRepo repository.TodoRepository
}

// NewTodoUseCase creates a new todo use case instance
func NewTodoUseCase(todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}

// GetAllTodos returns all todo items
func (uc *todoUseCase) GetAllTodos(ctx context.Context) ([]*dto.TodoResponse, error) {
	todos, err := uc.todoRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.TodoResponse
	for _, todo := range todos {
		responses = append(responses, dto.NewTodoResponse(todo))
	}

	return responses, nil
}