package repository

import (
	"context"
	"database/sql"
	"go_api/internal/model"

	"github.com/google/uuid"
)

type TripRepository interface {
	CreateTrip(ctx context.Context, model *model.TripModel, userID uuid.UUID) (uuid.UUID, error)
	GetTripByID(ctx context.Context, tripID uuid.UUID) (*model.TripModel, error)
	UpdateTrip(ctx context.Context, model *model.TripModel, tripID uuid.UUID) error
	DeleteTrip(ctx context.Context, tripID uuid.UUID) (int, error)
}
type tripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) TripRepository {
	return &tripRepository{
		db: db,
	}
}
