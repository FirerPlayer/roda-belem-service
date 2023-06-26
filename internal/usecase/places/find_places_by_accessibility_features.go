package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindPlacesByAccessibilityFeatureUseCase struct {
	PlacesGateway gateway.PlacesGateway
}

func NewFindPlacesByAccessibilityFeatureUseCase(placesGateway gateway.PlacesGateway) *FindPlacesByAccessibilityFeatureUseCase {
	return &FindPlacesByAccessibilityFeatureUseCase{
		PlacesGateway: placesGateway,
	}
}

func (uc FindPlacesByAccessibilityFeatureUseCase) Execute(ctx context.Context, input dto.FindPlacesByAccessibilityFeatureInputDTO) ([]*dto.FindPlacesByAccessibilityFeatureOutputDTO, error) {

	places, err := uc.PlacesGateway.FindPlacesByAccessibilityFeature(ctx, input.AccessibilityFeature)
	if err != nil {
		return nil, errors.New("failed to find places by accessibility features -> " + err.Error())
	}
	var output []*dto.FindPlacesByAccessibilityFeatureOutputDTO
	for _, place := range places {
		output = append(output, &dto.FindPlacesByAccessibilityFeatureOutputDTO{
			ID:              place.ID.String(),
			Name:            place.Name,
			FormatedAddress: place.FormattedAddress,
			Lat:             place.Lat,
			Lng:             place.Lng,
			Icon:            place.Icon,
			Types:           place.Types,
			OpeningPeriods:  place.OpeningPeriods,
			Photos:          place.Photos,
			Rating:          place.Rating,
		})
	}
	return output, nil
}
