package dto

import (
	"github.com/google/uuid"
)

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
	Picture struct {
		ID  uuid.UUID `json:"id"`
		Url string    `json:"url"`
	}
	GetTripResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Start_Date  string    `json:"start_date"`
		End_Date    string    `json:"end_date"`
		CreateAt    string    `json:"created_at"`
		Pictures    []Picture `json:"pictures"`
	}
)
