package service

import (
	"context"
	"errors"
	"go_api/internal/dto"
	"go_api/internal/model"
	"net/http"

	"github.com/google/uuid"
)

func (s *tripService) UpdateTrip(ctx context.Context, req *dto.CreateOrUpdateTripRequest, tripID, userID uuid.UUID) (uuid.UUID, int, error) {
	tripExists, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return uuid.Nil, http.StatusContinue, errors.New("")
	}
	if tripExists == nil {
		return uuid.Nil, http.StatusNotFound, errors.New("trip not found")
	}
	if tripExists.User_id != userID {
		return uuid.Nil, http.StatusNotFound, errors.New("trip not found")
	}
	err = s.tripRepo.UpdateTrip(ctx, &model.TripModel{
		Title:       req.Title,
		Description: req.Description,
		Start_date:  req.Start_Date,
		End_date:    req.End_Date,
	}, tripID)
	if err != nil {
		return uuid.Nil, http.StatusInternalServerError, err
	}
	return tripExists.ID, http.StatusOK, nil
}
