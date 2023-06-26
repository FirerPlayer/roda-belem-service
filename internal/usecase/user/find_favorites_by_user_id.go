package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
)

type FindFavoritesByUserIDUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewFindFavoritesByUserIDUseCase(usersGateway gateway.UsersGateway) *FindFavoritesByUserIDUseCase {
	return &FindFavoritesByUserIDUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *FindFavoritesByUserIDUseCase) Execute(ctx context.Context, input dto.FindFavoritesByUserIdInputDTO) ([]*dto.FindFavoritesByUserIdOutputDTO, error) {
	favorites, err := uc.UsersGateway.FindFavoritesByUserId(ctx, input.UserId)
	if err != nil {
		return nil, errors.New("Failed to find favorites -> " + err.Error())
	}
	var out []*dto.FindFavoritesByUserIdOutputDTO

	for _, favorite := range favorites {
		out = append(out, &dto.FindFavoritesByUserIdOutputDTO{
			PlaceId: favorite,
		})
	}
	return out, nil
}
