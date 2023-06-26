package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FincdReviewsByPlaceIDUseCase struct {
	ReviewsGateway gateway.ReviewsGateway
}

func NewFindReviewsByPlaceIDUseCase(reviewsGateway gateway.ReviewsGateway) *FincdReviewsByPlaceIDUseCase {
	return &FincdReviewsByPlaceIDUseCase{
		ReviewsGateway: reviewsGateway,
	}
}

func (uc *FincdReviewsByPlaceIDUseCase) Execute(ctx context.Context, input dto.FindReviewsByPlaceIDInputDTO) ([]*dto.FindReviewsByPlaceIDOutputDTO, error) {
	reviews, err := uc.ReviewsGateway.FindReviewsByPlaceID(ctx, input.PlaceID, input.Limit, input.Offset)
	if err != nil {
		return nil, errors.New("review not found -> " + err.Error())
	}
	var output []*dto.FindReviewsByPlaceIDOutputDTO
	for _, review := range reviews {
		output = append(output, &dto.FindReviewsByPlaceIDOutputDTO{
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
