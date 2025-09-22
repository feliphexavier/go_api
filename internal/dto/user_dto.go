package dto

import "github.com/google/uuid"

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=10"`
	}
	RegisterResponse struct {
		ID uuid.UUID
	}
)

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)

type (
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
	RefreshTokenResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
