package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Avatar    []byte    `json:"avatar"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Points    int       `json:"points"`
	Missions  []Mission `json:"missions"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(email, username, password string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Username:  username,
		Password:  password,
		Points:    0,
		Missions:  []Mission{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
