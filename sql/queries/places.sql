-- name: CreatePlace :exec

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
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: FindPlaceById :one

SELECT * FROM places WHERE id = ?;

-- name: FindPlaceByGooglePlaceId :one

SELECT * FROM places WHERE google_place_id = ?;

-- name: FindPlacesNearby :many

SELECT *
FROM
    places -- distance in meters
WHERE
    ST_DISTANCE_SPHERE(POINT(lat, lng), POINT(?, ?)) <= ?;

-- name: UpdatePlaceById :exec

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
WHERE id = ?;

-- name: DeletePlaceById :exec

DELETE FROM places WHERE id = ?;

-- name: FindPlacesByAccessibilityFeature :many

SELECT p.*
FROM places p
    JOIN reviews r ON p.id = r.place_id
WHERE
    FIND_IN_SET(?, r.accessibilityFeatures) > 0;