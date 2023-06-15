package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindReviewByIdUseCase struct {
	ReviewGateway gateway.ReviewsGateway
}

func NewFindReviewByIdUseCase(reviewGateway gateway.ReviewsGateway) *FindReviewByIdUseCase {
	return &FindReviewByIdUseCase{
		ReviewGateway: reviewGateway,
	}
}

func (u *FindReviewByIdUseCase) Execute(ctx context.Context, input dto.FindReviewByIDInputDTO) (*dto.FindReviewByIDOutputDTO, error) {
	review, err := u.ReviewGateway.FindReviewById(ctx, input.Id)
	if err != nil {
		return nil, errors.New("review not found: " + err.Error())
	}
	return &dto.FindReviewByIDOutputDTO{
		ID:        review.ID.String(),
		PlaceID:   review.PlaceID,
		UserID:    review.UserID,
		Text:      review.Text,
		Photos:    review.Photos,
		Rating:    review.Rating,
		Reactions: review.Reactions,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}, nil
}
