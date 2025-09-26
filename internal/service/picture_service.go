package service

import (
	"context"
	"go_api/internal/config"
	"go_api/internal/dto"
	"go_api/internal/repository"

	"github.com/google/uuid"
)

type PictureService interface {
	CreatePicture(ctx context.Context, req *dto.CreatePictureRequest, tripID uuid.UUID) (uuid.UUID, error)
	DeletePicture(ctx context.Context, pictureID uuid.UUID) error
}
type pictureService struct {
	cfg         *config.Config
	pictureRepo repository.PictureRepository
	tripRepo    repository.TripRepository
}

func NewPictureService(cfg *config.Config, pictureRepo repository.PictureRepository, tripRepo repository.TripRepository) PictureService {
	return &pictureService{
		cfg:         cfg,
		pictureRepo: pictureRepo,
		tripRepo:    tripRepo,
	}
}
