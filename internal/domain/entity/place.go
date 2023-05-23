package entity

import (
	"github.com/google/uuid"
)

type AccessibilityFeaturesEnum string

const (
	Ramp                      AccessibilityFeaturesEnum = "Ramp"
	Elevator                  AccessibilityFeaturesEnum = "Elevator"
	AdaptedBathroom           AccessibilityFeaturesEnum = "AdaptedBathroom"
	BrailleSignage            AccessibilityFeaturesEnum = "BrailleSignage"
	WideCirculationAreas      AccessibilityFeaturesEnum = "WideCirculationAreas"
	ReservedParking           AccessibilityFeaturesEnum = "ReservedParking"
	TactilePaving             AccessibilityFeaturesEnum = "TactilePaving"
	AdaptedTelephones         AccessibilityFeaturesEnum = "AdaptedTelephones"
	VideoIntercom             AccessibilityFeaturesEnum = "VideoIntercom"
	AdaptedTablesCounters     AccessibilityFeaturesEnum = "AdaptedTablesCounters"
	WheelchairAvailability    AccessibilityFeaturesEnum = "WheelchairAvailability"
	SignLanguageCommunication AccessibilityFeaturesEnum = "SignLanguageCommunication"
	GuideDogAllowed           AccessibilityFeaturesEnum = "GuideDogAllowed"
	OnlineAccessibility       AccessibilityFeaturesEnum = "OnlineAccessibility"
	AssistiveTechnologyAccess AccessibilityFeaturesEnum = "AssistiveTechnologyAccess"
)

type Place struct {
	ID                    uuid.UUID
	PlaceId               string
	Name                  string
	FormatedAddress       string
	Lat                   float64
	Lng                   float64
	Icon                  string
	Types                 []string
	OpeningPeriods        []string
	Photos                []Photo
	Rating                float64
	AccessibilityFeatures []AccessibilityFeaturesEnum
	Reviews               []Review
}

func NewPlace(placeId, name, formatedAddress string, lat, lng float64, icon string, types []string, openingPeriods []string, photos []Photo, rating float64, accessibilityFeatures []AccessibilityFeaturesEnum, reviews []Review) *Place {
	return &Place{
		ID:                    uuid.New(),
		PlaceId:               placeId,
		Name:                  name,
		FormatedAddress:       formatedAddress,
		Lat:                   lat,
		Lng:                   lng,
		Icon:                  icon,
		Types:                 types,
		OpeningPeriods:        openingPeriods,
		Photos:                photos,
		Rating:                rating,
		AccessibilityFeatures: accessibilityFeatures,
		Reviews:               reviews,
	}

}

func (p *Place) RefreshRatings() {
	rating := 0.0
	for _, review := range p.Reviews {
		rating += review.Rating
	}
	p.Rating = rating / float64(len(p.Reviews))
}

func (p *Place) GetRating() float64 {
	return p.Rating
}
