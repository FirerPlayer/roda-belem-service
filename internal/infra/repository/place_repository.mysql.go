package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	db "github.com/firerplayer/roda-belem-service/internal/infra/mysql"
	"github.com/google/uuid"
)

type PlaceRepositoryMySQL struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewPlaceRepositoryMySQL(dbt *sql.DB) *PlaceRepositoryMySQL {
	return &PlaceRepositoryMySQL{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

/*
type PlacesGateway interface {
	Create(ctx context.Context, place *entity.Place) error
	FindPlaceById(ctx context.Context, id string) (*entity.Place, error)
	FindPlaceByGooglePlaceId(ctx context.Context, GooglePlaceId string) (*entity.Place, error)
	FindPlacesByAccessibilityFeatures(ctx context.Context, features []entity.AccessibilityFeaturesEnum) ([]*entity.Place, error)
	FindNearbyPlaces(ctx context.Context, latitude float64, longitude float64, radius float64) ([]*entity.Place, error)
	UpdatePlaceById(ctx context.Context, id string, place *entity.Place) error
	DeletePlaceById(ctx context.Context, id string) error
}
*/

func (p *PlaceRepositoryMySQL) Create(ctx context.Context, place *entity.Place) error {
	types, err := json.Marshal(place.Types)
	if err != nil {
		return err
	}
	openingPeriods, err := json.Marshal(place.OpeningPeriods)
	if err != nil {
		return err
	}
	photos, err := json.Marshal(place.Photos)
	if err != nil {
		return err
	}
	args := db.CreatePlaceParams{
		ID:               place.ID.String(),
		GooglePlaceID:    sql.NullString{String: place.GooglePlaceId},
		Name:             sql.NullString{String: place.Name},
		FormattedAddress: sql.NullString{String: place.FormattedAddress},
		Lat:              sql.NullFloat64{Float64: place.Lat},
		Lng:              sql.NullFloat64{Float64: place.Lng},
		Icon:             sql.NullString{String: place.Icon},
		Types:            types,
		OpeningPeriods:   openingPeriods,
		Photos:           photos,
		Rating:           sql.NullFloat64{Float64: place.Rating},
	}

	if err := p.Queries.CreatePlace(ctx, args); err != nil {
		return err
	}

	return nil
}

func (p *PlaceRepositoryMySQL) FindPlaceById(ctx context.Context, id string) (*entity.Place, error) {
	placeDb, err := p.Queries.FindPlaceById(ctx, id)
	if err != nil {
		return nil, err
	}
	place := &entity.Place{}
	place.ID = uuid.MustParse(placeDb.ID)
	place.GooglePlaceId = placeDb.GooglePlaceID.String
	place.Name = placeDb.Name.String
	place.FormattedAddress = placeDb.FormattedAddress.String
	place.Lat = placeDb.Lat.Float64
	place.Lng = placeDb.Lng.Float64
	place.Icon = placeDb.Icon.String
	json.Unmarshal(placeDb.Types, &place.Types)
	json.Unmarshal(placeDb.OpeningPeriods, &place.OpeningPeriods)
	json.Unmarshal(placeDb.Photos, &place.Photos)
	place.Rating = placeDb.Rating.Float64
	place.CreatedAt = placeDb.CreatedAt.Time
	place.UpdatedAt = placeDb.UpdatedAt.Time

	return place, nil

}
