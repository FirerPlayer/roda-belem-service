// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: places.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createPlace = `-- name: CreatePlace :exec

INSERT INTO
    places (
        id,
        google_place_id,
        name,
        formatted_address,
        lat,
        lng,
        icon,
        types,
        opening_periods,
        photos,
        rating,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreatePlaceParams struct {
	ID               string
	GooglePlaceID    sql.NullString
	Name             sql.NullString
	FormattedAddress sql.NullString
	Lat              sql.NullFloat64
	Lng              sql.NullFloat64
	Icon             sql.NullString
	Types            json.RawMessage
	OpeningPeriods   json.RawMessage
	Photos           json.RawMessage
	Rating           sql.NullFloat64
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
}

func (q *Queries) CreatePlace(ctx context.Context, arg CreatePlaceParams) error {
	_, err := q.db.ExecContext(ctx, createPlace,
		arg.ID,
		arg.GooglePlaceID,
		arg.Name,
		arg.FormattedAddress,
		arg.Lat,
		arg.Lng,
		arg.Icon,
		arg.Types,
		arg.OpeningPeriods,
		arg.Photos,
		arg.Rating,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deletePlaceById = `-- name: DeletePlaceById :exec

DELETE FROM places WHERE id = ?
`

func (q *Queries) DeletePlaceById(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePlaceById, id)
	return err
}

const findPlaceByGooglePlaceId = `-- name: FindPlaceByGooglePlaceId :one

SELECT id, google_place_id, name, formatted_address, lat, lng, icon, types, opening_periods, photos, rating, created_at, updated_at FROM places WHERE google_place_id = ?
`

func (q *Queries) FindPlaceByGooglePlaceId(ctx context.Context, googlePlaceID sql.NullString) (Place, error) {
	row := q.db.QueryRowContext(ctx, findPlaceByGooglePlaceId, googlePlaceID)
	var i Place
	err := row.Scan(
		&i.ID,
		&i.GooglePlaceID,
		&i.Name,
		&i.FormattedAddress,
		&i.Lat,
		&i.Lng,
		&i.Icon,
		&i.Types,
		&i.OpeningPeriods,
		&i.Photos,
		&i.Rating,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPlaceById = `-- name: FindPlaceById :one

SELECT id, google_place_id, name, formatted_address, lat, lng, icon, types, opening_periods, photos, rating, created_at, updated_at FROM places WHERE id = ?
`

func (q *Queries) FindPlaceById(ctx context.Context, id string) (Place, error) {
	row := q.db.QueryRowContext(ctx, findPlaceById, id)
	var i Place
	err := row.Scan(
		&i.ID,
		&i.GooglePlaceID,
		&i.Name,
		&i.FormattedAddress,
		&i.Lat,
		&i.Lng,
		&i.Icon,
		&i.Types,
		&i.OpeningPeriods,
		&i.Photos,
		&i.Rating,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPlacesByAccessibilityFeature = `-- name: FindPlacesByAccessibilityFeature :many

SELECT p.id, p.google_place_id, p.name, p.formatted_address, p.lat, p.lng, p.icon, p.types, p.opening_periods, p.photos, p.rating, p.created_at, p.updated_at
FROM places p
    JOIN reviews r ON p.id = r.place_id
WHERE
    FIND_IN_SET(?, r.accessibilityFeatures) > 0
`

func (q *Queries) FindPlacesByAccessibilityFeature(ctx context.Context, findINSET string) ([]Place, error) {
	rows, err := q.db.QueryContext(ctx, findPlacesByAccessibilityFeature, findINSET)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Place
	for rows.Next() {
		var i Place
		if err := rows.Scan(
			&i.ID,
			&i.GooglePlaceID,
			&i.Name,
			&i.FormattedAddress,
			&i.Lat,
			&i.Lng,
			&i.Icon,
			&i.Types,
			&i.OpeningPeriods,
			&i.Photos,
			&i.Rating,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findPlacesNearby = `-- name: FindPlacesNearby :many

SELECT id, google_place_id, name, formatted_address, lat, lng, icon, types, opening_periods, photos, rating, created_at, updated_at
FROM
    places -- distance in meters
WHERE
    ST_DISTANCE_SPHERE(POINT(lat, lng), POINT(?, ?)) <= ?
`

type FindPlacesNearbyParams struct {
	POINT   float64
	POINT_2 float64
	Lat     sql.NullFloat64
}

func (q *Queries) FindPlacesNearby(ctx context.Context, arg FindPlacesNearbyParams) ([]Place, error) {
	rows, err := q.db.QueryContext(ctx, findPlacesNearby, arg.POINT, arg.POINT_2, arg.Lat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Place
	for rows.Next() {
		var i Place
		if err := rows.Scan(
			&i.ID,
			&i.GooglePlaceID,
			&i.Name,
			&i.FormattedAddress,
			&i.Lat,
			&i.Lng,
			&i.Icon,
			&i.Types,
			&i.OpeningPeriods,
			&i.Photos,
			&i.Rating,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlaceById = `-- name: UpdatePlaceById :exec

UPDATE places
SET
    google_place_id = ?,
    name = ?,
    formatted_address = ?,
    lat = ?,
    lng = ?,
    icon = ?,
    types = ?,
    opening_periods = ?,
    photos = ?,
    rating = ?,
    updated_at = ?
WHERE id = ?
`

type UpdatePlaceByIdParams struct {
	GooglePlaceID    sql.NullString
	Name             sql.NullString
	FormattedAddress sql.NullString
	Lat              sql.NullFloat64
	Lng              sql.NullFloat64
	Icon             sql.NullString
	Types            json.RawMessage
	OpeningPeriods   json.RawMessage
	Photos           json.RawMessage
	Rating           sql.NullFloat64
	UpdatedAt        sql.NullTime
	ID               string
}

func (q *Queries) UpdatePlaceById(ctx context.Context, arg UpdatePlaceByIdParams) error {
	_, err := q.db.ExecContext(ctx, updatePlaceById,
		arg.GooglePlaceID,
		arg.Name,
		arg.FormattedAddress,
		arg.Lat,
		arg.Lng,
		arg.Icon,
		arg.Types,
		arg.OpeningPeriods,
		arg.Photos,
		arg.Rating,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
