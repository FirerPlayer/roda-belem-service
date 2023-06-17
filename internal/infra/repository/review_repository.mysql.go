package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	db "github.com/firerplayer/roda-belem-service/internal/infra/mysql"
)

// type ReviewsGateway interface {
// 	Create(ctx context.Context, review *entity.Review) error
// 	FindReviewById(ctx context.Context, id string) (*entity.Review, error)
// 	FindReviewsByPlaceId(ctx context.Context, placeId string) ([]*entity.Review, error)
// 	FindReviewsByUserId(ctx context.Context, userId string) ([]*entity.Review, error)
// }

type ReviewRepositoryMySQL struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewReviewRepositoryMySQL(dbt *sql.DB) *ReviewRepositoryMySQL {
	return &ReviewRepositoryMySQL{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *ReviewRepositoryMySQL) Create(ctx context.Context, review *entity.Review) error {
	params := db.CreateReviewParams{
		ID:        review.ID.String(),
		PlaceID:   sql.NullString{String: review.PlaceID},
		UserID:    sql.NullString{String: review.UserID},
		Text:      sql.NullString{String: review.Text},
		Rating:    sql.NullFloat64{Float64: review.Rating},
		CreatedAt: sql.NullTime{Time: review.CreatedAt},
		UpdatedAt: sql.NullTime{Time: review.UpdatedAt},
	}
	photos, err := json.Marshal(review.Photos)
	if err != nil {
		return err
	}
	params.Images = photos
	reactions, err := json.Marshal(review.Reactions)
	if err != nil {
		return err
	}
	params.Reactions = reactions
	return r.Queries.CreateReview(ctx, params)
}

func (r *ReviewRepositoryMySQL) FindReviewById(ctx context.Context, id string) (*entity.Review, error) {
	reviewDb, err := r.Queries.FindReviewById(ctx, id)
	if err != nil {
		return nil, err
	}
	var review entity.Review
	err = json.Unmarshal(reviewDb.Images, &review.Photos)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reviewDb.Reactions, &review.Reactions)
	if err != nil {
		return nil, err
	}

}
