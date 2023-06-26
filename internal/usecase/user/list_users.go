package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type ListAllUsersUsecase struct {
	UsersGateway gateway.UsersGateway
}

func NewListAllUsersUsecase(usersGateway gateway.UsersGateway) *ListAllUsersUsecase {
	return &ListAllUsersUsecase{
		UsersGateway: usersGateway,
	}
}

func (uc *ListAllUsersUsecase) Execute(ctx context.Context) ([]*dto.ListAllUsersOutputDTO, error) {
	users, err := uc.UsersGateway.ListAllUsers(ctx)
	if err != nil {
		return nil, errors.New("Failed to list users -> " + err.Error())
	}
	var output []*dto.ListAllUsersOutputDTO
	for _, user := range users {
		output = append(output, &dto.ListAllUsersOutputDTO{
			UserId:    user.ID.String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}
	return output, nil
}
