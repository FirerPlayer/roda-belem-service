package entity

import "github.com/google/uuid"

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
