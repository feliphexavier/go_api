package repository

import (
	"context"
	"database/sql"
	"go_api/internal/model"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.UserModel, error) {
	query := `SELECT id,username,email, password, created_at, updated_at  FROM users WHERE email = ? OR username = ?`
	row := r.db.QueryRowContext(ctx, query, email, username)
	var result model.UserModel
	err := row.Scan(&result.ID, &result.Username, &result.Email, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &result, nil
}
