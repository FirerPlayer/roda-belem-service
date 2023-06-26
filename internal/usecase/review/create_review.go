package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type CreateReviewUsecase struct {
	ReviewsGateway gateway.ReviewsGateway
}

func NewCreateReviewUseCase(reviewsGateway gateway.ReviewsGateway) *CreateReviewUsecase {
	return &CreateReviewUsecase{
		ReviewsGateway: reviewsGateway,
	}
}

func (uc *CreateReviewUsecase) Execute(ctx context.Context, input dto.CreateReviewInputDTO) error {
	newReview := entity.NewReview(
		input.PlaceId,
		input.UserID,
		input.Content,
		input.Photos,
		input.Rating,
		input.AccessibilityFeatures,
	)

	err := uc.ReviewsGateway.Create(ctx, newReview)
	if err != nil {
		return errors.New("failed to create review -> " + err.Error())
	}
	return nil
}
