package dto

import (
	"time"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateUserInputDTO struct {
	ID        uuid.UUID `json:"id"`
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
	UserId     string
	ChagedUser entity.User
}

type FindUserByEmailInputDTO struct {
	Email string
}
type FindUserByEmailOutputDTO struct {
	User entity.User
}

type AuthenticateJwtUserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateJwtUserOutputDTO struct {
	Token string `json:"token"`
}

type AddFavoritesInputDTO struct {
	UserId  string `json:"userId"`
	PlaceId string `json:"placeId"`
}
