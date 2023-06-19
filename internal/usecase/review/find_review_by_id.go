package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindReviewByIDUseCase struct {
	ReviewGateway gateway.ReviewsGateway
}

func NewFindReviewByIdUseCase(reviewGateway gateway.ReviewsGateway) *FindReviewByIDUseCase {
	return &FindReviewByIDUseCase{
		ReviewGateway: reviewGateway,
	}
}

func (u *FindReviewByIDUseCase) Execute(ctx context.Context, input dto.FindReviewByIDInputDTO) (*dto.FindReviewByIDOutputDTO, error) {
	review, err := u.ReviewGateway.FindReviewByID(ctx, input.Id)
	if err != nil {
		return nil, errors.New("review not found: " + err.Error())
	}
	return &dto.FindReviewByIDOutputDTO{
		ID:                    review.ID.String(),
		PlaceID:               review.PlaceID,
		UserID:                review.UserID,
		Text:                  review.Content,
		Photos:                review.Photos,
		Rating:                review.Rating,
		Reactions:             review.Reactions,
		AccessibilityFeatures: review.AccessibilityFeatures,
		CreatedAt:             review.CreatedAt,
		UpdatedAt:             review.UpdatedAt,
	}, nil
}
