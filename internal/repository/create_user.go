package repository

import (
	"context"
	"go_api/internal/model"
	"net/http"

	"github.com/google/uuid"
)

func (r *userRepository) CreateUser(ctx context.Context, model *model.UserModel) (uuid.UUID, int, error) {
	query := `INSERT INTO users (id,email, username, hashed_pass, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := r.db.ExecContext(ctx, query, model.ID, model.Email, model.Username, model.Hashed_password, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return uuid.Nil, http.StatusInternalServerError, err
	}

	return model.ID, http.StatusCreated, nil
}
