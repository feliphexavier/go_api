package repository

import (
	"context"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (r *tripRepository) GetTripByID(ctx context.Context, tripID uuid.UUID) (*model.TripModel, error) {
	query := `SELECT id, title, description, start_date, end_date, user_id FROM trips WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, tripID)
	var result model.TripModel
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Start_date,
		&result.End_date,
		&result.User_id,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
