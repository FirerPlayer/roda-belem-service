package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/infra/blooms"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	"googlemaps.github.io/maps"
)

type FindNearbyPlacesUseCase struct {
	PlacesGateway    gateway.PlacesGateway
	ReviewsGateway   gateway.ReviewsGateway
	GoogleMapsClient *maps.Client
	BloomFilter      *blooms.BloomFilter
}

func NewFindNearbyPlacesUseCase(placesGateway gateway.PlacesGateway, reviewsGateway gateway.ReviewsGateway, googleMapsClient *maps.Client, bloomFilter *blooms.BloomFilter) *FindNearbyPlacesUseCase {
	return &FindNearbyPlacesUseCase{
		PlacesGateway:    placesGateway,
		ReviewsGateway:   reviewsGateway,
		GoogleMapsClient: googleMapsClient,
		BloomFilter:      bloomFilter,
	}
}

func (u *FindNearbyPlacesUseCase) GetNearbyPlacesFromGoogleMapsWithPersistence(ctx context.Context, input dto.FindNearbyPlacesInputDTO) ([]*entity.Place, error) {
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
		if u.BloomFilter.NotContains(place.PlaceID) {
			newPLace := entity.NewPlace(
				place.PlaceID,
				place.Name,
				place.FormattedAddress,
				place.Geometry.Location.Lat,
				place.Geometry.Location.Lng,
				place.Icon,
				place.Types,
				place.OpeningHours.WeekdayText,
			)
			err := u.PlacesGateway.Create(ctx, newPLace)
			if err != nil {
				return nil, errors.New("Failed to persist place while google maps nearby search: " + err.Error())
			}
			u.BloomFilter.Add(place.PlaceID)
			output = append(output, newPLace)
		}
	}

	return output, nil
}

func (u *FindNearbyPlacesUseCase) GetNearbyPlacesFromGoogleOrRepository(ctx context.Context, input dto.FindNearbyPlacesInputDTO) ([]*entity.Place, error) {
	if input.IsFromGoogle {
		places, err := u.GetNearbyPlacesFromGoogleMapsWithPersistence(ctx, input)
		if err != nil {
			return nil, errors.New("Failed to find nearby places from Google: " + err.Error())
		}
		return places, nil
	}
	places, err := u.PlacesGateway.FindNearbyPlaces(ctx, input.Lat, input.Lng, input.Radius)
	if err != nil {
		return nil, errors.New("Failed to find nearby places " + err.Error())
	}
	return places, nil
}

// Execute finds nearby places and returns an array of related output DTO objects or an error.
//
// ctx - context object.
// input - input DTO object for finding nearby places.
//
// []*dto.FindNearbyPlacesOutputDTO - an array of output DTO objects containing information about nearby places.
// error - an error object if the function fails to find nearby places.
func (u *FindNearbyPlacesUseCase) Execute(ctx context.Context, input dto.FindNearbyPlacesInputDTO) ([]*dto.FindNearbyPlacesOutputDTO, error) {
	places, err := u.GetNearbyPlacesFromGoogleOrRepository(ctx, input)
	if err != nil {
		return nil, err
	}

	var output []*dto.FindNearbyPlacesOutputDTO
	for _, place := range places {
		output = append(output, &dto.FindNearbyPlacesOutputDTO{
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
