package repository

import (
	"context"
	"database/sql"
	"go_api/internal/model"

	"github.com/google/uuid"
)

type PictureRepository interface {
	CreatePicture(ctx context.Context, model *model.PictureModel) (uuid.UUID, error)
	DeletePicture(ctx context.Context, pictureID uuid.UUID) error
}
type pictureRepository struct {
	db *sql.DB
}

func NewPictureRepository(db *sql.DB) PictureRepository {
	return &pictureRepository{
		db: db,
	}
}
