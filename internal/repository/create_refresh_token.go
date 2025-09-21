package repository

import (
	"context"
	"go_api/internal/model"
)

func (r *userRepository) CreateRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error {

	query := `INSERT INTO refresh_tokens (id , user_id, refresh_token, expired_at, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := r.db.Exec(query, model.ID, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt)
	return err
}
