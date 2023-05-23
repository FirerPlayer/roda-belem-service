package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/infra/blooms"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
	"github.com/google/uuid"
	"googlemaps.github.io/maps"
)

type FindNearbyPlacesUseCase struct {
	PlacesGateway    gateway.PlacesGateway
	GoogleMapsClient *maps.Client
	BloomFilter      *blooms.BloomFilter
}

func NewFindNearbyPlacesUseCase(placesGateway gateway.PlacesGateway, googleMapsClient *maps.Client, bloomFilter *blooms.BloomFilter) *FindNearbyPlacesUseCase {
	return &FindNearbyPlacesUseCase{
		PlacesGateway:    placesGateway,
		GoogleMapsClient: googleMapsClient,
		BloomFilter:      bloomFilter,
	}
}

func (u *FindNearbyPlacesUseCase) GetNearbyPlacesFromGoogleMaps(ctx context.Context, input *dto.FindNearbyPlacesInputDTO) ([]*entity.Place, error) {
	request := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: input.Lat,
			Lng: input.Lng,
		},
		Radius: uint(math.Ceil(input.Radius)),
	}
	response, err := u.GoogleMapsClient.NearbySearch(ctx, request)
	if err != nil {
		return nil, errors.New("Failed to find nearby places " + err.Error())
	}
	var output []*entity.Place
	for _, place := range response.Results {
		newPlace := &entity.Place{
			ID:              uuid.New(),
			PlaceId:         place.PlaceID,
			Name:            place.Name,
			FormatedAddress: place.FormattedAddress,
			Lat:             place.Geometry.Location.Lat,
			Lng:             place.Geometry.Location.Lng,
			Icon:            place.Icon,
			Types:           place.Types,
			OpeningPeriods:  place.OpeningHours.WeekdayText,
		}
		err := u.PlacesGateway.Create(ctx, newPlace)
		if err != nil {
			return nil, errors.New("Failed to create places from nearby places request: " + err.Error())
		}
		u.BloomFilter.Add(newPlace.ID.String())
		output = append(output, newPlace)
	}

	return output, nil
}

func (u *FindNearbyPlacesUseCase) Execute(ctx context.Context, input dto.FindNearbyPlacesInputDTO) ([]*dto.FindNearbyPlacesOutputDTO, error) {
	places, err := u.PlacesGateway.FindNearbyPlaces(ctx, input.Lat, input.Lng, input.Radius)
	if err != nil {
		return nil, errors.New("Failed to find nearby places " + err.Error())
	}
	var output []*dto.FindNearbyPlacesOutputDTO
	for _, place := range places {
		output = append(output, &dto.FindNearbyPlacesOutputDTO{
			ID:                    place.ID.String(),
			Name:                  place.GooglePlace.Name,
			FormatedAddress:       place.GooglePlace.FormattedAddress,
			Lat:                   place.GooglePlace.Geometry.Location.Lat,
			Lng:                   place.GooglePlace.Geometry.Location.Lng,
			Icon:                  place.GooglePlace.Icon,
			Types:                 place.GooglePlace.Types,
			OpeningPeriod:         place.GooglePlace.OpeningHours.WeekdayText,
			Photos:                place.Photos,
			Rating:                place.Rating,
			AccessibilityFeatures: place.AccessibilityFeatures,
		})
	}
	return output, nil
}
