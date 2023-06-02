package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
)

type FindReviewsByUserIDUseCase struct {
	ReviewsGateway gateway.ReviewsGateway
}

func NewFindReviewsByUserIDUseCase(reviewsGateway gateway.ReviewsGateway) *FindReviewsByUserIDUseCase {
	return &FindReviewsByUserIDUseCase{
		ReviewsGateway: reviewsGateway,
	}
}

func (uc *FindReviewsByUserIDUseCase) Execute(ctx context.Context, input *dto.FindReviewsByUserIDInputDTO) ([]*dto.FindReviewsByUserIDOutputDTO, error) {
	reviews, err := uc.ReviewsGateway.FindReviewsByUserId(ctx, input.UserID)
	if err != nil {
		return nil, errors.New("review not found: " + err.Error())
	}
	var output []*dto.FindReviewsByUserIDOutputDTO
	for _, review := range reviews {
		output = append(output, &dto.FindReviewsByUserIDOutputDTO{
			ID:        review.ID.String(),
			PlaceID:   review.PlaceID,
			UserID:    review.UserID,
			Text:      review.Text,
			Photos:    review.Photos,
			Rating:    review.Rating,
			Reactions: review.Reactions,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
	}
	return output, nil

}
