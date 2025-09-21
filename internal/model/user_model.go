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
	RefreshTokenModel struct {
		ID           uuid.UUID
		UserID       uuid.UUID
		RefreshToken string
		ExpiredAt    time.Time
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)
