package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type DeletePlaceByIDUseCase struct {
	PlacesGateway gateway.PlacesGateway
}

func NewDeletePlaceByIDUseCase(placesGateway gateway.PlacesGateway) *DeletePlaceByIDUseCase {
	return &DeletePlaceByIDUseCase{
		PlacesGateway: placesGateway,
	}
}

// Execute deletes a place by ID from the PlacesGateway.
//
// ctx - context object used to cancel the operation if needed.
// input - input DTO that contains the ID of the place to be deleted.
// error - returns any errors that occurred while deleting the place.
func (uc *DeletePlaceByIDUseCase) Execute(ctx context.Context, input dto.DeletePlaceByIDInputDTO) error {
	err := uc.PlacesGateway.DeletePlaceById(ctx, input.ID)
	if err != nil {
		return errors.New("Failed to delete place " + err.Error())
	}
	return nil
}
