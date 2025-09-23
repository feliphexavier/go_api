package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (r *tripRepository) DeleteTrip(ctx context.Context, tripID uuid.UUID) error {
	query := `DELETE FROM trips WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, tripID)
	if err != nil {
		return errors.New("query failed")
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("nothing to delete")
	}
	return nil
}
