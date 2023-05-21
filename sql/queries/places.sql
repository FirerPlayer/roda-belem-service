-- name: CreatePlace :execute
INSERT INTO places (
    id,
    place_id,
    name,
    formatted_address,
    coordinates,
    icon,
    types,
    opening_periods,
    photos,
    rating,
    accessibility_features
  )
VALUES (?, ?, ?, ?, POINT(?, ?), ?, ?, ?, ?, ?, ?);
-- name: FindPlaceById :one
SELECT *
FROM places
WHERE id = ?;
-- name: FindPlaceByPlaceId :one
SELECT *
FROM places
WHERE place_id = ?;
-- name: FindPlacesNearby :many
SELECT *
FROM places -- distance in meters
WHERE ST_DISTANCE_SPHERE(coordinates, POINT(?, ?)) <= ?;
-- name: FindPlacesByAccessibilityFeatures :many
SELECT *
FROM places
WHERE AccessibilityFeatures = ?;
-- name: UpdatePlaceById :execute
UPDATE places
SET ?
WHERE id = ?;
-- name: DeletePlaceById :execute
DELETE FROM places
WHERE id = ?;