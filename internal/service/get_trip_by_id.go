package service

import (
	"context"
	"go_api/internal/dto"
	"net/http"

	"github.com/google/uuid"
)

func (s *tripService) GetTripByID(ctx context.Context, tripID, userID uuid.UUID) (*dto.GetTripResponse, int, error) {
	trip, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if trip == nil {
		return nil, http.StatusNotFound, nil
	}
	if trip.User_id != userID {
		return nil, http.StatusForbidden, nil
	}
	tripIDs := []uuid.UUID{trip.ID}
	pictures, err := s.pictureRepo.GetPicturesByTripID(ctx, tripIDs)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	pictureDTOs := make([]dto.Picture, 0, len(pictures))
	for _, pic := range pictures {
		pictureDTOs = append(pictureDTOs, dto.Picture{
			ID:  pic.ID,
			Url: pic.Url,
		})
	}
	return &dto.GetTripResponse{
		ID:          trip.ID,
		Title:       trip.Title,
		Description: trip.Description,
		Start_Date:  trip.Start_date,
		End_Date:    trip.End_date,
		CreateAt:    trip.CreatedAt.String(),
		Pictures:    pictureDTOs,
	}, http.StatusOK, nil
}
