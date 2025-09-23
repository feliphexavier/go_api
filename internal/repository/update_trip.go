package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (r *tripRepository) UpdateTrip(ctx context.Context, model *model.TripModel, tripID uuid.UUID) error {
	query := `
		UPDATE trips
		SET title=$1, description=$2, start_date=$3, end_date=$4
		WHERE id=$5
	`
	result, err := r.db.ExecContext(ctx, query,
		model.Title,
		model.Description,
		model.Start_date,
		model.End_date,
		tripID,
	)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
