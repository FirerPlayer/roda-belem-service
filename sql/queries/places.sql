-- name: CreatePlace :exec
INSERT INTO places (
    id,
    place_id,
    name,
    formatted_address,
    lat,
    lng,
    icon,
    types,
    opening_periods,
    photos,
    rating,
    accessibility_features
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
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
WHERE ST_DISTANCE_SPHERE(POINT(lat, lng), POINT(?, ?)) <= ?;
-- name: FindPlacesByAccessibilityFeatures :many
SELECT *
FROM places
WHERE accessibility_features = ?;
-- name: UpdatePlaceById :exec
UPDATE places
SET place_id = ?,
  name = ?,
  formatted_address = ?,
  lat = ?,
  lng = ?,
  icon = ?,
  types = ?,
  opening_periods = ?,
  photos = ?,
  rating = ?,
  accessibility_features = ?
WHERE id = ?;
-- name: DeletePlaceById :exec
DELETE FROM places
WHERE id = $1;