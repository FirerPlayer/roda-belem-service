package dto

import (
	"time"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
)

type FindReviewByIDInputDTO struct {
	Id string
}

type FindReviewByIDOutputDTO struct {
	ID        string            `json:"id"`
	PlaceID   string            `json:"placeId"`
	UserID    string            `json:"userId"`
	Text      string            `json:"text"`
	Photos    []entity.Photo    `json:"photos"`
	Rating    float64           `json:"rating"`
	Reactions []entity.Reaction `json:"reactions"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

type FindReviewsByPlaceIDInputDTO struct {
	PlaceID string
	Limit   int
	Offset  int
}

type FindReviewsByPlaceIDOutputDTO struct {
	ID        string
	PlaceID   string
	UserID    string
	Text      string
	Photos    []entity.Photo
	Rating    float64
	Reactions []entity.Reaction
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FindReviewsByUserIDInputDTO struct {
	UserID string
}

type FindReviewsByUserIDOutputDTO struct {
	ID        string
	PlaceID   string
	UserID    string
	Text      string
	Photos    []entity.Photo
	Rating    float64
	Reactions []entity.Reaction
	CreatedAt time.Time
	UpdatedAt time.Time
}
