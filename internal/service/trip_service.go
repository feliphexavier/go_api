package service

import (
	"context"
	"go_api/internal/config"
	"go_api/internal/dto"
	"go_api/internal/repository"

	"github.com/google/uuid"
)

type TripService interface {
	CreateTrip(ctx context.Context, req *dto.CreateOrUpdateTripRequest, userID uuid.UUID) (uuid.UUID, int, error)
	UpdateTrip(ctx context.Context, req *dto.CreateOrUpdateTripRequest, tripID, userID uuid.UUID) (uuid.UUID, int, error)
	DeleteTrip(ctx context.Context, tripID, userID uuid.UUID) (int, error)
	GetTripByID(ctx context.Context, tripID, userID uuid.UUID) (*dto.GetTripResponse, int, error)
	GetAllTrip(ctx context.Context, param *dto.GetAllTripsRequest, userID uuid.UUID) (*dto.GetAllTripsResponse, int, error)
}
type tripService struct {
	cfg         *config.Config
	tripRepo    repository.TripRepository
	pictureRepo repository.PictureRepository
}

func NewTripService(cfg *config.Config, tripRepo repository.TripRepository, pictureRepo repository.PictureRepository) TripService {
	return &tripService{
		cfg:         cfg,
		tripRepo:    tripRepo,
		pictureRepo: pictureRepo,
	}
}
