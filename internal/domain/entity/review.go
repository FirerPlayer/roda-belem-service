package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID
	PlaceID   string
	UserID    string
	Text      string
	Images    []Image
	Rating    float64
	Reactions []Reaction
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewReview(placeID, userID, text string, images []Image, rating float64, reactions []Reaction) *Review {
	return &Review{
		ID:        uuid.New(),
		PlaceID:   placeID,
		UserID:    userID,
		Text:      text,
		Images:    images,
		Rating:    rating,
		Reactions: reactions,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
