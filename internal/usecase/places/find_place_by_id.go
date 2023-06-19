package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/infra/blooms"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindPlaceByIDUseCase struct {
	PlacesGateway gateway.PlacesGateway
	BloomFilter   *blooms.BloomFilter
}

func NewFindPlaceByIdUseCase(placesGateway gateway.PlacesGateway, bloomFilter *blooms.BloomFilter) *FindPlaceByIDUseCase {
	return &FindPlaceByIDUseCase{
		PlacesGateway: placesGateway,
		BloomFilter:   bloomFilter,
	}
}

func (u *FindPlaceByIDUseCase) Execute(ctx context.Context, input dto.FindPlaceByIDInputDTO) (*dto.FindPlaceByIDOutputDTO, error) {
	place, err := u.PlacesGateway.FindPlaceById(ctx, input.ID)
	if err != nil {
		return nil, errors.New("place not found: " + err.Error())
	}
	return &dto.FindPlaceByIDOutputDTO{
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
		Reviews:         place.Reviews,
	}, nil

}
