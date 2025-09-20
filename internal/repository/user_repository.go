package repository

import (
	"context"
	"database/sql"
	"go_api/internal/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, model *model.UserModel) (uuid.UUID, error)
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
