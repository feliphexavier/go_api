package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (r *userRepository) GetUserById(ctx context.Context, userId uuid.UUID) (*model.UserModel, error) {
	query := `SELECT id, email, username, hashed_pass from users 
	WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, userId)
	var result model.UserModel
	err := row.Scan(&result.ID, &result.Email, &result.Username, &result.Hashed_password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
