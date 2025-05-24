package usecase

import (
	"context"
	"copilot/internal/domain/repository"
	"copilot/internal/usecase/dto"
)

// TodoUseCase はTodoユースケースのインターフェースを定義する
type TodoUseCase interface {
	// GetAllTodos は全てのTodoアイテムを返す
	GetAllTodos(ctx context.Context) ([]*dto.TodoResponse, error)
}

// todoUseCase はTodoUseCaseインターフェースを実装する
type todoUseCase struct {
	todoRepo repository.TodoRepository
}

// NewTodoUseCase は新しいTodoユースケースのインスタンスを作成する
func NewTodoUseCase(todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}

// GetAllTodos は全てのTodoアイテムを返す
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