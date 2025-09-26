package service

import (
	"context"
	"fmt"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (s *pictureService) CreatePicture(ctx context.Context, filesPath []string, tripID uuid.UUID) ([]string, error) {
	tripExists, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return nil, err
	}
	if tripExists == nil {
		return nil, fmt.Errorf("trip not found")
	}

	var picturesPath []string

	for _, filePath := range filesPath {
		id := uuid.New()

		_, err := s.pictureRepo.CreatePicture(ctx, &model.PictureModel{
			ID:      id,
			Url:     filePath,
			Trip_id: tripID,
		})
		if err != nil {
			return nil, err
		}

		picturesPath = append(picturesPath, filePath)
	}

	return picturesPath, nil
}
