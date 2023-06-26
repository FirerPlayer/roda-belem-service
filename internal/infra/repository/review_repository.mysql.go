package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	db "github.com/firerplayer/roda-belem-service/internal/infra/mysql"
	"github.com/google/uuid"
)

//	type ReviewsGateway interface {
//		Create(ctx context.Context, review *entity.Review) error
//		FindReviewByID(ctx context.Context, id string) (*entity.Review, error)
//		FindReviewsByPlaceID(ctx context.Context, placeId string, limit int, offset int) ([]*entity.Review, error)
//		FindReviewsByUserID(ctx context.Context, userId string, limit int, offset int) ([]*entity.Review, error)
//		UpdateReviewByID(ctx context.Context, id string, review *entity.Review) error
//		DeleteReviewByID(ctx context.Context, id string) error
//		AddAccessibilityFeatureByReviewID(ctx context.Context, reviewId string, feature string) error
//	}
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
		PlaceID:   sql.NullString{String: review.PlaceID, Valid: true},
		UserID:    sql.NullString{String: review.UserID, Valid: true},
		Content:   sql.NullString{String: review.Content, Valid: true},
		Rating:    sql.NullFloat64{Float64: review.Rating, Valid: true},
		CreatedAt: sql.NullTime{Time: review.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: review.UpdatedAt, Valid: true},
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
	var afs string
	for _, af := range review.AccessibilityFeatures {
		afs = afs + string(af) + ","
	}
	params.AccessibilityFeatures = sql.NullString{String: afs}
	return r.Queries.CreateReview(ctx, params)
}

func (r *ReviewRepositoryMySQL) FindReviewByID(ctx context.Context, id string) (*entity.Review, error) {
	reviewDb, err := r.Queries.FindReviewById(ctx, id)
	if err != nil {
		return nil, err
	}
	var review entity.Review
	err = HydrateReview(reviewDb, &review)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepositoryMySQL) FindReviewsByPlaceID(ctx context.Context, placeId string, limit int, offset int) ([]*entity.Review, error) {
	reviewsDb, err := r.Queries.FindReviewsByPlaceId(ctx, db.FindReviewsByPlaceIdParams{
		PlaceID: sql.NullString{String: placeId, Valid: true},
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}
	var reviews []*entity.Review
	for _, reviewDb := range reviewsDb {
		var review entity.Review
		err = HydrateReview(reviewDb, &review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func (r *ReviewRepositoryMySQL) FindReviewsByUserID(ctx context.Context, userId string, limit int, offset int) ([]*entity.Review, error) {
	reviewsDb, err := r.Queries.FindReviewsByUserID(ctx, db.FindReviewsByUserIDParams{
		UserID: sql.NullString{String: userId, Valid: true},
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	var reviews []*entity.Review
	for _, reviewDb := range reviewsDb {
		var review entity.Review
		err = HydrateReview(reviewDb, &review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}
	return reviews, nil
}

func (r *ReviewRepositoryMySQL) UpdateReviewByID(ctx context.Context, id string, review *entity.Review) error {
	params := db.UpdateReviewByIdParams{
		Content:   sql.NullString{String: review.Content, Valid: true},
		Rating:    sql.NullFloat64{Float64: review.Rating, Valid: true},
		UpdatedAt: sql.NullTime{Time: review.UpdatedAt, Valid: true},
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
	afs := strings.Join(review.AccessibilityFeatures, ",")
	params.AccessibilityFeatures = sql.NullString{String: afs, Valid: true}
	return r.Queries.UpdateReviewById(ctx, params)
}

func (r *ReviewRepositoryMySQL) DeleteReviewByID(ctx context.Context, id string) error {
	return r.Queries.DeleteReviewById(ctx, id)
}

func (r *ReviewRepositoryMySQL) AddAccessibilityFeatureByReviewID(ctx context.Context, reviewId string, feature string) error {
	return r.Queries.AddAccessibilityFeatureByReviewID(ctx, db.AddAccessibilityFeatureByReviewIDParams{
		CONCAT: feature,
		ID:     reviewId,
	})
}

// HydrateReview hydrates a review entity with data from a review database object.
//
// reviewDb is the review object from the database.
// review is the review entity to be hydrated.
// Returns an error if unmarshalling of any field fails.
func HydrateReview(reviewDb db.Review, review *entity.Review) error {
	err := json.Unmarshal(reviewDb.Images, &review.Photos)
	if err != nil {
		return err
	}
	err = json.Unmarshal(reviewDb.Reactions, &review.Reactions)
	if err != nil {
		return err
	}
	id, err := uuid.Parse(reviewDb.ID)
	if err != nil {
		return err
	}
	review.ID = id
	review.CreatedAt = reviewDb.CreatedAt.Time
	review.UpdatedAt = reviewDb.UpdatedAt.Time
	review.UserID = reviewDb.UserID.String
	review.PlaceID = reviewDb.PlaceID.String
	review.Content = reviewDb.Content.String
	review.Rating = reviewDb.Rating.Float64
	review.AccessibilityFeatures = strings.Split(reviewDb.AccessibilityFeatures.String, ",")
	return nil
}
