package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	TripModel struct {
		ID          uuid.UUID
		User_id     uuid.UUID
		Title       string
		Description string
		Start_date  string
		End_date    string
		CreatedAt   time.Time
	}
)
