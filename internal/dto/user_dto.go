package dto

import "github.com/google/uuid"

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RegisterResponse struct {
		ID uuid.UUID
	}
)
