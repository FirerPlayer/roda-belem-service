package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindUserByIDUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewFindUserByIdUseCase(usersGateway gateway.UsersGateway) *FindUserByIDUseCase {
	return &FindUserByIDUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *FindUserByIDUseCase) Execute(ctx context.Context, id string) (*dto.FindUserByIDOutputDTO, error) {
	user, err := uc.UsersGateway.FindUserById(ctx, id)
	if err != nil {
		return nil, errors.New("Failed to find user -> " + err.Error())
	}

	return &dto.FindUserByIDOutputDTO{
		ID:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		Points:    user.Points,
		Avatar:    user.Avatar,
		Missions:  user.Missions,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
