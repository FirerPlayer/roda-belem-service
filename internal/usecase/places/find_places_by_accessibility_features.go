package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindPlacesByAccessibilityFeaturesUseCase struct {
	PlacesGateway gateway.PlacesGateway
}

func NewFindPlacesByAccessibilityFeaturesUseCase(placesGateway gateway.PlacesGateway) *FindPlacesByAccessibilityFeaturesUseCase {
	return &FindPlacesByAccessibilityFeaturesUseCase{
		PlacesGateway: placesGateway,
	}
}

var strToEnumMap = map[string]entity.AccessibilityFeaturesEnum{
	"Ramp":                      entity.Ramp,
	"Elevator":                  entity.Elevator,
	"AdaptedBathroom":           entity.AdaptedBathroom,
	"BrailleSignage":            entity.BrailleSignage,
	"WideCirculationAreas":      entity.WideCirculationAreas,
	"ReservedParking":           entity.ReservedParking,
	"TactilePaving":             entity.TactilePaving,
	"AdaptedTelephones":         entity.AdaptedTelephones,
	"VideoIntercom":             entity.VideoIntercom,
	"AdaptedTablesCounters":     entity.AdaptedTablesCounters,
	"WheelchairAvailability":    entity.WheelchairAvailability,
	"SignLanguageCommunication": entity.SignLanguageCommunication,
	"GuideDogAllowed":           entity.GuideDogAllowed,
	"OnlineAccessibility":       entity.OnlineAccessibility,
	"AssistiveTechnologyAccess": entity.AssistiveTechnologyAccess,
}

func (uc FindPlacesByAccessibilityFeaturesUseCase) Execute(ctx context.Context, input dto.FindPlacesByAccessibilityFeaturesInputDTO) ([]*dto.FindPlacesByAccessibilityFeaturesOutputDTO, error) {
	var af []entity.AccessibilityFeaturesEnum
	for _, v := range input.AccessibilityFeatures {
		af = append(af, strToEnumMap[v])
	}

	places, err := uc.PlacesGateway.FindPlacesByAccessibilityFeatures(ctx, af)
	if err != nil {
		return nil, errors.New("failed to find places by accessibility features: " + err.Error())
	}
	var output []*dto.FindPlacesByAccessibilityFeaturesOutputDTO
	for _, place := range places {
		output = append(output, &dto.FindPlacesByAccessibilityFeaturesOutputDTO{
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
