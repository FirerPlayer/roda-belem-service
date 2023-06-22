package dto

import (
	"time"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
)

type CreateUserInputDTO struct {
	Email     string    `json:"email"`
	Avatar    []byte    `json:"avatar"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListAllUsersOutputDTO struct {
	UserId    string    `json:"userId"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateUserInputDTO struct {
	UserId    string           `json:"userId"`
	Email     string           `json:"email"`
	Avatar    []byte           `json:"avatar"`
	Username  string           `json:"username"`
	Password  string           `json:"password"`
	Points    int              `json:"points"`
	Missions  []entity.Mission `json:"missions"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

type FindUserByIDInputDTO struct {
	UserId string `json:"userId"`
}

type FindUserByIDOutputDTO struct {
	ID        string           `json:"id"`
	Email     string           `json:"email"`
	Username  string           `json:"username"`
	Avatar    []byte           `json:"avatar"`
	Points    int              `json:"points"`
	Missions  []entity.Mission `json:"missions"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

type FindUserByEmailInputDTO struct {
	Email string `json:"email"`
}
type FindUserByEmailOutputDTO struct {
	ID        string           `json:"id"`
	Email     string           `json:"email"`
	Username  string           `json:"username"`
	Avatar    []byte           `json:"avatar"`
	Points    int              `json:"points"`
	Missions  []entity.Mission `json:"missions"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

type AuthenticateJwtUserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateJwtUserOutputDTO struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type AddFavoritesInputDTO struct {
	UserId  string `json:"userId"`
	PlaceId string `json:"placeId"`
}

type DeleteUserByIdInputDTO struct {
	UserId string `json:"userId"`
}

type DeleteFavoriteByUserIdAndPlaceIdInputDTO struct {
	UserId  string `json:"userId"`
	PlaceId string `json:"placeId"`
}

type FindFavoritesByUserIdInputDTO struct {
	UserId string `json:"userId"`
}

type FindFavoritesByUserIdOutputDTO struct {
	PlaceId string `json:"placeId"`
}

type UpdateUserPointsByUserIDInputDTO struct {
	UserId string `json:"userId"`
	Points int    `json:"points"`
}
