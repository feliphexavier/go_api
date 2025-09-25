package service

import (
	"context"
	"go_api/internal/dto"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (s *pictureService) CreatePicture(ctx context.Context, req *dto.CreatePictureRequest, tripID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()
	tripExists, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return uuid.Nil, err
	}
	if tripExists == nil {
		return uuid.Nil, nil
	}

	_, err = s.pictureRepo.CreatePicture(ctx, &model.PictureModel{
		ID:      id,
		Url:     req.Url,
		Trip_id: tripID,
	})
	if err != nil {
		return uuid.Nil, err
	}
	return tripID, nil
}
