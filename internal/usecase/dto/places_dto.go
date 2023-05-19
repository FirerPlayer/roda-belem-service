package dto

import "github.com/firerplayer/hexagonal-arch-go/internal/domain/entity"

type FindNearbyPlacesInputDTO struct {
	Lat    float64
	Lng    float64
	Radius float64
}

type FindNearbyPlacesOutputDTO struct {
	Places []*entity.Place
}
