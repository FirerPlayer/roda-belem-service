package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type DeleteFavoriteUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewDeleteFavoriteUseCase(usersGateway gateway.UsersGateway) *DeleteFavoriteUseCase {
	return &DeleteFavoriteUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *DeleteFavoriteUseCase) Execute(ctx context.Context, input dto.DeleteFavoriteByUserIdAndPlaceIdInputDTO) error {
	err := uc.UsersGateway.DeleteFavoriteByUserIdAndPlaceId(ctx, input.UserId, input.PlaceId)
	if err != nil {
		return errors.New("failed to delete favorite -> " + err.Error())
	}
	return nil
}
