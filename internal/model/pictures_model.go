package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	PictureModel struct {
		ID        uuid.UUID
		Trip_id   uuid.UUID
		Url       string
		CreatedAt time.Time
	}
)
