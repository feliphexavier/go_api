package repository

import (
	"context"
	"database/sql"
	"go_api/internal/model"

	"github.com/google/uuid"
)

type TripRepository interface {
	CreateTrip(ctx context.Context, model *model.TripModel, userID uuid.UUID) (uuid.UUID, error)
}
type tripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) TripRepository {
	return &tripRepository{
		db: db,
	}
}
