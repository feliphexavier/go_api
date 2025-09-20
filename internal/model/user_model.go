package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserModel struct {
		ID              uuid.UUID
		Email           string
		Username        string
		Hashed_password string
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}
)
