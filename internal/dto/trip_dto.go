package dto

import "github.com/google/uuid"

type (
	CreateOrUpdateTripRequest struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Start_Date  string `json:"start_date" validate:"required,datetime=2006-01-02"`
		End_Date    string `json:"end_date" validate:"required,datetime=2006-01-02"`
	}
	CreateOrUpdateTripResponse struct {
		ID uuid.UUID `json:"id"`
	}
)
