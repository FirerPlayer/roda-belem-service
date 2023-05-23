package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
)

type UpdateUserUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewUpdateUserUseCase(usersGateway gateway.UsersGateway) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UsersGateway: usersGateway,
	}
}

// Execute updates the user with the given ID according to the provided user data.
//
// ctx is the context that the function executes under.
// input is the input data for updating the user.
// Returns an error if the update operation fails.
func (uc *UpdateUserUseCase) Execute(ctx context.Context, input dto.UpdateUserInputDTO) error {
	err := uc.UsersGateway.UpdateUserById(ctx, input.UserId, &entity.User{
		Email:     input.Email,
		Avatar:    input.Avatar,
		Username:  input.Username,
		Password:  input.Password,
		Points:    input.Points,
		Missions:  input.Missions,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return errors.New("Failed to update user " + err.Error())
	}
	return nil
}
