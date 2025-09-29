package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go_api/internal/model"
	"strings"

	"github.com/google/uuid"
)

func (r *pictureRepository) GetPicturesByTripID(ctx context.Context, tripID []uuid.UUID) ([]model.PictureModel, error) {
	if len(tripID) == 0 {
		return []model.PictureModel{}, nil
	}
	placeholders := make([]string, len(tripID))
	args := make([]interface{}, len(tripID))
	for i, id := range tripID {
		placeholders[i] = "$" + fmt.Sprint(i+1)
		args[i] = id
	}
	query := fmt.Sprintf(`SELECT p.id, p.trip_id, p.url, p.created_at FROM pictures p
	JOIN trips t ON p.trip_id = t.id
	WHERE p.trip_id IN (%s) GROUP BY p.id, p.trip_id, p.url, p.created_at`, strings.Join(placeholders, ","))
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return []model.PictureModel{}, nil
		}
		return []model.PictureModel{}, err
	}
	result := make([]model.PictureModel, 0)
	for rows.Next() {
		var picture model.PictureModel
		err = rows.Scan(&picture.ID, &picture.Trip_id, &picture.Url, &picture.CreatedAt)
		if err != nil {
			return []model.PictureModel{}, err
		}
		result = append(result, model.PictureModel{
			ID:        picture.ID,
			Trip_id:   picture.Trip_id,
			Url:       picture.Url,
			CreatedAt: picture.CreatedAt,
		})
	}
	return result, nil
}
