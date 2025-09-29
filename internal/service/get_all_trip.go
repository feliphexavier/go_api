package service

import (
	"context"
	"go_api/internal/dto"
	"net/http"

	"github.com/google/uuid"
)

func (s *tripService) GetAllTrip(ctx context.Context, param *dto.GetAllTripsRequest, userID uuid.UUID) (*dto.GetAllTripsResponse, int, error) {
	trips, err := s.tripRepo.GetAllTrips(ctx, param, userID, (param.Page-1)*param.Limit)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	tripIDs := make([]uuid.UUID, 0, len(trips))
	for _, trip := range trips {
		tripIDs = append(tripIDs, trip.ID)
	}
	pictures, err := s.pictureRepo.GetPicturesByTripID(ctx, tripIDs)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	pictureMap := make(map[uuid.UUID][]dto.Picture)
	for _, pic := range pictures {
		pictureMap[pic.Trip_id] = append(pictureMap[pic.Trip_id], dto.Picture{
			ID:  pic.ID,
			Url: pic.Url,
		})
	}
	var tripDTOs []dto.GetTripResponse
	for _, trip := range trips {
		picture := pictureMap[trip.ID]
		if picture == nil {
			picture = []dto.Picture{}
		}
		tripDTOs = append(tripDTOs, dto.GetTripResponse{
			ID:          trip.ID,
			Title:       trip.Title,
			Description: trip.Description,
			Start_Date:  trip.Start_date,
			End_Date:    trip.End_date,
			CreateAt:    trip.CreatedAt.String(),
			Pictures:    picture,
		})

	}
	totalTrips := len(tripDTOs)
	totalPages := (totalTrips + param.Limit - 1) / param.Limit
	return &dto.GetAllTripsResponse{
		TotalPages:  totalPages,
		CurrentPage: param.Page,
		Limit:       param.Limit,
		Data:        tripDTOs,
	}, http.StatusOK, nil
}
