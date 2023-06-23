package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type CreateUserUsecase struct {
	UsersGateway gateway.UsersGateway
}

func NewCreateUserUseCase(usersGateway gateway.UsersGateway) *CreateUserUsecase {
	return &CreateUserUsecase{
		UsersGateway: usersGateway,
	}
}

// Execute creates a new user in the database using the CreateUser method of the UsersGateway struct.
//
// ctx is the context of the request.
// input is a CreateUserInputDTO struct containing the information of the user to be created.
// error is returned if there are any issues creating the user.
func (uc *CreateUserUsecase) Execute(ctx context.Context, input dto.CreateUserInputDTO) error {
	newUser, err := entity.NewUser(input.Email, input.Username, input.Password)
	if err != nil {
		return errors.New("failed to create user " + err.Error())
	}
	err = uc.UsersGateway.CreateUser(ctx, newUser)
	if err != nil {
		return errors.New("failed to create user " + err.Error())
	}
	return nil
}
