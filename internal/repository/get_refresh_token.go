package repository

import (
	"context"
	"database/sql"
	"go_api/internal/model"
	"time"

	"github.com/google/uuid"
)

func (r *userRepository) GetRefreshToken(ctx context.Context, userID uuid.UUID, now time.Time) (*model.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expired_at FROM refresh_tokens WHERE user_id = $1 AND expired_at >= $2`
	row := r.db.QueryRow(query, userID, now)
	var result model.RefreshTokenModel
	err := row.Scan(&result.ID, &result.UserID, &result.RefreshToken, &result.ExpiredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &result, nil
}
