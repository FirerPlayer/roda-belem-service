package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Mission struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Reward       int
	MinimumLevel int
}

func NewMission(name string, description string, reward int, minimumLevel int) *Mission {
	return &Mission{
		ID:           uuid.New(),
		Name:         name,
		Description:  description,
		Reward:       reward,
		MinimumLevel: minimumLevel,
	}
}

type User struct {
	ID        uuid.UUID
	Email     string
	Avatar    []byte
	Username  string
	Password  string
	Points    int
	Missions  []Mission
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email, username, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Username:  username,
		Password:  string(hashedPassword),
		Points:    0,
		Missions:  []Mission{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
