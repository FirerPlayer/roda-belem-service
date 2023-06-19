package gateway

import (
	"context"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
)

type UsersGateway interface {
	CreateUser(ctx context.Context, user *entity.User) error
	ListAllUsers(ctx context.Context) ([]*entity.User, error)
	FindUserById(ctx context.Context, id string) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUserById(ctx context.Context, id string, user *entity.User) error
	DeleteUserById(ctx context.Context, id string) error
	UpdateUserPointsByUserId(ctx context.Context, userId string, points int) error
	AddFavoriteByUserIdAndPlaceId(ctx context.Context, userId string, placeId string) error
	DeleteFavoriteByUserIdAndPlaceId(ctx context.Context, userId string, placeId string) error
	// FindFavoritesByUserId returns a slice of strings containing the favorites of a user given their user ID.
	//
	// ctx is the context of the request.
	// userId is the ID of the user whose favorites are being searched.
	// It returns a slice of strings representing the favorites of the user and an error if any occurred.
	//
	FindFavoritesByUserId(ctx context.Context, userId string) ([]string, error)
}
