-- name: CreateReview :execute
INSERT INTO reviews (
    id,
    placeId,
    userId,
    text,
    images,
    rating,
    reactions,
    createdAt,
    updatedAt
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: FindReviewsByPlaceId :many
SELECT *
FROM reviews
WHERE placeId = ?;
-- name: FindReviewsByUserId :many
SELECT *
FROM reviews
WHERE userId = ?;
-- name: FindReviewById :one
SELECT *
FROM reviews
WHERE id = ?;