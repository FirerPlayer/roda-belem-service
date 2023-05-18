package repository

import "github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"

type ProductRepository interface {
	Create(place *entity.Place) error
	FindPlaceById(id string) (*entity.Place, error)
	FindPlaceByPlaceId(placeId string) (*entity.Place, error)
	FindPlacesByAccessibilityFeatures(features *entity.AccessibilityFeaturesEnum) ([]*entity.Place, error)
	UpdatePlaceById(id string, place *entity.Place) error
}
