package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type AddAccessibilityFeaturesByReviewIDUseCase struct {
	ReviewGateway gateway.ReviewsGateway
}

func NewAddAccessibilityFeaturesByReviewIdUseCase(reviewGateway gateway.ReviewsGateway) *AddAccessibilityFeaturesByReviewIDUseCase {
	return &AddAccessibilityFeaturesByReviewIDUseCase{
		ReviewGateway: reviewGateway,
	}
}

func (u *AddAccessibilityFeaturesByReviewIDUseCase) Execute(ctx context.Context, input dto.AddAccessibilityFeaturesByReviewIDInputDTO) error {
	formatedFeatures := strings.Join(input.Features, ",")
	err := u.ReviewGateway.AddAccessibilityFeatureByReviewID(ctx, input.ReviewID, formatedFeatures)
	if err != nil {
		return errors.New("error adding accessibility features: " + err.Error())
	}
	return nil
}
