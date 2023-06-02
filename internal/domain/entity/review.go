package entity

import (
	"time"

	"github.com/google/uuid"
)

type ReactionEnum string

const (
	Like   ReactionEnum = "like"
	Love   ReactionEnum = "love"
	Thanks ReactionEnum = "thanks"
)

type Reaction struct {
	React  ReactionEnum
	UserID string
}

func NewReaction(react ReactionEnum, userID string) *Reaction {
	return &Reaction{
		React:  react,
		UserID: userID,
	}
}

type Photo struct {
	Id          uuid.UUID
	Data        []byte
	Description string
}

func NewPhoto(data []byte, description string) *Photo {
	return &Photo{
		Id:          uuid.New(),
		Data:        data,
		Description: description,
	}
}

type AccessibilityFeaturesEnum int

const (
	Ramp = iota
	Elevator
	AdaptedBathroom
	BrailleSignage
	WideCirculationAreas
	ReservedParking
	TactilePaving
	AdaptedTelephones
	VideoIntercom
	AdaptedTablesCounters
	WheelchairAvailability
	SignLanguageCommunication
	GuideDogAllowed
	OnlineAccessibility
	AssistiveTechnologyAccess
)

type Review struct {
	ID                    uuid.UUID
	PlaceID               string
	UserID                string
	Text                  string
	Photos                []Photo
	Rating                float64
	Reactions             []Reaction
	AccessibilityFeatures []AccessibilityFeaturesEnum
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func NewReview(placeID, userID, text string, photos []Photo, rating float64, reactions []Reaction) *Review {
	return &Review{
		ID:                    uuid.New(),
		PlaceID:               placeID,
		UserID:                userID,
		Text:                  text,
		Photos:                photos,
		Rating:                rating,
		Reactions:             reactions,
		AccessibilityFeatures: []AccessibilityFeaturesEnum{},
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
}
