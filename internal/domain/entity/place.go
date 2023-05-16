package entity

import (
	"googlemaps.github.io/maps"
)

type AccessibilityFeatures uint32

const (
	AccessibilityFeaturesNone AccessibilityFeatures = iota
	Ramp
)

type GooglePlace struct {
	Place maps.PlacesSearchResult
}

type Place struct {
	Place GooglePlace
	AccessibilityFeatures
}
