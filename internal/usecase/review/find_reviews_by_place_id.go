package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
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

func (uc *FincdReviewsByPlaceIDUseCase) Execute(ctx context.Context, input *dto.FindReviewsByPlaceIDInputDTO) ([]*dto.FindReviewsByPlaceIDOutputDTO, error) {
	reviews, err := uc.ReviewsGateway.FindReviewsByPlaceId(ctx, input.PlaceID)
	if err != nil {
		return nil, errors.New("review not found: " + err.Error())
	}
	var output []*dto.FindReviewsByPlaceIDOutputDTO
	var actualReviews []*entity.Review
	if input.Offset > len(reviews) {
		return output, nil
	} else if input.Offset+input.Limit > len(reviews) {
		actualReviews = reviews[input.Offset:]
	} else {
		actualReviews = reviews[input.Offset : input.Offset+input.Limit]
	}
	for _, review := range actualReviews {
		output = append(output, &dto.FindReviewsByPlaceIDOutputDTO{
			ID:                    review.ID.String(),
			PlaceID:               review.PlaceID,
			UserID:                review.UserID,
			Text:                  review.Text,
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
