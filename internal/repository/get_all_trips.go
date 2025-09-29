package repository

import (
	"context"
	"fmt"
	"go_api/internal/dto"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (r *tripRepository) GetAllTrips(ctx context.Context, param *dto.GetAllTripsRequest, userID uuid.UUID, offSet int) ([]*model.TripModel, error) {
	query := `SELECT id, title,description, start_date, end_date FROM trips WHERE user_id = $1
	LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, userID, param.Limit, offSet)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []*model.TripModel
	for rows.Next() {
		var trip model.TripModel
		if err := rows.Scan(&trip.ID, &trip.Title, &trip.Description, &trip.Start_date, &trip.End_date); err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if trips == nil {
		return []*model.TripModel{}, nil
	}
	fmt.Println(trips)
	return trips, nil
}
