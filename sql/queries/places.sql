-- name: CreatePlace :exec
INSERT INTO places (
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
    rating
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: FindPlaceById :one
SELECT *
FROM places
WHERE id = ?;
-- name: FindPlaceByGooglePlaceId :one
SELECT *
FROM places
WHERE google_place_id = ?;
-- name: FindPlacesNearby :many
SELECT *
FROM places -- distance in meters
WHERE ST_DISTANCE_SPHERE(POINT(lat, lng), POINT(?, ?)) <= ?;
-- name: UpdatePlaceById :exec
UPDATE places
SET google_place_id = ?,
  name = ?,
  formatted_address = ?,
  lat = ?,
  lng = ?,
  icon = ?,
  types = ?,
  opening_periods = ?,
  photos = ?,
  rating = ?
WHERE id = ?;
-- name: DeletePlaceById :exec
DELETE FROM places
WHERE id = $1;
-- name: FindPlacesByAccessibilityFeature :many
SELECT p.id,
  COUNT(*) AS num_reviews
FROM places p
  JOIN reviews r ON p.id = r.place_id
  JOIN accessibility_features af ON r.review_id = af.review_id
WHERE af.feature IN (?)
GROUP BY p.id
HAVING COUNT(DISTINCT af.feature) = ?