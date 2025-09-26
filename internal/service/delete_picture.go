package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *pictureService) DeletePicture(ctx context.Context, pictureID uuid.UUID) error {
	return s.pictureRepo.DeletePicture(ctx, pictureID)
}
