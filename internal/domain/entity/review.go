package entity

import (
	"time"

	"github.com/google/uuid"
)

type ReactionEnum string

const (
	Like   ReactionEnum = "Like"
	Love   ReactionEnum = "Love"
	Thanks ReactionEnum = "Thanks"
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

// type AccessibilityFeaturesEnum string

// const (
// 	Ramp                      = "Ramp"
// 	Elevator                  = "Elevator"
// 	AdaptedBathroom           = "AdaptedBathroom"
// 	BrailleSignage            = "BrailleSignage"
// 	WideCirculationAreas      = "WideCirculationAreas"
// 	ReservedParking           = "ReservedParking"
// 	TactilePaving             = "TactilePaving"
// 	AdaptedTelephones         = "AdaptedTelephones"
// 	VideoIntercom             = "VideoIntercom"
// 	AdaptedTablesCounters     = "AdaptedTablesCounters"
// 	WheelchairAvailability    = "WheelchairAvailability"
// 	SignLanguageCommunication = "SignLanguageCommunication"
// 	GuideDogAllowed           = "GuideDogAllowed"
// 	OnlineAccessibility       = "OnlineAccessibility"
// 	AssistiveTechnologyAccess = "AssistiveTechnologyAccess"
// )

type Review struct {
	ID                    uuid.UUID
	PlaceID               string
	UserID                string
	Content               string
	Photos                []Photo
	Rating                float64
	Reactions             []Reaction
	AccessibilityFeatures []string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func NewReview(placeID, userID, text string, photos []Photo, rating float64, reactions []Reaction) *Review {
	return &Review{
		ID:                    uuid.New(),
		PlaceID:               placeID,
		UserID:                userID,
		Content:               text,
		Photos:                photos,
		Rating:                rating,
		Reactions:             reactions,
		AccessibilityFeatures: []string{},
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
}
