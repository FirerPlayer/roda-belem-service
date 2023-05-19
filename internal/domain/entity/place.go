package entity

import (
	"googlemaps.github.io/maps"
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
	Place                 maps.PlacesSearchResult
	AccessibilityFeatures []AccessibilityFeaturesEnum
	Reviews               []Review
}

func NewPlace(place maps.PlacesSearchResult, accessibilityFeatures []AccessibilityFeaturesEnum, reviews []Review) *Place {
	return &Place{
		Place:                 place,
		AccessibilityFeatures: accessibilityFeatures,
		Reviews:               reviews,
	}
}
