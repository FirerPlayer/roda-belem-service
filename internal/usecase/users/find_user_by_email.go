package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
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
		ID:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Points:    user.Points,
		Missions:  user.Missions,
		Favorites: user.Favorites,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
