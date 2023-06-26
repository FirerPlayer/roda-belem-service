package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/infra/filters"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	"googlemaps.github.io/maps"
)

type FindNearbyPlacesUseCase struct {
	PlacesGateway    gateway.PlacesGateway
	GoogleMapsClient *maps.Client
	CuckooFilter     *filters.CuckooFilter
}

func NewFindNearbyPlacesUseCase(placesGateway gateway.PlacesGateway, googleMapsClient *maps.Client, cuckooFilter *filters.CuckooFilter) *FindNearbyPlacesUseCase {
	return &FindNearbyPlacesUseCase{
		PlacesGateway:    placesGateway,
		GoogleMapsClient: googleMapsClient,
		CuckooFilter:     cuckooFilter,
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
		return nil, errors.New("Failed to find nearby places -> " + err.Error())
	}
	var output []*entity.Place
	for _, place := range response.Results {
		if u.CuckooFilter.NotContains(place.PlaceID) {
			var wkt []string
			if place.OpeningHours != nil {
				wkt = place.OpeningHours.WeekdayText
			} else {
				wkt = []string{}
			}
			newPlace := entity.NewPlace(
				place.PlaceID,
				place.Name,
				place.FormattedAddress,
				place.Geometry.Location.Lat,
				place.Geometry.Location.Lng,
				place.Icon,
				place.Types,
				wkt,
			)
			err := u.PlacesGateway.Create(ctx, newPlace)
			if err != nil {
				return nil, errors.New("Failed to persist place while google maps nearby search -> " + err.Error())
			}
			u.CuckooFilter.Add(place.PlaceID)
			output = append(output, newPlace)
		}
	}

	return output, nil
}

func (u *FindNearbyPlacesUseCase) GetNearbyPlacesFromGoogleOrRepository(ctx context.Context, input dto.FindNearbyPlacesInputDTO) ([]*entity.Place, error) {
	if input.IsFromGoogle {
		places, err := u.GetNearbyPlacesFromGoogleMapsWithPersistence(ctx, input)
		if err != nil {
			return nil, errors.New("Failed to find nearby places from Google -> " + err.Error())
		}
		return places, nil
	}
	places, err := u.PlacesGateway.FindNearbyPlaces(ctx, input.Lat, input.Lng, input.Radius)
	if err != nil {
		return nil, errors.New("Failed to find nearby places -> " + err.Error())
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
