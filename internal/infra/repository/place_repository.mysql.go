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
	FindPlacesByAccessibilityFeature(ctx context.Context, feature string) ([]*entity.Place, error)
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
		GooglePlaceID:    sql.NullString{String: place.GooglePlaceId, Valid: true},
		Name:             sql.NullString{String: place.Name, Valid: true},
		FormattedAddress: sql.NullString{String: place.FormattedAddress, Valid: true},
		Lat:              sql.NullFloat64{Float64: place.Lat, Valid: true},
		Lng:              sql.NullFloat64{Float64: place.Lng, Valid: true},
		Icon:             sql.NullString{String: place.Icon, Valid: true},
		Types:            types,
		OpeningPeriods:   openingPeriods,
		Photos:           photos,
		Rating:           sql.NullFloat64{Float64: place.Rating, Valid: true},
		CreatedAt:        sql.NullTime{Time: place.CreatedAt, Valid: true},
		UpdatedAt:        sql.NullTime{Time: place.UpdatedAt, Valid: true},
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
	var place entity.Place
	err = HydratePlace(placeDb, &place)
	if err != nil {
		return nil, err
	}

	return &place, nil

}

func (p *PlaceRepositoryMySQL) FindPlacesByAccessibilityFeature(ctx context.Context, feature string) ([]*entity.Place, error) {
	placesDb, err := p.Queries.FindPlacesByAccessibilityFeature(ctx, feature)
	if err != nil {
		return nil, err
	}
	var output []*entity.Place
	for _, placeDb := range placesDb {
		var place entity.Place
		err := HydratePlace(placeDb, &place)
		if err != nil {
			return nil, err
		}
		output = append(output, &place)
	}

	return output, nil
}

func (p *PlaceRepositoryMySQL) FindNearbyPlaces(ctx context.Context, latitude float64, longitude float64, radius float64) ([]*entity.Place, error) {
	params := db.FindPlacesNearbyParams{
		POINT:   latitude,
		POINT_2: longitude,
		Lat:     sql.NullFloat64{Float64: radius, Valid: true},
	}

	places, err := p.Queries.FindPlacesNearby(ctx, params)
	if err != nil {
		return nil, err
	}
	var output []*entity.Place
	for _, pl := range places {
		var place entity.Place
		err := HydratePlace(pl, &place)
		if err != nil {
			return nil, err
		}
		output = append(output, &place)
	}
	return output, nil
}

func (p *PlaceRepositoryMySQL) UpdatePlaceById(ctx context.Context, id string, place *entity.Place) error {
	gooId := sql.NullString{String: place.GooglePlaceId, Valid: true}
	if place.GooglePlaceId == "" {
		gooId.Valid = false
	}
	args := db.UpdatePlaceByIdParams{
		ID:               place.ID.String(),
		GooglePlaceID:    gooId,
		Name:             sql.NullString{String: place.Name, Valid: true},
		FormattedAddress: sql.NullString{String: place.FormattedAddress, Valid: true},
		Lat:              sql.NullFloat64{Float64: place.Lat, Valid: true},
		Lng:              sql.NullFloat64{Float64: place.Lng, Valid: true},
		Icon:             sql.NullString{String: place.Icon, Valid: true},
		Rating:           sql.NullFloat64{Float64: place.Rating, Valid: true},
		UpdatedAt:        sql.NullTime{Time: place.UpdatedAt, Valid: true},
	}
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
	args.Types = types
	args.OpeningPeriods = openingPeriods
	args.Photos = photos

	if err := p.Queries.UpdatePlaceById(ctx, args); err != nil {
		return err
	}
	return nil
}

func (p *PlaceRepositoryMySQL) DeletePlaceById(ctx context.Context, id string) error {
	if err := p.Queries.DeletePlaceById(ctx, id); err != nil {
		return err
	}
	return nil
}

// HydratePlace hydrates a Place entity with data from a Place database object.
//
// placeDb: The Place database object to hydrate from.
// place: The Place entity to hydrate.
// error: Returns an error if there was a problem with hydrating the Place entity.
func HydratePlace(placeDb db.Place, place *entity.Place) error {
	id, err := uuid.Parse(placeDb.ID)
	if err != nil {
		return err
	}
	place.ID = id
	place.GooglePlaceId = placeDb.GooglePlaceID.String
	place.Name = placeDb.Name.String
	place.FormattedAddress = placeDb.FormattedAddress.String
	place.Lat = placeDb.Lat.Float64
	place.Lng = placeDb.Lng.Float64
	place.Icon = placeDb.Icon.String
	err = json.Unmarshal(placeDb.Types, &place.Types)
	if err != nil {
		return err
	}
	err = json.Unmarshal(placeDb.OpeningPeriods, &place.OpeningPeriods)
	if err != nil {
		return err
	}
	err = json.Unmarshal(placeDb.Photos, &place.Photos)
	if err != nil {
		return err
	}
	place.Rating = placeDb.Rating.Float64
	place.CreatedAt = placeDb.CreatedAt.Time
	place.UpdatedAt = placeDb.UpdatedAt.Time
	return nil
}
