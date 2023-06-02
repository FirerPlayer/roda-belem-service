package entity

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID              uuid.UUID
	GooglePlaceId   string
	Name            string
	FormatedAddress string
	Lat             float64
	Lng             float64
	Icon            string
	Types           []string
	OpeningPeriods  []string
	Photos          []Photo
	Rating          float64
	Reviews         []Review
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewPlace(googlePlaceId string, name string, formatedAddress string, lat float64, lng float64, icon string, types []string, openingPeriods []string) *Place {
	return &Place{
		ID:              uuid.New(),
		GooglePlaceId:   googlePlaceId,
		Name:            name,
		FormatedAddress: formatedAddress,
		Lat:             lat,
		Lng:             lng,
		Icon:            icon,
		Types:           types,
		OpeningPeriods:  openingPeriods,
		Photos:          []Photo{},
		Rating:          0.0,
		Reviews:         []Review{},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
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
