package dto

import (
	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
)

type CreatePlaceInputDTO struct {
	GooglePlaceId   string   `json:"GooglePlaceId"`
	Name            string   `json:"name"`
	FormatedAddress string   `json:"formatedAddress"`
	Lat             float64  `json:"lat"`
	Lng             float64  `json:"lng"`
	Icon            string   `json:"icon"`
	Types           []string `json:"types"`
	OpeningPeriods  []string `json:"openingPeriods"`
}

type FindNearbyPlacesInputDTO struct {
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	Radius       float64 `json:"radius"`
	IsFromGoogle bool    `json:"isFromGoogle"`
}

type FindNearbyPlacesOutputDTO struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	FormatedAddress string         `json:"formatedAddress"`
	Lat             float64        `json:"lat"`
	Lng             float64        `json:"lng"`
	Icon            string         `json:"icon"`
	Types           []string       `json:"types"`
	OpeningPeriods  []string       `json:"openingPeriods"`
	Photos          []entity.Photo `json:"photo"`
	Rating          float64        `json:"rating"`
}

type FindPlaceByIDInputDTO struct {
	ID string `json:"id"`
}

type FindPlaceByIDOutputDTO struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	FormatedAddress string         `json:"formatedAddress"`
	Lat             float64        `json:"lat"`
	Lng             float64        `json:"lng"`
	Icon            string         `json:"icon"`
	Types           []string       `json:"types"`
	OpeningPeriods  []string       `json:"openingPeriods"`
	Photos          []entity.Photo `json:"photo"`
	Rating          float64        `json:"rating"`
}

type FindPlacesByAccessibilityFeatureInputDTO struct {
	AccessibilityFeature string `json:"accessibilityFeature"`
}

type FindPlacesByAccessibilityFeatureOutputDTO struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	FormatedAddress string         `json:"formatedAddress"`
	Lat             float64        `json:"lat"`
	Lng             float64        `json:"lng"`
	Icon            string         `json:"icon"`
	Types           []string       `json:"types"`
	OpeningPeriods  []string       `json:"openingPeriods"`
	Photos          []entity.Photo `json:"photo"`
	Rating          float64        `json:"rating"`
}

type UpdatePlaceByIDInputDTO struct {
	PlaceToUpadteID string         `json:"id"`
	GooglePlaceId   string         `json:"googlePlaceId"`
	Name            string         `json:"name"`
	FormatedAddress string         `json:"formatedAddress"`
	Lat             float64        `json:"lat"`
	Lng             float64        `json:"lng"`
	Icon            string         `json:"icon"`
	Types           []string       `json:"types"`
	OpeningPeriods  []string       `json:"openingPeriods"`
	Photos          []entity.Photo `json:"photos"`
}

type DeletePlaceByIDInputDTO struct {
	ID string `json:"id"`
}
