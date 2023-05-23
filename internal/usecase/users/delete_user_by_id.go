package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
)

type DeleteUserByIDUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewDeleteUserByIDUseCase(usersGateway gateway.UsersGateway) *DeleteUserByIDUseCase {
	return &DeleteUserByIDUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *DeleteUserByIDUseCase) Execute(ctx context.Context, input dto.DeleteUserByIdInputDTO) error {
	err := uc.UsersGateway.DeleteUserById(ctx, input.UserId)
	if err != nil {
		return errors.New("Failed to delete user " + err.Error())
	}
	return nil
}
