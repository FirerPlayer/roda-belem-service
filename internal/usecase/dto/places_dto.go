package dto

import (
	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
)

type CreatePlaceInputDTO struct {
	PlaceId               string                             `json:"placeId"`
	Name                  string                             `json:"name"`
	FormatedAddress       string                             `json:"formatedAddress"`
	Lat                   float64                            `json:"lat"`
	Lng                   float64                            `json:"lng"`
	Icon                  string                             `json:"icon"`
	Types                 []string                           `json:"types"`
	OpeningPeriods        []string                           `json:"openingPeriods"`
	Photos                []entity.Photo                     `json:"photos"`
	Rating                float64                            `json:"rating"`
	AccessibilityFeatures []entity.AccessibilityFeaturesEnum `json:"accessibilityFeatures"`
	Reviews               []entity.Review                    `json:"reviews"`
}

type FindNearbyPlacesInputDTO struct {
	Lat    float64
	Lng    float64
	Radius float64
}

type PhotoDTO struct {
	ID    string `json:"id"`
	Data  []byte `json:"data"`
	Place string `json:"place"`
	
}

type FindNearbyPlacesOutputDTO struct {
	ID                    string                             `json:"id"`
	Name                  string                             `json:"name"`
	FormatedAddress       string                             `json:"formatedAddress"`
	Lat                   float64                            `json:"lat"`
	Lng                   float64                            `json:"lng"`
	Icon                  string                             `json:"icon"`
	Types                 []string                           `json:"types"`
	OpeningPeriod         []string                           `json:"openingPeriod"`
	Photos                []PhotoDTO                         `json:"photo"`
	Rating                float64                            `json:"rating"`
	AccessibilityFeatures []entity.AccessibilityFeaturesEnum `json:"accessibilityFeatures"`
}
