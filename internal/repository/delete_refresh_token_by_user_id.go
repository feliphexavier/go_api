package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (r *userRepository) DeleteRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (int64, error) {
	query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	result, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAffected == 0 {
		return 0, errors.New("nothing to delete")
	}
	return rowsAffected, nil
}
