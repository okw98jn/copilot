package repository

import (
	"context"
	"strconv"

	"copilot/internal/domain/entity"
	"copilot/internal/domain/repository"
	"copilot/internal/infrastructure/db/model"

	"github.com/uptrace/bun"
)

type userRepository struct {
	db *bun.DB
}

// NewUserRepository は新しいUserRepositoryを作成します
func NewUserRepository(db *bun.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// FindAll は全てのユーザーを取得します
func (ur *userRepository) FindAll() ([]*entity.User, error) {
	var userModels []model.UserModel

	err := ur.db.NewSelect().
		Model(&userModels).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	// モデルからエンティティへの変換
	users := make([]*entity.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = &entity.User{
			ID:    strconv.FormatInt(userModel.ID, 10),
			Name:  userModel.Name,
			Email: userModel.Email,
		}
	}

	return users, nil
}
