package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindReviewsByUserIDUseCase struct {
	ReviewsGateway gateway.ReviewsGateway
}

func NewFindReviewsByUserIDUseCase(reviewsGateway gateway.ReviewsGateway) *FindReviewsByUserIDUseCase {
	return &FindReviewsByUserIDUseCase{
		ReviewsGateway: reviewsGateway,
	}
}

func (uc *FindReviewsByUserIDUseCase) Execute(ctx context.Context, input dto.FindReviewsByUserIDInputDTO) ([]*dto.FindReviewsByUserIDOutputDTO, error) {
	reviews, err := uc.ReviewsGateway.FindReviewsByPlaceID(ctx, input.UserID, input.Limit, input.Offset)
	if err != nil {
		return nil, errors.New("review not found -> " + err.Error())
	}
	var output []*dto.FindReviewsByUserIDOutputDTO
	for _, review := range reviews {
		output = append(output, &dto.FindReviewsByUserIDOutputDTO{
			ID:                    review.ID.String(),
			PlaceID:               review.PlaceID,
			UserID:                review.UserID,
			Content:               review.Content,
			Photos:                review.Photos,
			Rating:                review.Rating,
			Reactions:             review.Reactions,
			AccessibilityFeatures: review.AccessibilityFeatures,
			CreatedAt:             review.CreatedAt,
			UpdatedAt:             review.UpdatedAt,
		})
	}
	return output, nil

}
