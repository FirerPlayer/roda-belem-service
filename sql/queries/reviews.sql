-- name: CreateReview :exec
INSERT INTO reviews (
    id,
    place_id,
    user_id,
    text,
    images,
    rating,
    reactions,
    created_at,
    updated_at
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: FindReviewsByPlaceId :many
SELECT *
FROM reviews
WHERE place_id = ?;
-- name: FindReviewsByUserId :many
SELECT *
FROM reviews
WHERE user_id = ?;
-- name: FindReviewById :one
SELECT *
FROM reviews
WHERE id = ?;