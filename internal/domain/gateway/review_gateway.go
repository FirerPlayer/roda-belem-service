package gateway

import (
	"context"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
)

type ReviewsGateway interface {
	Create(ctx context.Context, review *entity.Review) error
	FindReviewByID(ctx context.Context, id string) (*entity.Review, error)
	FindReviewsByPlaceID(ctx context.Context, placeId string, limit int, offset int) ([]*entity.Review, error)
	FindReviewsByUserID(ctx context.Context, userId string, limit int, offset int) ([]*entity.Review, error)
	UpdateReviewByID(ctx context.Context, id string, review *entity.Review) error
	DeleteReviewByID(ctx context.Context, id string) error
	AddAccessibilityFeatureByReviewID(ctx context.Context, reviewId string, feature string) error
}
