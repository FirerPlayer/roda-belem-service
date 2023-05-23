package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/infra/blooms"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
	"github.com/google/uuid"
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
	err := uc.PlacesGateway.Create(ctx, &entity.Place{
		ID:              uuid.New(),
		PlaceId:         input.PlaceId,
		Name:            input.Name,
		FormatedAddress: input.FormatedAddress,
		Lat:             input.Lat,
		Lng:             input.Lng,
		Icon:            input.Icon,
		Types:           input.Types,
		OpeningPeriods:  input.OpeningPeriods,
	})
	if err != nil {
		return errors.New("Failed to create place " + err.Error())
	}
	uc.BloomFilter.Add(input.PlaceId)

	return nil
}
