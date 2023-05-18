-- name: CreatePlace :execute
INSERT INTO places (
    id,
    placeId,
    name,
    formatted_address,
    lat,
    lng,
    icon,
    types,
    opening_periods,
    photos,
    rating,
    AccessibilityFeatures
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: FindPlaceById :one
SELECT *
FROM places
WHERE id = ?;
-- name: FindPlaceByPlaceId :one
SELECT *
FROM places
WHERE placeId = ?;
-- name: FindPlacesByAccessibilityFeatures :many
SELECT *
FROM places
WHERE AccessibilityFeatures = ?;
-- name: FindPlacesByRating :many
SELECT *
FROM places
WHERE ABS(rating - ?) = 0.5;
-- name: UpdatePlaceById :execute
UPDATE places
SET ?
WHERE id = ?;
-- name: DeletePlaceById :execute
DELETE FROM places
WHERE id = ?;