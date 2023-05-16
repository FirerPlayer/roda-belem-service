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