package repository

import (
	"context"
	"database/sql"
	"go_api/internal/model"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, model *model.UserModel) (uuid.UUID, int, error)
	GetRefreshToken(ctx context.Context, userId uuid.UUID, now time.Time) (*model.RefreshTokenModel, error)
	CreateRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error
	GetUserById(ctx context.Context, userId uuid.UUID) (*model.UserModel, error)
	DeleteRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (int64, error)
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
