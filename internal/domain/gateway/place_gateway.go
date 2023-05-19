package gateway

import (
	"context"

	"github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"
)

type PlacesGateway interface {
	Create(ctx context.Context, place *entity.Place) error
	FindPlaceById(ctx context.Context, id string) (*entity.Place, error)
	FindPlaceByPlaceId(ctx context.Context, placeId string) (*entity.Place, error)
	FindPlacesByAccessibilityFeatures(ctx context.Context, features *entity.AccessibilityFeaturesEnum) ([]*entity.Place, error)
	UpdatePlaceById(ctx context.Context, id string, place *entity.Place) error
	DeletePlaceById(ctx context.Context, id string) error
}
