package repository

import (
	"context"
	"go_api/internal/model"

	"github.com/google/uuid"
)

func (r *pictureRepository) CreatePicture(ctx context.Context, model *model.PictureModel) (uuid.UUID, error) {
	query := `INSERT INTO pictures (id, url, trip_id) VALUES ($1,$2,$3)`
	_, err := r.db.ExecContext(ctx, query, model.ID, model.Url, model.Trip_id)
	if err != nil {
		return uuid.Nil, err
	}
	return model.ID, nil
}
