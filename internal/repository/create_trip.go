package repository

import (
	"context"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (r *tripRepository) CreateTrip(ctx context.Context, model *model.TripModel, userID uuid.UUID) (uuid.UUID, error) {
	query := `INSERT INTO trips (id, title,description, start_date, end_date, user_id) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, model.ID, model.Title, model.Description, model.Start_date, model.End_date, model.User_id)
	if err != nil {
		return uuid.Nil, err
	}
	return model.ID, nil
}
