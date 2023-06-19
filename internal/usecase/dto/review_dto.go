package dto

import (
	"time"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
)

type FindReviewByIDInputDTO struct {
	Id string
}

type FindReviewByIDOutputDTO struct {
	ID                    string            `json:"id"`
	PlaceID               string            `json:"placeId"`
	UserID                string            `json:"userId"`
	Text                  string            `json:"text"`
	Photos                []entity.Photo    `json:"photos"`
	Rating                float64           `json:"rating"`
	Reactions             []entity.Reaction `json:"reactions"`
	AccessibilityFeatures []string          `json:"accessibilityFeatures"`
	CreatedAt             time.Time         `json:"createdAt"`
	UpdatedAt             time.Time         `json:"updatedAt"`
}

type FindReviewsByPlaceIDInputDTO struct {
	PlaceID string
	Limit   int
	Offset  int
}

type FindReviewsByPlaceIDOutputDTO struct {
	ID                    string            `json:"id"`
	PlaceID               string            `json:"placeId"`
	UserID                string            `json:"userId"`
	Text                  string            `json:"text"`
	Photos                []entity.Photo    `json:"photos"`
	Rating                float64           `json:"rating"`
	Reactions             []entity.Reaction `json:"reactions"`
	AccessibilityFeatures []string          `json:"accessibilityFeatures"`
	CreatedAt             time.Time         `json:"createdAt"`
	UpdatedAt             time.Time         `json:"updatedAt"`
}

type FindReviewsByUserIDInputDTO struct {
	UserID string
	Limit  int
	Offset int
}

type FindReviewsByUserIDOutputDTO struct {
	ID                    string            `json:"id"`
	PlaceID               string            `json:"placeId"`
	UserID                string            `json:"userId"`
	Text                  string            `json:"text"`
	Photos                []entity.Photo    `json:"photos"`
	Rating                float64           `json:"rating"`
	Reactions             []entity.Reaction `json:"reactions"`
	AccessibilityFeatures []string          `json:"accessibilityFeatures"`
	CreatedAt             time.Time         `json:"createdAt"`
	UpdatedAt             time.Time         `json:"updatedAt"`
}

type AddAccessibilityFeaturesByReviewIDInputDTO struct {
	ReviewID string
	Features []string
}

type DeleteReviewByIDInputDTO struct {
	ID string
}

type UpdateReviewByIDInputDTO struct {
	ID                    string            `json:"id"`
	Content               string            `json:"content"`
	Images                []entity.Photo    `json:"images"`
	Rating                float64           `json:"rating"`
	Reactions             []entity.Reaction `json:"reactions"`
	AccessibilityFeatures []string          `json:"accessibilityFeatures"`
}
