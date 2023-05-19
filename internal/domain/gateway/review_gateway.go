package gateway

import (
	"context"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
)

type ReviewsGateway interface {
	Create(ctx context.Context, review *entity.Review) error
	FindReviewById(ctx context.Context, id string) (*entity.Review, error)
	FindReviewsByPlaceId(ctx context.Context, placeId string) ([]*entity.Review, error)
	FindReviewsByUserId(ctx context.Context, userId string) ([]*entity.Review, error)
}
