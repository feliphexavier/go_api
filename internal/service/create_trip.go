package service

import (
	"context"
	"go_api/internal/dto"
	"go_api/internal/model"
	"net/http"

	"github.com/google/uuid"
)

func (s *tripService) CreateTrip(ctx context.Context, req *dto.CreateOrUpdateTripRequest, userID uuid.UUID) (uuid.UUID, int, error) {
	id := uuid.New()
	_, err := s.tripRepo.CreateTrip(ctx, &model.TripModel{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Start_date:  req.Start_Date,
		End_date:    req.End_Date,
		User_id:     userID,
	}, uuid.Nil)
	if err != nil {
		return uuid.Nil, http.StatusInternalServerError, err
	}
	return id, http.StatusOK, nil
}
