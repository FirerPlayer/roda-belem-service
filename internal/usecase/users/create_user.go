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
	err := uc.UsersGateway.CreateUser(ctx, &entity.User{
		ID:        input.ID,
		Email:     input.Email,
		Avatar:    input.Avatar,
		Username:  input.Username,
		Password:  input.Password,
		Points:    input.Points,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})
	if err != nil {
		return errors.New("failed to create user " + err.Error())
	}
	return nil
}
