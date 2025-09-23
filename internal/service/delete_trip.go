package service

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *tripService) DeleteTrip(ctx context.Context, tripID, userID uuid.UUID) (int, error) {
	tripExist, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return http.StatusNotFound, err
	}
	if tripExist == nil {
		return http.StatusNotFound, err
	}
	if tripExist.User_id != userID {
		return http.StatusNotFound, err
	}
	err = s.tripRepo.DeleteTrip(ctx, tripID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
