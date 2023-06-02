package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/infra/blooms"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
)

type CreatePlaceUseCase struct {
	PlacesGateway gateway.PlacesGateway
	BloomFilter   *blooms.BloomFilter
}

func NewCreatePlaceUseCase(placesGateway gateway.PlacesGateway, bloomFilter *blooms.BloomFilter) *CreatePlaceUseCase {
	return &CreatePlaceUseCase{
		PlacesGateway: placesGateway,
		BloomFilter:   bloomFilter,
	}
}

func (uc *CreatePlaceUseCase) Execute(ctx context.Context, input *dto.CreatePlaceInputDTO) error {
	newPlace := entity.NewPlace(
		input.GooglePlaceId,
		input.Name,
		input.FormatedAddress,
		input.Lat,
		input.Lng,
		input.Icon,
		input.Types,
		input.OpeningPeriods,
	)
	err := uc.PlacesGateway.Create(ctx, newPlace)
	if err != nil {
		return errors.New("Failed to create place: " + err.Error())
	}
	uc.BloomFilter.Add(newPlace.ID.String())

	return nil
}
