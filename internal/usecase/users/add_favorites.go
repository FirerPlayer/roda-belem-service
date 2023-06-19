package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type AddFavoritesUseCase struct {
	UsersGateway  gateway.UsersGateway
	PlacesGateway gateway.PlacesGateway
}

func NewAddFavoritesUseCase(usersGateway gateway.UsersGateway) *AddFavoritesUseCase {
	return &AddFavoritesUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *AddFavoritesUseCase) Execute(ctx context.Context, input dto.AddFavoritesInputDTO) error {
	_, err := uc.UsersGateway.FindUserById(ctx, input.UserId)
	if err != nil {
		return errors.New("User not found " + err.Error())
	}
	_, err = uc.PlacesGateway.FindPlaceById(ctx, input.PlaceId)
	if err != nil {
		return errors.New("Place not found " + err.Error())
	}

	err = uc.UsersGateway.AddFavoriteByUserIdAndPlaceId(ctx, input.UserId, input.PlaceId)
	if err != nil {
		return errors.New("Failed to add favorite place " + err.Error())
	}
	return nil

}
