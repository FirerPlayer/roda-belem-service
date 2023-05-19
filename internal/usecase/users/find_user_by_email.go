package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/gateway"
	"github.com/firerplayer/hexagonal-arch-go/internal/usecase/dto"
)

type FindUserByEmailUsecase struct {
	UsersGateway gateway.UsersGateway
}

func NewFindUserByEmailUsecase(usersGateway gateway.UsersGateway) *FindUserByEmailUsecase {
	return &FindUserByEmailUsecase{
		UsersGateway: usersGateway,
	}
}

// Execute retrieves a user by email.
//
// ctx is the context of the request.
// email is the email of the user to retrieve.
// Returns a pointer to a User and an error if any occurred.
func (uc *FindUserByEmailUsecase) Execute(ctx context.Context, email string) (*dto.FindUserByEmailOutputDTO, error) {
	user, err := uc.UsersGateway.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("failed to find user " + err.Error())
	}
	return &dto.FindUserByEmailOutputDTO{
		User: *user,
	}, nil
}
