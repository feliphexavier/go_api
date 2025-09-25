package dto

import "github.com/google/uuid"

type (
	CreatePictureRequest struct {
		Url string `json:"url" validade:"required"`
	}
	CreatePictureReponse struct {
		ID uuid.UUID `json:"id"`
	}
)
