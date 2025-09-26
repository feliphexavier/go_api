package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (r *pictureRepository) DeletePicture(ctx context.Context, pictureID uuid.UUID) error {
	query := `DELETE FROM pictures WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, pictureID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Nothing to delete")
	}
	return nil
}
