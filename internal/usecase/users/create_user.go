package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
)

type CreateUserUsecase struct {
	UsersGateway gateway.UsersGateway
}

func NewCreateUserUsecase(usersGateway gateway.UsersGateway) *CreateUserUsecase {
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
	newUser := entity.NewUser(input.Email, input.Username, input.Password)
	err := uc.UsersGateway.CreateUser(ctx, newUser)
	if err != nil {
		return errors.New("failed to create user " + err.Error())
	}
	return nil
}
