package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func (r *tripRepository) DeleteTrip(ctx context.Context, tripID uuid.UUID) (int, error) {
	query := `DELETE FROM trips WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, tripID)
	if err != nil {
		return 0, errors.New("query failed")
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return 0, errors.New("Nothing to delete")
	}
	return http.StatusNoContent, nil
}
