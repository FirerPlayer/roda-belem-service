package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/infra/filters"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type CreatePlaceUseCase struct {
	PlacesGateway gateway.PlacesGateway
	CuckooFilter  *filters.CuckooFilter
}

func NewCreatePlaceUseCase(placesGateway gateway.PlacesGateway, cuckooFilter *filters.CuckooFilter) *CreatePlaceUseCase {
	return &CreatePlaceUseCase{
		PlacesGateway: placesGateway,
		CuckooFilter:  cuckooFilter,
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
		return errors.New("Failed to create place -> " + err.Error())
	}
	uc.CuckooFilter.Add(newPlace.ID.String())

	return nil
}
