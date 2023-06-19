package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type DeleteReviewByIDUseCase struct {
	ReviewGateway gateway.ReviewsGateway
}

func NewDeleteReviewByIDUseCase(reviewGateway gateway.ReviewsGateway) *DeleteReviewByIDUseCase {
	return &DeleteReviewByIDUseCase{
		ReviewGateway: reviewGateway,
	}
}

func (u *DeleteReviewByIDUseCase) Execute(ctx context.Context, input dto.DeleteReviewByIDInputDTO) error {
	err := u.ReviewGateway.DeleteReviewByID(ctx, input.ID)
	if err != nil {
		return errors.New("error deleting review: " + err.Error())
	}
	return nil
}
