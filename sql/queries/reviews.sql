-- name: CreateReview :exec
INSERT INTO reviews (
    id,
    place_id,
    user_id,
    content,
    images,
    rating,
    reactions,
    accessibility_features,
    created_at,
    updated_at
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: FindReviewsByPlaceId :many
SELECT *
FROM reviews
WHERE place_id = ?
LIMIT ?
OFFSET ?;
-- name: FindReviewsByUserID :many
SELECT *
FROM reviews
WHERE user_id = ?
LIMIT ?
OFFSET ?;
-- name: FindReviewById :one
SELECT *
FROM reviews
WHERE id = ?;
-- name: UpdateReviewById :exec
UPDATE reviews
SET content = ?,
  images = ?,
  rating = ?,
  reactions = ?,
  accessibility_features = ?,
  updated_at = ?
WHERE id = ?;
-- name: DeleteReviewById :exec
DELETE FROM reviews
WHERE id = ?;
-- name: AddAccessibilityFeatureByReviewID :exec
UPDATE reviews
SET accessibility_features = CONCAT(accessibility_features, ',', ?)
WHERE id = ?;
