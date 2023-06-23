package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type UpdateUserPointsByUserIdUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewUpdateUserPointsByUserIDUseCase(usersGateway gateway.UsersGateway) *UpdateUserPointsByUserIdUseCase {
	return &UpdateUserPointsByUserIdUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *UpdateUserPointsByUserIdUseCase) Execute(ctx context.Context, input dto.UpdateUserPointsByUserIDInputDTO) error {
	err := uc.UsersGateway.UpdateUserPointsByUserId(ctx, input.UserId, input.Points)
	if err != nil {
		return errors.New("failed to update user points " + err.Error())
	}
	return nil
}
