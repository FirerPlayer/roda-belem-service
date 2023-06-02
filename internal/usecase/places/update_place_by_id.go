package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
)

type UpdatePLaceByIDUseCase struct {
	PlacesGateway gateway.PlacesGateway
}

func NewUpdatePlaceByIDUseCase(placesGateway gateway.PlacesGateway) *UpdatePLaceByIDUseCase {
	return &UpdatePLaceByIDUseCase{
		PlacesGateway: placesGateway,
	}
}

func (uc *UpdatePLaceByIDUseCase) Execute(ctx context.Context, input dto.UpdatePlaceByIDInputDTO) error {
	err := uc.PlacesGateway.UpdatePlaceById(ctx, input.PlaceToUpadteID, &entity.Place{
		GooglePlaceId:   input.GooglePlaceId,
		Name:            input.Name,
		FormatedAddress: input.FormatedAddress,
		Lat:             input.Lat,
		Lng:             input.Lng,
		Icon:            input.Icon,
		Types:           input.Types,
		OpeningPeriods:  input.OpeningPeriods,
		Photos:          input.Photos,
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		return errors.New("Failed to update place " + err.Error())
	}
	return nil

}
