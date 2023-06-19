package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type UpdateReviewByIDUseCase struct {
	ReviewsGateway gateway.ReviewsGateway
}

func NewUpdateReviewByIDUseCase(reviewsGateway gateway.ReviewsGateway) *UpdateReviewByIDUseCase {
	return &UpdateReviewByIDUseCase{
		ReviewsGateway: reviewsGateway,
	}
}

func (u *UpdateReviewByIDUseCase) Execute(ctx context.Context, input dto.UpdateReviewByIDInputDTO) error {
	review := &entity.Review{
		Content:               input.Content,
		Photos:                input.Images,
		Rating:                input.Rating,
		Reactions:             input.Reactions,
		AccessibilityFeatures: input.AccessibilityFeatures,
	}
	err := u.ReviewsGateway.UpdateReviewByID(ctx, input.ID, review)
	if err != nil {
		return errors.New("failed to update review: " + err.Error())
	}
	return nil

}
