package repository

import (
	"context"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (r *userRepository) CreateUser(ctx context.Context, model *model.UserModel) (uuid.UUID, error) {

}
