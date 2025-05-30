package repository

import "copilot/internal/domain/entity"

type UserRepository interface {
	FindAll() ([]*entity.User, error)
}
