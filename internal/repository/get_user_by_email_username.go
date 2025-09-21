package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_api/internal/model"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.UserModel, error) {
	query := `SELECT id, email, username, hashed_pass, created_at, updated_at 
              FROM users 
              WHERE email = $1 
			  OR username = $2`

	row := r.db.QueryRowContext(ctx, query, email, username)

	var result model.UserModel
	err := row.Scan(
		&result.ID,
		&result.Email,
		&result.Username,
		&result.Hashed_password,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
