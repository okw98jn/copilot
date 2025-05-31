package usecase

import (
	"copilot/internal/domain/entity"
	"copilot/internal/domain/repository"
)

// UserUseCase はユーザーに関するビジネスロジックを定義します
type UserUseCase interface {
	GetAllUsers() ([]*entity.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase は新しいUserUseCaseを作成します
func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

// GetAllUsers は全てのユーザーを取得します
func (uu *userUseCase) GetAllUsers() ([]*entity.User, error) {
	return uu.userRepository.FindAll()
}
