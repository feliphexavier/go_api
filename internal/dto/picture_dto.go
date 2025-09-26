package dto

import "github.com/google/uuid"

type (
	CreatePictureRequest struct {
		Url string `json:"url" validade:"required"`
	}
	CreatePictureReponse struct {
		IDS []uuid.UUID `json:"id"`
	}
)
