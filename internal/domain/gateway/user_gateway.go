package gateway

import (
	"context"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
)

type UsersGateway interface {
	CreateUser(ctx context.Context, user *entity.User) error
	ListAllUsers(ctx context.Context) ([]*entity.User, error)
	FindUserById(ctx context.Context, id string) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUserById(ctx context.Context, id string, user *entity.User) error
	DeleteUserById(ctx context.Context, id string) error
	UpdateUserPointsByUserId(ctx context.Context, userId string, points int) error
	AddFavoritePlaceByUserId(ctx context.Context, userId string, placeId string) error
}
